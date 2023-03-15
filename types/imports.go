package types

import (
	"github.com/CherkashinEvgeny/gogen"
	"github.com/pkg/errors"
	"go/types"
)

func TypeImports(t types.Type) []gen.Code {
	pkgs := map[*types.Package]struct{}{}
	addTypePackages(t, pkgs)
	imports := make([]gen.Code, 0, len(pkgs))
	for pkg := range pkgs {
		imports = append(imports, gen.Import(pkg.Name(), "", pkg.Path()))
	}
	return imports
}

func addTypePackages(t types.Type, imports map[*types.Package]struct{}) {
	switch v := t.(type) {
	case *types.Interface:
		addInterfacePackages(v, imports)
	case *types.Struct:
		addStructPackages(v, imports)
	case *types.Signature:
		addFuncPackages(v, imports)
	case *types.Tuple:
		addTuplePackages(v, imports)
	case *types.Map:
		addMapPackages(v, imports)
	case *types.Chan:
		addChanPackages(v, imports)
	case *types.Slice:
		addSlicePackages(v, imports)
	case *types.Array:
		addArrayPackages(v, imports)
	case *types.Pointer:
		addPointerPackages(v, imports)
	case *types.Named:
		addNamedPackages(v, imports)
	case *types.Basic:
		addBasicPackages(v, imports)
	default:
		panic(errors.Errorf("unknown type = %v", t))
	}
}

func InterfaceImports(t *types.Interface) []gen.Code {
	pkgs := map[*types.Package]struct{}{}
	addInterfacePackages(t, pkgs)
	imports := make([]gen.Code, 0, len(pkgs))
	for pkg := range pkgs {
		imports = append(imports, gen.Import(pkg.Name(), "", pkg.Path()))
	}
	return imports
}

func addInterfacePackages(t *types.Interface, imports map[*types.Package]struct{}) {
	n := t.NumMethods()
	for i := 0; i < n; i++ {
		method := t.Method(i)
		addTypePackages(method.Type(), imports)
	}
}

func StructureImports(t *types.Struct) []gen.Code {
	pkgs := map[*types.Package]struct{}{}
	addStructPackages(t, pkgs)
	imports := make([]gen.Code, 0, len(pkgs))
	for pkg := range pkgs {
		imports = append(imports, gen.Import(pkg.Name(), "", pkg.Path()))
	}
	return imports
}

func addStructPackages(t *types.Struct, imports map[*types.Package]struct{}) {
	n := t.NumFields()
	for i := 0; i < n; i++ {
		field := t.Field(i)
		addTypePackages(field.Type(), imports)
	}
}

func FuncImports(t *types.Signature) []gen.Code {
	pkgs := map[*types.Package]struct{}{}
	addFuncPackages(t, pkgs)
	imports := make([]gen.Code, 0, len(pkgs))
	for pkg := range pkgs {
		imports = append(imports, gen.Import(pkg.Name(), "", pkg.Path()))
	}
	return imports
}

func addFuncPackages(f *types.Signature, imports map[*types.Package]struct{}) {
	addTuplePackages(f.Params(), imports)
	addTuplePackages(f.Results(), imports)
}

func TupleImports(t *types.Tuple) []gen.Code {
	pkgs := map[*types.Package]struct{}{}
	addTuplePackages(t, pkgs)
	imports := make([]gen.Code, 0, len(pkgs))
	for pkg := range pkgs {
		imports = append(imports, gen.Import(pkg.Name(), "", pkg.Path()))
	}
	return imports
}

func addTuplePackages(t *types.Tuple, imports map[*types.Package]struct{}) {
	for i := 0; i < t.Len(); i++ {
		param := t.At(i)
		addTypePackages(param.Type(), imports)
	}
}

func MapImports(t *types.Map) []gen.Code {
	pkgs := map[*types.Package]struct{}{}
	addMapPackages(t, pkgs)
	imports := make([]gen.Code, 0, len(pkgs))
	for pkg := range pkgs {
		imports = append(imports, gen.Import(pkg.Name(), "", pkg.Path()))
	}
	return imports
}

func addMapPackages(t *types.Map, imports map[*types.Package]struct{}) {
	addTypePackages(t.Key(), imports)
	addTypePackages(t.Elem(), imports)
}

func ChanImports(t *types.Chan) []gen.Code {
	pkgs := map[*types.Package]struct{}{}
	addChanPackages(t, pkgs)
	imports := make([]gen.Code, 0, len(pkgs))
	for pkg := range pkgs {
		imports = append(imports, gen.Import(pkg.Name(), "", pkg.Path()))
	}
	return imports
}

func addChanPackages(t *types.Chan, imports map[*types.Package]struct{}) {
	addTypePackages(t.Elem(), imports)
}

func SliceImports(t *types.Slice) []gen.Code {
	pkgs := map[*types.Package]struct{}{}
	addSlicePackages(t, pkgs)
	imports := make([]gen.Code, 0, len(pkgs))
	for pkg := range pkgs {
		imports = append(imports, gen.Import(pkg.Name(), "", pkg.Path()))
	}
	return imports
}

func addSlicePackages(t *types.Slice, imports map[*types.Package]struct{}) {
	addTypePackages(t.Elem(), imports)
}

func ArrayImports(t *types.Array) []gen.Code {
	pkgs := map[*types.Package]struct{}{}
	addArrayPackages(t, pkgs)
	imports := make([]gen.Code, 0, len(pkgs))
	for pkg := range pkgs {
		imports = append(imports, gen.Import(pkg.Name(), "", pkg.Path()))
	}
	return imports
}

func addArrayPackages(t *types.Array, imports map[*types.Package]struct{}) {
	addTypePackages(t.Elem(), imports)
}

func PointerImports(t *types.Pointer) []gen.Code {
	pkgs := map[*types.Package]struct{}{}
	addTypePackages(t, pkgs)
	imports := make([]gen.Code, 0, len(pkgs))
	for pkg := range pkgs {
		imports = append(imports, gen.Import(pkg.Name(), "", pkg.Path()))
	}
	return imports
}

func addPointerPackages(t *types.Pointer, imports map[*types.Package]struct{}) {
	addTypePackages(t.Elem(), imports)
}

func NamedImports(t *types.Named) []gen.Code {
	pkgs := map[*types.Package]struct{}{}
	addNamedPackages(t, pkgs)
	imports := make([]gen.Code, 0, len(pkgs))
	for pkg := range pkgs {
		imports = append(imports, gen.Import(pkg.Name(), "", pkg.Path()))
	}
	return imports
}

func addNamedPackages(t *types.Named, imports map[*types.Package]struct{}) {
	obj := t.Obj()
	pkg := obj.Pkg()
	if pkg != nil {
		imports[pkg] = struct{}{}
	}
}

func addBasicPackages(_ *types.Basic, _ map[*types.Package]struct{}) {
}
