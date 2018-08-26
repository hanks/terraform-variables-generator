package parser

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// CustVar is the struct of customized variable
type CustVar struct {
	Type    string `yaml:"type"`
	Desc    string `yaml:"description"`
	Default string `yaml:"default"`
}

// CustVars is the struct contains list of CustVar
type CustVars struct {
	Vars []map[string]CustVar `yaml:"vars"`
}

// ParseCustVars is used to get the content from customized variable config file
func ParseCustVars(path string) (CustVars, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return CustVars{}, err
	}

	custVars := CustVars{}
	err = yaml.Unmarshal(buf, &custVars)
	if err != nil {
		return CustVars{}, err
	}

	return custVars, nil
}
