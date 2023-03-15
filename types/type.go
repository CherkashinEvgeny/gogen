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
		methods.Add(gen.Embedded(gen.Named(
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

func Struct(t *types.Struct) (code gen.Code) {
	fields := gen.FieldDecls()
	n := t.NumFields()
	for i := 0; i < n; i++ {
		field := t.Field(i)
		fields.Add(gen.FieldDecl(field.Name(), Type(field.Type())))
	}
	return gen.Struct(fields)
}

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

func Map(t *types.Map) (code gen.Code) {
	return gen.Map(Type(t.Key()), Type(t.Elem()))
}

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

func Slice(t *types.Slice) (code gen.Code) {
	return gen.Slice(Type(t.Elem()))
}

func Array(t *types.Array) (code gen.Code) {
	return gen.Array(int(t.Len()), Type(t.Elem()))
}

func Pointer(t *types.Pointer) (code gen.Code) {
	return gen.Ptr(Type(t.Elem()))
}

func Named(t *types.Named) (code gen.Code) {
	var path string
	var name string
	obj := t.Obj()
	pkg := obj.Pkg()
	if pkg != nil {
		path = pkg.Path()
	}
	name = obj.Name()
	return gen.Named(path, name)
}

func Basic(t *types.Basic) (code gen.Code) {
	switch t.String() {
	case "bool":
		return gen.Bool
	case "int":
		return gen.Int
	case "int8":
		return gen.Int8
	case "int16":
		return gen.Int16
	case "int32":
		return gen.Int32
	case "int64":
		return gen.Int64
	case "uint":
		return gen.Uint
	case "uint8":
		return gen.Uint8
	case "uint16":
		return gen.Uint16
	case "uint32":
		return gen.Uint32
	case "uint64":
		return gen.Uint64
	case "uintptr":
		return gen.Uintptr
	case "float32":
		return gen.Float32
	case "float64":
		return gen.Float64
	case "complex64":
		return gen.Complex64
	case "complex128":
		return gen.Complex128
	case "string":
		return gen.String
	case "byte":
		return gen.Byte
	case "rune":
		return gen.Rune
	default:
		panic("unsupported type")
	}
}
