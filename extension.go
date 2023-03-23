package ents

import (
	"embed"

	"entgo.io/ent/entc/gen"
)

//go:embed templates
var templates embed.FS

func NewExtension(path string) *extension {
	e := new(extension)
	e.path = path
	e.hooks = append(e.hooks, e.generate)
	return e
}

func (e *extension) Hooks() []gen.Hook {
	return e.hooks
}
