package config

import (
	"strings"
	"text/template"
)

// VarPrefix is the prefix of var name
var VarPrefix = "var."

// TFFileExt is the extension name of target terraform config file
var TFFileExt = "*.tf"

// VarConfName is the name of customized var config filename
var VarConfName = "vars.yml"

// DstFile is the name of target variables.tf filename
var DstFile = "variables.tf"

// VarTemplate is the template object to define 'variable block' in 'variables.tf'
var VarTemplate = template.Must(template.New("var_file").Parse(`{{ range . }}
variable "{{ .Name }}" {
  {{ if .Desc -}}
  description = "{{ .Desc }}"
  {{ end }}
  {{- if .Type -}}
  type = "{{ .Type }}"
  {{ end }}
  {{- if .Default }}
  default = {{ .Default }}
  {{ end }}
}
{{ end }}`))

// Replacer is used to simple clean up var name strings found in *.tf
var Replacer = strings.NewReplacer(":", ".",
	"]", "",
	"}", "",
	"{", "",
	"\"", "",
	")", "",
	"(", "",
	"[", "",
	",", "",
	"var.", "",
	" ", "",
)