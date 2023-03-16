package gen

var _ Code = (*AddrRenderer)(nil)

type AddrRenderer struct {
	item Code
	ctx  Code
}

func Addr(elem Code) *AddrRenderer {
	r := &AddrRenderer{}
	r.SetItem(elem)
	return r
}

func (r *AddrRenderer) GetItem() Code {
	return r.item
}

func (r *AddrRenderer) SetItem(item Code) {
	r.item = item
	if item != nil {
		item.SetContext(r)
	}
}

func (r *AddrRenderer) GetContext() Code {
	return r.ctx
}

func (r *AddrRenderer) SetContext(ctx Code) {
	r.ctx = ctx
	if r.item != nil {
		r.item.SetContext(r)
	}
}

func (r *AddrRenderer) Render(w Writer) {
	w.Write("&")
	r.item.Render(w)
}
