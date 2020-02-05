package xbindata_build

import (
	"github.com/ecletus/hooks"
)

func Run() error {
	return hooks.NewCmd("xbindata", "build").Run()
}
