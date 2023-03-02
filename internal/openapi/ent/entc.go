//go:build ignore
// +build ignore

package main

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/chestarss/elk"
	"github.com/chestarss/elk/spec"
	"log"
)

func main() {
	ex, err := elk.NewExtension(
		elk.GenerateSpec(
			"openapi.json",
			elk.SpecTitle("My Pets API"),
			elk.SpecDescription("Awesome, Mega Cool API to manage Ariel's Pet Leopards!"),
			elk.SpecVersion("0.0.1"),
			elk.SpecSecuritySchemes(map[string]spec.SecurityScheme{
				"apiKeySample": {
					Type: "apiKey",
					In:   "header",
					Name: "X-API-KEY",
				},
			}),
			elk.SpecSecurity([]map[string][]string{
				{"apiKeySample": {}},
			}),
		),
	)
	if err != nil {
		log.Fatalf("creating elk extension: %v", err)
	}
	err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(ex))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
