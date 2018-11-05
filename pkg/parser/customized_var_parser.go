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

// equal does a simple equal compare between two CustVars structs
func (c CustVars) equal(other CustVars) bool {
	if len(c.Vars) != len(other.Vars) {
		return false
	}

	for idx, m := range c.Vars {
		for k := range m {
			if _, ok := other.Vars[idx][k]; !ok {
				return false
			}
		}
	}

	return true
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
