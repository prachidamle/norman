package generator

var schemeTemplate = `package {{.version.Version}}

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	GroupName = "{{.version.Group}}"
	Version = "{{.version.Version}}"
)

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: Version}

// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

{{- if not .external }}
var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	// TODO this gets cleaned up when the types are fixed
	scheme.AddKnownTypes(SchemeGroupVersion,
	{{range .names}}
	&{{.}}{},{{end}}
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
{{- end }}
`
