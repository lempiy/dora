package help

import (
	"log"
	"os"
	"path/filepath"
)

func GetCurrentDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Printf("GetCurrentDir: %s", err)
	}
	return dir
}
