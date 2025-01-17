package ccnop

import (
	"text/template"

	"github.com/nvx/protoc-gen-validate/templates/cc"
	pgs "github.com/lyft/protoc-gen-star"
)

func RegisterModule(tpl *template.Template, params pgs.Parameters) {
	cc.RegisterModule(tpl, params)
	template.Must(tpl.Parse(moduleFileTpl))
}

func RegisterHeader(tpl *template.Template, params pgs.Parameters) {
	cc.RegisterHeader(tpl, params)
	template.Must(tpl.Parse(headerFileTpl))
}
