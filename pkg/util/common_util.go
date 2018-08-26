package util

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

// CheckError helps to abort when error happens
func CheckError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// UserPromt helps to ask user to overwrite existed 'variables.tf' or not
func UserPromt(dstFile string) {
	var response string
	log.Warnf("File %q already exists, type yes if you want overwrite it", dstFile)
	fmt.Print("-> ")
	_, err := fmt.Scanln(&response)
	CheckError(err)
	if response != "yes" {
		os.Exit(0)
	}
}
