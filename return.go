package gen

var _ Code = (*ReturnRenderer)(nil)

type ReturnRenderer struct {
	item Code
	ctx  Code
}

func Return(elem Code) *ReturnRenderer {
	r := &ReturnRenderer{}
	r.SetItem(elem)
	return r
}

func (r *ReturnRenderer) GetItem() Code {
	return r.item
}

func (r *ReturnRenderer) SetItem(elem Code) {
	r.item = elem
	if elem != nil {
		elem.SetContext(r)
	}
}

func (r *ReturnRenderer) GetContext() Code {
	return r.ctx
}

func (r *ReturnRenderer) SetContext(ctx Code) {
	r.ctx = ctx
	if r.item != nil {
		r.item.SetContext(r)
	}
}

func (r *ReturnRenderer) Render(w Writer) {
	w.Write("return")
	if r.item == nil {
		return
	}
	if IsEmpty(r.item) {
		return
	}
	w.Write(" ")
	r.item.Render(w)
}
