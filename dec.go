package gen

var _ Code = (*DecRenderer)(nil)

type DecRenderer struct {
	item Code
	ctx  Code
}

func Dec(item Code) *DecRenderer {
	r := &DecRenderer{}
	r.SetItem(item)
	return r
}

func (r *DecRenderer) GetItem() Code {
	return r.item
}

func (r *DecRenderer) SetItem(item Code) {
	r.item = item
	if item != nil {
		item.SetContext(r)
	}
}

func (r *DecRenderer) GetContext() Code {
	return r.ctx
}

func (r *DecRenderer) SetContext(ctx Code) {
	r.ctx = ctx
}

func (r *DecRenderer) Render(w Writer) {
	r.item.Render(w)
	w.Write("--")
}
