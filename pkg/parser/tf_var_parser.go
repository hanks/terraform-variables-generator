package parser

import (
	"regexp"
	"sort"
	"strings"
)

// TFVar represents struct of 'variable block' in 'variables.tf'
type TFVar struct {
	Name    string
	Type    string
	Desc    string
	Default string
}

// TerraformVars contains list of TFVar
type TerraformVars struct {
	Variables []TFVar
}

// MatchVarPref to fetch all 'var name' defined in *.tf files
func (t *TerraformVars) MatchVarPref(row string, varPrefix string, replacer *strings.Replacer) {
	if strings.Contains(row, varPrefix) {
		pattern := regexp.MustCompile(`var.([a-z?A-Z?0-9?_][a-z?A-Z?0-9?_?-]*)`)
		match := pattern.FindAllStringSubmatch(row, -1)
		for _, m := range match {
			res := replacer.Replace(m[0])
			if !t.ContainsElement(t.Variables, res) {
				tfvar := TFVar{
					Name: res,
				}
				t.Variables = append(t.Variables, tfvar)
			}
		}
	}
}

// SortVars to sort 'TFVar' by its 'Name' field in a alphabetical order
func (t *TerraformVars) SortVars() {
	sort.Slice(t.Variables, func(i, j int) bool {
		return t.Variables[i].Name < t.Variables[j].Name
	})
}

// ContainsElement is a helper to deduplicate var names in *.tf files
func (t *TerraformVars) ContainsElement(slice []TFVar, value string) bool {
	if len(slice) == 0 {
		return false
	}
	for _, tfvar := range slice {
		if value == tfvar.Name {
			return true
		}
	}
	return false
}

// MergeConfVars is used to merge var config between *.tf and customized var conf
func (t *TerraformVars) MergeConfVars(custVars *CustVars) error {
	if len(custVars.Vars) == 0 {
		return nil
	}

	mapping := make(map[string]CustVar)
	for _, m := range custVars.Vars {
		for name, varconf := range m {
			mapping[name] = varconf
		}
	}

	for i, tfvar := range t.Variables {
		if _, exists := mapping[tfvar.Name]; exists {
			t.Variables[i].Type = mapping[tfvar.Name].Type
			t.Variables[i].Desc = mapping[tfvar.Name].Desc
			t.Variables[i].Default = mapping[tfvar.Name].Default
		}
	}

	return nil
}
