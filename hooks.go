package ents

import (
	"encoding/json"
	"log"
	"os"
	"path"

	"entgo.io/ent/entc/gen"
)

func (e *extension) generate(next gen.Generator) gen.Generator {
	return gen.GenerateFunc(func(g *gen.Graph) error {
		b, _ := json.Marshal(g.Schemas)
		os.WriteFile("debug.json", b, 0666)
		s := parseTemplate("types", g)
		err := os.WriteFile(path.Join(e.path, "ent.ts"), []byte(s), 0666)
		if err != nil {
			log.Fatalln(err)
		}
		return next.Generate(g)
	})
}
