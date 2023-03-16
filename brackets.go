package gen

var _ Code = (*BracketsRenderer)(nil)

type BracketsRenderer struct {
	item Code
	ctx  Code
}

func Brackets(item Code) *BracketsRenderer {
	r := &BracketsRenderer{}
	r.SetItem(item)
	return r
}

func (r *BracketsRenderer) GetItem() Code {
	return r.item
}

func (r *BracketsRenderer) SetItem(item Code) {
	r.item = item
	if item != nil {
		item.SetContext(r)
	}
}

func (r *BracketsRenderer) GetContext() Code {
	return r.ctx
}

func (r *BracketsRenderer) SetContext(ctx Code) {
	r.ctx = ctx
	if r.item != nil {
		r.item.SetContext(r)
	}
}

func (r *BracketsRenderer) Render(w Writer) {
	w.Write("(")
	r.item.Render(w)
	w.Write(")")
}
