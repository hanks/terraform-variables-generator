package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/hanks/terraform-variables-generator/configs"
	"github.com/hanks/terraform-variables-generator/pkg/parser"
	"github.com/hanks/terraform-variables-generator/pkg/util"
)

func TestContainsElement(t *testing.T) {
	ter := &parser.TerraformVars{}
	testSlice := []parser.TFVar{parser.TFVar{Name: "Terraform"}, parser.TFVar{Name: "Puppet"}, parser.TFVar{Name: "Ansible"}}
	if ter.ContainsElement(testSlice, "Chef") {
		t.Error("Should return false, but return true")
	}
}

func TestGetAllFiles(t *testing.T) {
	files, _, err := util.GetAllFiles("", config.TFFileExt)
	util.CheckError(err)
	if len(files) == 0 {
		t.Error("Should found at least one file")
	}
}

func TestMatchVariable(t *testing.T) {
	ter := &parser.TerraformVars{}
	var messages []string

	file, _, err := util.GetAllFiles("", config.TFFileExt)
	util.CheckError(err)

	fileHandle, _ := os.Open(file[0])
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
		messages = append(messages, fileScanner.Text())
	}
	for _, text := range messages {
		ter.MatchVarPref(text, config.VarPrefix, config.Replacer)
	}
	if len(ter.Variables) != 6 {
		t.Errorf("Should return five variable. but returned %d", len(ter.Variables))
		t.Errorf("Variables found: %s", ter.Variables)
	}
}
