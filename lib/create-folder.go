package lib

import (
	"os"
)

func CreateFolder(path string) error {
	return os.Mkdir(path, 0755)
}
