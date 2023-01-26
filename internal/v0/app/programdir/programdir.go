// модуль для определения главного каталога программы
package programdir

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/common/pkg/v0/micro"
)

// CurrentFilename - возвращает полное имя текущего исполняемого файла
func CurrentFilename() string {
	_, filename, _, _ := runtime.Caller(0)
	return filename
}

// ProgramDir - возвращает главный каталог программы, в конце "/"
func ProgramDir() string {
	filename := os.Args[0]
	dir := filepath.Dir(filename)
	sdir := strings.ToLower(dir)

	if micro.SubstringLeft(sdir, 5) == "/tmp/" {
		filename = CurrentFilename()
		dir = filepath.Dir(filename)

		if dir[len(dir)-10:] == "programdir" {
			dir = micro.FindDirUp(dir)
			dir = micro.FindDirUp(dir)
			dir = micro.FindDirUp(dir)
			dir = micro.FindDirUp(dir)
		}
	}

	dir = micro.AddSeparator(dir)
	return dir
}
