package gen

var _ Code = (*SliceRenderer)(nil)

type SliceRenderer struct {
	ctx  Code
	elem Code
}

func Slice(elem Code) *SliceRenderer {
	s := &SliceRenderer{}
	s.SetElem(elem)
	return s
}

func (r *SliceRenderer) GetElem() Code {
	return r.elem
}

func (r *SliceRenderer) SetElem(elem Code) {
	r.elem = elem
	if elem != nil {
		elem.SetContext(r)
	}
}

func (r *SliceRenderer) GetContext() Code {
	return r.ctx
}

func (r *SliceRenderer) SetContext(ctx Code) {
	r.ctx = ctx
	if r.elem != nil {
		r.elem.SetContext(r)
	}
}

func (r *SliceRenderer) Render(w Writer) {
	w.Write("[]")
	r.elem.Render(w)
}
