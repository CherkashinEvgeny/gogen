package types

import (
	gen "github.com/CherkashinEvgeny/gogen"
	"go/types"
)

func Struct(t *types.Struct) (code gen.Code) {
	fields := gen.FieldDecls()
	n := t.NumFields()
	for i := 0; i < n; i++ {
		field := t.Field(i)
		fields.Add(gen.FieldDecl(field.Name(), Type(field.Type())))
	}
	return gen.Struct(fields)
}
