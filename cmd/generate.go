package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/hashicorp/hcl/hcl/printer"
	log "github.com/sirupsen/logrus"

	"github.com/hanks/terraform-variables-generator/pkg/parser"
	"github.com/hanks/terraform-variables-generator/pkg/util"

	c "github.com/hanks/terraform-variables-generator/configs"
)

// Generate cmd to generate variables.tf from *.tf content
func Generate(dir string, varConfName string) {
	t, dirPath := gatherVars(dir, varConfName)
	output(t, dirPath)
}

func gatherVars(dir string, varConfName string) (*parser.TerraformVars, string) {
	if util.FileExists(c.DstFile) {
		util.UserPromt(fmt.Scanln, c.DstFile)
	}

	tfFiles, dirPath, err := util.GetAllFiles(filepath.Glob, dir, c.TFFileExt)
	if len(tfFiles) == 0 {
		log.Warn("No terraform files to proceed, exiting")
		os.Exit(0)
	}
	util.CheckError(err)

	var wg sync.WaitGroup
	messages := make(chan string)
	wg.Add(len(tfFiles))
	t := &parser.TerraformVars{}

	for _, file := range tfFiles {
		go func(file string) {
			defer wg.Done()
			fileHandle, _ := os.Open(file)
			defer fileHandle.Close()
			fileScanner := bufio.NewScanner(fileHandle)
			for fileScanner.Scan() {
				messages <- fileScanner.Text()
			}
		}(file)
	}
	go func() {
		for text := range messages {
			t.MatchVarPref(text, c.VarPrefix, c.Replacer)
		}
	}()
	wg.Wait()

	// merge customized vars in 'vars.yml' to t.Variables
	varConfPath := fmt.Sprintf("%s/%s", dirPath, varConfName)
	if util.FileExists(varConfPath) {
		varConfigs, err := parser.ParseCustVars(varConfPath)
		util.CheckError(err)

		err = t.MergeConfVars(&varConfigs)
		util.CheckError(err)
	}

	t.SortVars()

	return t, dirPath
}

func output(t *parser.TerraformVars, dirPath string) {
	destPath := dirPath + "/" + c.DstFile

	output, err := os.Create(destPath)
	defer output.Close()
	util.CheckError(err)

	err = c.VarTemplate.Execute(output, t.Variables)
	util.CheckError(err)

	// use terraform fmt to format again
	b, err := ioutil.ReadFile(destPath)
	util.CheckError(err)
	res, err := printer.Format(b)
	err = ioutil.WriteFile(destPath, res, 0644)
	util.CheckError(err)

	log.Infof("Variables are generated to %q file", destPath)
}
