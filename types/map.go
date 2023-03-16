package types

import (
	gen "github.com/CherkashinEvgeny/gogen"
	"go/types"
)

func Map(t *types.Map) (code gen.Code) {
	return gen.Map(Type(t.Key()), Type(t.Elem()))
}
