package types

import (
	gen "github.com/CherkashinEvgeny/gogen"
	"go/types"
)

func Named(t *types.Named) (code gen.Code) {
	obj := t.Obj()
	pkg := obj.Pkg()
	if pkg == nil {
		return gen.Named("", obj.Name())
	}
	return gen.SmartNamed(pkg.Name(), pkg.Path(), obj.Name())
}
