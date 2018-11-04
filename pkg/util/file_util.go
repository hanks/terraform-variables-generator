package util

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// GetAllFiles helps to get all blob type of files with specified extension in specified directory
func GetAllFiles(glob func(s string) ([]string, error), dir string, ext string) ([]string, string, error) {
	var err error
	var globPath string

	if dir == "" {
		dir, err = os.Getwd()
		globPath = ext
	} else {
		globPath = dir + "/" + ext
	}

	CheckError(err)
	var files []string
	log.Infof("Finding files in %q directory", dir)
	files, err = glob(globPath)
	CheckError(err)

	if len(files) == 0 {
		log.Infof("No files with %q extensions found in %q", ext, dir)
	}
	return files, dir, nil
}

// FileExists helps to judge file exists or not
func FileExists(name string) bool {
	if _, err := os.Stat(name); err == nil {
		return true
	}
	return false
}
