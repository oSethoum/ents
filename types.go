package ents

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

type extension struct {
	entc.DefaultExtension
	hooks []gen.Hook
	path  string
}

var gots = map[string]string{
	"time.Time": "string",
	"bool":      "boolean",
	"int":       "number",
	"uint":      "number",
	"float":     "number",
	"enum":      "string",
	"any":       "any",
	"other":     "any",
	"json":      "any",
}
