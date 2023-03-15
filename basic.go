package gen

const (
	Bool       = basicRenderer("bool")
	Int        = basicRenderer("int")
	Int8       = basicRenderer("int8")
	Int16      = basicRenderer("int16")
	Int32      = basicRenderer("int32")
	Int64      = basicRenderer("int64")
	Uint       = basicRenderer("uint")
	Uint8      = basicRenderer("uint8")
	Uint16     = basicRenderer("uint16")
	Uint32     = basicRenderer("uint32")
	Uint64     = basicRenderer("uint64")
	Uintptr    = basicRenderer("uintptr")
	Float32    = basicRenderer("float32")
	Float64    = basicRenderer("float64")
	Complex64  = basicRenderer("complex64")
	Complex128 = basicRenderer("complex128")
	String     = basicRenderer("string")
	Byte       = basicRenderer("byte")
	Rune       = basicRenderer("rune")
	Any        = basicRenderer("any")
	Error      = basicRenderer("error")
)

var _ Code = basicRenderer("")

type basicRenderer string

func (r basicRenderer) GetContext() Code {
	return nil
}

func (r basicRenderer) SetContext(_ Code) {
}

func (r basicRenderer) Render(w Writer) {
	w.Write(string(r))
}
