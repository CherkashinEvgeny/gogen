package types

import (
	"github.com/CherkashinEvgeny/gogen"
	"github.com/pkg/errors"
	"go/types"
)

func Type(t types.Type) (code gen.Code) {
	switch v := t.(type) {
	case *types.Interface:
		return Interface(v)
	case *types.Struct:
		return Struct(v)
	case *types.Signature:
		return Func(v)
	case *types.Tuple:
		return Tuple(v)
	case *types.Map:
		return Map(v)
	case *types.Chan:
		return Chan(v)
	case *types.Slice:
		return Slice(v)
	case *types.Array:
		return Array(v)
	case *types.Pointer:
		return Pointer(v)
	case *types.Named:
		return Named(v)
	case *types.Basic:
		return Basic(v)
	default:
		panic(errors.Errorf("unknown type = %v", t))
	}
}
