package ents

import (
	"log"
	"os"
	"path"

	"entgo.io/ent/entc/gen"
)

func (e *extension) generate(next gen.Generator) gen.Generator {
	return gen.GenerateFunc(func(g *gen.Graph) error {
		s := parseTemplate("types", g)
		err := os.WriteFile(path.Join(e.path, "ent.ts"), []byte(s), 0666)
		if err != nil {
			log.Fatalln(err)
		}
		return next.Generate(g)
	})
}
