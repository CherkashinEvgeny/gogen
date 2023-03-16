package types

import (
	gen "github.com/CherkashinEvgeny/gogen"
	"go/types"
)

func Pointer(t *types.Pointer) (code gen.Code) {
	return gen.Ptr(Type(t.Elem()))
}
