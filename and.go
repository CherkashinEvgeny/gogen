package gen

var _ Sequence = (*AndRenderer)(nil)

type AndRenderer struct {
	items []Code
	ctx   Code
}

func And(items ...Code) *AndRenderer {
	r := &AndRenderer{}
	r.Add(items...)
	return r
}

func (r *AndRenderer) Len() int {
	return len(r.items)
}

func (r *AndRenderer) At(i int) Code {
	return r.items[i]
}

func (r *AndRenderer) Add(items ...Code) {
	r.items = append(r.items, items...)
	for _, item := range items {
		item.SetContext(r)
	}
}

func (r *AndRenderer) GetContext() Code {
	return r.ctx
}

func (r *AndRenderer) SetContext(ctx Code) {
	r.ctx = ctx
	for _, item := range r.items {
		item.SetContext(r)
	}
}

func (r *AndRenderer) Render(w Writer) {
	for i, item := range r.items {
		if i != 0 {
			w.Write(" && ")
		}
		item.Render(w)
	}
}
