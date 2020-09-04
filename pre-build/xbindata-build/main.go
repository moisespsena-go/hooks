package xbindata_build

import (
	"github.com/moisespsena-go/hooks"
)

func Run() error {
	return hooks.NewCmd("xbindata", "build", "--prod").Run()
}
