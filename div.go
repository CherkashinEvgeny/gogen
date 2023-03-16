package gen

var _ Sequence = (*DivRenderer)(nil)

type DivRenderer struct {
	items []Code
	ctx   Code
}

func Div(items ...Code) *DivRenderer {
	r := &DivRenderer{}
	r.Add(items...)
	return r
}

func (r *DivRenderer) Len() int {
	return len(r.items)
}

func (r *DivRenderer) At(i int) Code {
	return r.items[i]
}

func (r *DivRenderer) Add(items ...Code) {
	r.items = append(r.items, items...)
	for _, item := range items {
		item.SetContext(r)
	}
}

func (r *DivRenderer) GetContext() Code {
	return r.ctx
}

func (r *DivRenderer) SetContext(ctx Code) {
	r.ctx = ctx
	for _, item := range r.items {
		item.SetContext(r)
	}
}

func (r *DivRenderer) Render(w Writer) {
	for i, item := range r.items {
		if i != 0 {
			w.Write(" / ")
		}
		item.Render(w)
	}
}
