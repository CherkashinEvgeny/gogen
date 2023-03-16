package types

import (
	gen "github.com/CherkashinEvgeny/gogen"
	"go/types"
)

func Slice(t *types.Slice) (code gen.Code) {
	return gen.Slice(Type(t.Elem()))
}
