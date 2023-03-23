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

type Comparable interface{ ~string | ~int | ~float32 }
