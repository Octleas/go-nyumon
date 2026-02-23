package img

import (
	"embed"
	"io/fs"
)

//go:embed dog.txt
var img embed.FS

func Open(name string) (fs.File, error) {
	return img.Open(name)
}