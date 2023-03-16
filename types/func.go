package types

import (
	gen "github.com/CherkashinEvgeny/gogen"
	"go/types"
)

func Func(t *types.Signature) (code gen.Code) {
	return gen.FuncType(gen.Sign(funcIn(t), funcOut(t)))
}

func funcIn(t *types.Signature) gen.Code {
	in := gen.In()
	params := t.Params()
	n := params.Len()
	if t.Variadic() {
		n--
	}
	for i := 0; i < n; i++ {
		param := params.At(i)
		in.Add(gen.Param(param.Name(), Type(param.Type()), false))
	}
	if t.Variadic() {
		param := params.At(n)
		in.Add(gen.Param(param.Name(), Type(param.Type()), true))
	}
	return in
}

func funcOut(t *types.Signature) gen.Code {
	out := gen.Out()
	params := t.Results()
	n := params.Len()
	for i := 0; i < n; i++ {
		variable := params.At(i)
		out.Add(gen.Param(variable.Name(), Type(variable.Type()), false))
	}
	return out
}

func Tuple(t *types.Tuple) (code gen.Code) {
	n := t.Len()
	params := make([]gen.Code, 0, n)
	for i := 0; i < n; i++ {
		param := t.At(i)
		params = append(params, Type(param.Type()))
	}
	return gen.Params(params...)
}
