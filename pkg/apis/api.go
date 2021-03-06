// +nirvana:api=descriptors:"Descriptor"

package apis

import (
	"github.com/quasimodo7614/nirvanatest/pkg/apis/middlewares"

	def "github.com/caicloud/nirvana/definition"

	v1 "github.com/quasimodo7614/nirvanatest/pkg/apis/v1/descriptors"
)

// Descriptor returns a combined descriptor for APIs of all versions.
func Descriptor() def.Descriptor {
	return def.Descriptor{
		Description: "APIs",
		Path:        "/apis",
		Middlewares: middlewares.Middlewares(),
		Consumes:    []string{def.MIMEJSON},
		Produces:    []string{def.MIMEJSON},
		Children: []def.Descriptor{
			v1.Descriptor(),
		},
	}
}
