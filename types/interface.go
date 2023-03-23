package types

import (
	"go/types"

	gen "github.com/CherkashinEvgeny/gogen"
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

func ForEachInterfaceMethod(iface *types.Interface, f func(name string, sign *types.Signature)) {
	n := iface.NumEmbeddeds()
	for i := 0; i < n; i++ {
		embedded := iface.EmbeddedType(i)
		underlying := embedded.Underlying()
		embeddedIface, ok := underlying.(*types.Interface)
		if !ok {
			continue
		}
		ForEachInterfaceMethod(embeddedIface, f)
	}
	n = iface.NumMethods()
	for i := 0; i < n; i++ {
		method := iface.Method(i)
		methodType := method.Type()
		sign, ok := methodType.(*types.Signature)
		if !ok {
			continue
		}
		f(method.Name(), sign)
	}
}
