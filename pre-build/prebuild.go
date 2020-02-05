package pre_build

import (
	"github.com/ecletus/hooks"
	fix_jinzhu_configor_print_to_stdout "github.com/ecletus/hooks/pre-build/fix-jinzhu-configor-print-to-stdout"
	xbindata_build "github.com/ecletus/hooks/pre-build/xbindata-build"
)

func Hooks() hooks.Jobs {
	return hooks.Jobs{
		hooks.JobFunc(fix_jinzhu_configor_print_to_stdout.Run),
		hooks.JobFunc(xbindata_build.Run),
	}
}
