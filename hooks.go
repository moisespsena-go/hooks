package hooks

import (
	"io"
	"os"
	"reflect"
	"runtime"
	"sync"

	_ "github.com/moisespsena-go/default-logger"
	path_helpers "github.com/moisespsena-go/path-helpers"
	"github.com/op/go-logging"
)

type Status bool

func (this Status) Exit() {
	if !this {
		os.Exit(1)
	}
}

type Jobs []Job
type PostJobs []PostJob

func (this PostJobs) Jobs(binNames ...string) (jobs Jobs) {
	jobs = make(Jobs, len(this))
	for i, j := range this {
		func(j PostJob) {
			name := JobName(j)
			jobs[i] = &NamedJob{JobFunc(func() error {
				return j.Run(binNames...)
			}), name}
		}(j)
	}
	return
}

type Runner struct {
	NumCpu    int
	Jobs      Jobs
	Stderr    io.Writer
	failedMux sync.Mutex
	failed    bool
	wg        *sync.WaitGroup
}

func NewRunner(jobs Jobs) *Runner {
	return &Runner{Jobs: jobs}
}

func (this *Runner) SetNumCpu(v int) *Runner {
	this.NumCpu = v
	return this
}

func (this *Runner) Append(jobs ...Job) *Runner {
	this.Jobs = append(this.Jobs, jobs...)
	return this
}

func (this *Runner) worker(id int, jobs <-chan Job) {
	for j := range jobs {
		func() {
			defer this.wg.Done()
			var name string
			if namer, ok := j.(JobNamer); ok {
				name = namer.JobName()
			} else {
				name = runtime.FuncForPC(reflect.ValueOf(j).Pointer()).Name()
			}
			l := logging.MustGetLogger(name)
			l.Info("starting")
			err := j.Run()
			if err != nil {
				this.failedMux.Lock()
				this.failed = true
				this.failedMux.Unlock()
				l.Error(err.Error())
			}
			l.Info("done")
		}()
	}
}

func (this Runner) Ok() bool {
	return !this.failed
}

func (this Runner) runP() (status Status) {
	numJobs := len(this.Jobs)
	jobs := make(chan Job, numJobs)

	for w := 0; w < this.NumCpu; w++ {
		go this.worker(w+1, jobs)
	}

	this.wg = &sync.WaitGroup{}
	this.wg.Add(numJobs)

	for _, j := range this.Jobs {
		jobs <- j
	}

	close(jobs)

	this.wg.Wait()

	return Status(!this.failed)
}

func (this *Runner) run() (status Status) {
	for _, j := range this.Jobs {
		log := logging.MustGetLogger(path_helpers.PkgPathOf(j))
		log.Info("starting")
		err := j.Run()
		if err != nil {
			log.Error(err.Error())
		}
		log.Info("done")
		if err != nil {
			return Status(false)
		}
	}
	return Status(true)
}

func (this Runner) Run() (status Status) {
	if len(this.Jobs) == 0 {
		return true
	}

	if this.Stderr == nil {
		this.Stderr = os.Stderr
	}
	if this.NumCpu == 0 {
		this.NumCpu = runtime.NumCPU()
		if this.NumCpu == 1 {
			return this.run()
		}
		if this.NumCpu > 1 {
			this.NumCpu--
		}
	}
	return this.runP()
}
