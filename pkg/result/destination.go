package result

import (
	"os"
	"path/filepath"
	"strings"
)

func CreateDestination(destination, ext string) (*os.File, error) {
	if strings.Contains(ext, ".") {
		destination += ext
	} else {
		destination += "." + ext
	}
	return os.Create(filepath.Join(destination))
}

func Write(file *os.File, elem ...string) error {
	for _, s := range elem {
		if _, err := file.WriteString(s); err != nil {
			return err
		}
	}

	return nil
}
