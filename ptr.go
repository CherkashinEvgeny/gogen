package gen

var _ Code = (*PtrRenderer)(nil)

type PtrRenderer struct {
	elem Code
	ctx  Code
}

func Ptr(elem Code) *PtrRenderer {
	r := &PtrRenderer{}
	r.SetElem(elem)
	return r
}

func (r *PtrRenderer) GetElem() Code {
	return r.elem
}

func (r *PtrRenderer) SetElem(elem Code) {
	r.elem = elem
	if elem != nil {
		elem.SetContext(r)
	}
}

func (r *PtrRenderer) GetContext() Code {
	return r.ctx
}

func (r *PtrRenderer) SetContext(ctx Code) {
	r.ctx = ctx
}

func (r *PtrRenderer) Render(w Writer) {
	w.Write("*")
	r.elem.Render(w)
}
