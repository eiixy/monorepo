//go:build ignore
// +build ignore

package main

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/eiixy/monorepo/internal/data/annotations"
	"log"
	"runtime"
	"strings"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithConfigPath("./gqlgen.yml"),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("./ent.graphql"),
		entgql.WithWhereInputs(true),
		entgql.WithNodeDescriptor(true),
		entgql.WithSchemaHook(annotations.EnumsGQLSchemaHook),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	_, filename, _, _ := runtime.Caller(0)
	entPath := strings.TrimSuffix(filename, "account/ent/entc.go")
	if err = entc.Generate(entPath+"account/ent/schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeatureIntercept,
			gen.FeatureSnapshot,
			gen.FeatureModifier,
			gen.FeatureExecQuery,
			gen.FeatureEntQL,
		},
	}, entc.Extensions(ex), entc.TemplateDir(entPath+"template")); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
