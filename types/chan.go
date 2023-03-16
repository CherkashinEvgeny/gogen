package types

import (
	gen "github.com/CherkashinEvgeny/gogen"
	"go/types"
)

func Chan(t *types.Chan) (code gen.Code) {
	switch t.Dir() {
	case types.SendOnly:
		return gen.Chan(gen.Send, Type(t.Elem()))
	case types.RecvOnly:
		return gen.Chan(gen.Receive, Type(t.Elem()))
	default:
		return gen.Chan(gen.SendAndReceive, Type(t.Elem()))
	}
}
