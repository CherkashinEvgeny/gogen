package types

import (
	gen "github.com/CherkashinEvgeny/gogen"
	"go/types"
)

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
