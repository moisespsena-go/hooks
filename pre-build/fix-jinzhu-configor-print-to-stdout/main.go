package fix_jinzhu_configor_print_to_stdout

import (
	"io/ioutil"
	"os"
	"regexp"

	path_helpers "github.com/moisespsena-go/path-helpers"
)

func Run() (err error) {
	_, pth := path_helpers.ResolveGoSrcPath("github.com/jinzhu/configor/utils.go")
	s, err := os.Stat(pth)
	if err != nil {
		return err
	}
	f, err := os.Open(pth)
	if err != nil {
		return err
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	data = regexp.MustCompilePOSIX(`fmt.Printf\(`).ReplaceAll(data, []byte(`fmt.Fprintf(os.Stderr, `))
	return ioutil.WriteFile(pth, data, s.Mode())
}
