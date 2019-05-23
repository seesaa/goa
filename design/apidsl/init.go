package apidsl

import (
	"github.com/seesaa/goa/design"
	"github.com/seesaa/goa/dslengine"
)

// Setup API DSL roots.
func init() {
	design.Design = design.NewAPIDefinition()
	design.GeneratedMediaTypes = make(design.MediaTypeRoot)
	design.ProjectedMediaTypes = make(design.MediaTypeRoot)
	dslengine.Register(design.Design)
	dslengine.Register(design.GeneratedMediaTypes)
}
