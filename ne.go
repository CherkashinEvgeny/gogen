package gen

var _ Code = (*NeRenderer)(nil)

type NeRenderer struct {
	item Code
	ctx  Code
}

func Ne(item Code) *NeRenderer {
	r := &NeRenderer{}
	r.SetItem(item)
	return r
}

func (r *NeRenderer) GetItem() Code {
	return r.item
}

func (r *NeRenderer) SetItem(item Code) {
	r.item = item
	if item != nil {
		item.SetContext(r)
	}
}

func (r *NeRenderer) GetContext() Code {
	return r.ctx
}

func (r *NeRenderer) SetContext(ctx Code) {
	r.ctx = ctx
}

func (r *NeRenderer) Render(w Writer) {
	w.Write("!")
	r.item.Render(w)
}
