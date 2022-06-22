//go:build ignore

package main

// package ent

import (
	"log"

	"ariga.io/ogent"
	"entgo.io/contrib/entoas"
	"entgo.io/contrib/entproto"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
	"github.com/hedwigz/entviz"
	"github.com/ogen-go/ogen"
)

func main() {
	spec := new(ogen.Spec)
	oas, err := entoas.NewExtension(entoas.Spec(spec), entoas.Mutations(func(_ *gen.Graph, spec *ogen.Spec) error {
		spec.AddPathItem("/health", ogen.NewPathItem().SetDescription("return service healthy status").
			SetGet(ogen.NewOperation().SetOperationID("health").AddResponse("204", ogen.NewResponse())))
		return nil
	}), entoas.DefaultPolicy(entoas.PolicyExclude))

	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}
	ogent, err := ogent.NewExtension(spec)
	if err != nil {
		log.Fatalf("creating ogent extension: %v", err)
	}

	err = entc.Generate("./schema", &gen.Config{
		IDType: &field.TypeInfo{
			Type: field.TypeUUID,
		},
		Hooks: []gen.Hook{
			entproto.Hook(),
		},
	}, entc.Extensions(entviz.Extension{}, ogent, oas))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
