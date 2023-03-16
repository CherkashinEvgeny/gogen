package types

import (
	gen "github.com/CherkashinEvgeny/gogen"
	"go/types"
)

func Array(t *types.Array) (code gen.Code) {
	return gen.Array(int(t.Len()), Type(t.Elem()))
}
