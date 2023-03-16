package gen

var _ Code = (*IncRenderer)(nil)

type IncRenderer struct {
	item Code
	ctx  Code
}

func Inc(vvar Code) *IncRenderer {
	r := &IncRenderer{}
	r.SetItem(vvar)
	return r
}

func (r *IncRenderer) GetItem() Code {
	return r.item
}

func (r *IncRenderer) SetItem(item Code) {
	r.item = item
	if item != nil {
		item.SetContext(r)
	}
}

func (r *IncRenderer) GetContext() Code {
	return r.ctx
}

func (r *IncRenderer) SetContext(ctx Code) {
	r.ctx = ctx
	if r.item != nil {
		r.item.SetContext(r)
	}
}

func (r *IncRenderer) Render(w Writer) {
	r.item.Render(w)
	w.Write("++")
}
