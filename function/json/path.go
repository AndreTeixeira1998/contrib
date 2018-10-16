package json

import (
	"github.com/oliveagle/jsonpath"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnPath{})
}

type fnPath struct {
}

// Name returns the name of the function
func (fnPath) Name() string {
	return "json.path"
}

// Sig returns the function signature
func (fnPath) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeAny}, false
}

// Eval executes the function
func (fnPath) Eval(params ...interface{}) (interface{}, error) {
	expression := params[0].(string)
	return jsonpath.JsonPathLookup(params[1], expression)
}
