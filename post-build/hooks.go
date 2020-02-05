package post_build

import (
	"github.com/moisespsena-go/hooks"
	xbindata_build_program "github.com/moisespsena-go/hooks/post-build/xbindata-build-program"
)

func Hooks() (jobs hooks.PostJobs) {
	return hooks.PostJobs{
		hooks.PostJobFunc(xbindata_build_program.Run),
	}
}
