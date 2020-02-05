package hooks

import (
	"reflect"
	"runtime"
)

type PostJob interface {
	Run(binNames ...string) error
}

type Job interface {
	Run() error
}

type JobNamer interface {
	JobName() string
}

type JobFunc func() error

func (this JobFunc) Run() error {
	return this()
}

func (this JobFunc) JobName() string {
	return runtime.FuncForPC(reflect.ValueOf(this).Pointer()).Name()
}

type NamedJob struct {
	Job
	Name string
}

func (this NamedJob) JobName() string {
	return this.Name
}

type PostJobFunc func(binNames ...string) error

func (this PostJobFunc) Run(binNames ...string) error {
	return this(binNames...)
}

func (this PostJobFunc) JobName() string {
	return JobName(this)
}

type NamedPostJob struct {
	PostJob
	Name string
}

func (this NamedPostJob) JobName() string {
	return this.Name
}

func JobName(j interface{}) string {
	v := reflect.ValueOf(j)
	loop:
	for {
		switch v.Kind() {
		case reflect.Struct, reflect.Func:
			break loop
		default:
			v = v.Elem()
		}
	}
	switch v.Kind() {
	case reflect.Struct:
		return v.Type().PkgPath() + "." + v.Type().Name()
	case reflect.Func:
		return runtime.FuncForPC(v.Pointer()).Name()
	default:
		return "<no name>"
	}
}
