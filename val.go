package gen

import (
	"strconv"
)

var _ Sequence = (*ValsRenderer)(nil)

type ValsRenderer struct {
	items []Code
	ctx   Code
}

func Vals(items ...Code) *ValsRenderer {
	r := &ValsRenderer{}
	r.Add(items...)
	return r
}

func (r *ValsRenderer) Len() int {
	return len(r.items)
}

func (r *ValsRenderer) At(i int) Code {
	return r.items[i]
}

func (r *ValsRenderer) Add(items ...Code) {
	r.items = append(r.items, items...)
	for _, item := range items {
		item.SetContext(r)
	}
}

func (r *ValsRenderer) GetContext() Code {
	return r.ctx
}

func (r *ValsRenderer) SetContext(ctx Code) {
	r.ctx = ctx
}

func (r *ValsRenderer) Render(w Writer) {
	for i, item := range r.items {
		if i != 0 {
			w.Write(", ")
		}
		item.Render(w)
	}
}

var _ Code = (*ValRenderer)(nil)

type ValRenderer struct {
	value any
}

func Val(value any) ValRenderer {
	return ValRenderer{value}
}

func (r ValRenderer) GetContext() Code {
	return nil
}

func (r ValRenderer) SetContext(ctx Code) {
}

func (r ValRenderer) Render(w Writer) {
	switch v := r.value.(type) {
	case bool:
		renderBoolLiteral(v, w)
	case int:
		renderIntLiteral(v, w)
	case int8:
		renderInt8Literal(v, w)
	case int16:
		renderInt16Literal(v, w)
	case int32:
		renderInt32Literal(v, w)
	case int64:
		renderInt64Literal(v, w)
	case uint:
		renderUintLiteral(v, w)
	case uint8:
		renderUint8Literal(v, w)
	case uint16:
		renderUint16Literal(v, w)
	case uint32:
		renderUint32Literal(v, w)
	case uint64:
		renderUint64Literal(v, w)
	case float32:
		renderFloat32Literal(v, w)
	case float64:
		renderFloat64Literal(v, w)
	case string:
		renderStringLiteral(v, w)
	}
}

func renderBoolLiteral(v bool, w Writer) {
	w.Write(strconv.FormatBool(v))
}

func renderIntLiteral(v int, w Writer) {
	w.Write(strconv.Itoa(v))
}

func renderInt8Literal(v int8, w Writer) {
	w.Write(strconv.FormatInt(int64(v), 10))
}

func renderInt16Literal(v int16, w Writer) {
	w.Write(strconv.FormatInt(int64(v), 10))
}

func renderInt32Literal(v int32, w Writer) {
	w.Write(strconv.FormatInt(int64(v), 10))
}

func renderInt64Literal(v int64, w Writer) {
	w.Write(strconv.FormatInt(v, 10))
}

func renderUintLiteral(v uint, w Writer) {
	w.Write(strconv.FormatUint(uint64(v), 10))
}

func renderUint8Literal(v uint8, w Writer) {
	w.Write(strconv.FormatUint(uint64(v), 10))
}

func renderUint16Literal(v uint16, w Writer) {
	w.Write(strconv.FormatUint(uint64(v), 10))
}

func renderUint32Literal(v uint32, w Writer) {
	w.Write(strconv.FormatUint(uint64(v), 10))
}

func renderUint64Literal(v uint64, w Writer) {
	w.Write(strconv.FormatUint(v, 10))
}

func renderFloat32Literal(v float32, w Writer) {
	w.Write(strconv.FormatFloat(float64(v), 'f', -1, 10))
}

func renderFloat64Literal(v float64, w Writer) {
	w.Write(strconv.FormatFloat(v, 'f', -1, 10))
}

func renderStringLiteral(str string, w Writer) {
	w.Write("\"")
	for _, c := range str {
		if c == '"' || c == '\\' {
			w.Write("\\")
		}
		w.Write(string(c))
	}
	w.Write("\"")
}
