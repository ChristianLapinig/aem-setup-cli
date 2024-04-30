package lib

import (
	"github.com/otiai10/copy"
)

func Copy(src, dest string) error {
	return copy.Copy(src, dest)
}
