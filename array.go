package gen

import "strconv"

var _ Code = (*ArrayRenderer)(nil)

type ArrayRenderer struct {
	size int
	elem Code
	ctx  Code
}

func Array(size int, elem Code) *ArrayRenderer {
	r := &ArrayRenderer{}
	r.SetSize(size)
	r.SetElem(elem)
	return r
}

func (r *ArrayRenderer) GetSize() int {
	return r.size
}

func (r *ArrayRenderer) SetSize(size int) {
	r.size = size
}

func (r *ArrayRenderer) GetElem() Code {
	return r.elem
}

func (r *ArrayRenderer) SetElem(elem Code) {
	r.elem = elem
	if elem != nil {
		elem.SetContext(r)
	}
}

func (r *ArrayRenderer) GetContext() Code {
	return r.ctx
}

func (r *ArrayRenderer) SetContext(ctx Code) {
	r.ctx = ctx
}

func (r *ArrayRenderer) Render(w Writer) {
	w.Write("[")
	w.Write(strconv.Itoa(r.size))
	w.Write("]")
	r.elem.Render(w)
}
