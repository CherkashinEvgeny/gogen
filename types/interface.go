package types

import (
	gen "github.com/CherkashinEvgeny/gogen"
	"go/types"
)

func Interface(t *types.Interface) (code gen.Code) {
	methods := gen.Methods()
	n := t.NumMethods()
	for i := 0; i < n; i++ {
		method := t.Method(i)
		methodType := method.Type()
		nd, ok := methodType.(*types.Named)
		if !ok {
			continue
		}
		methods.Add(gen.Embedded(gen.Qual(
			nd.Obj().Pkg().Path(),
			nd.Obj().Name(),
		)))
	}
	for i := 0; i < n; i++ {
		method := t.Method(i)
		methodType := method.Type()
		s, ok := methodType.(*types.Signature)
		if !ok {
			continue
		}
		methods.Add(gen.MethodDecl(method.Name(), gen.Sign(
			funcIn(s),
			funcOut(s),
		)))
	}
	return gen.Iface(methods)
}
