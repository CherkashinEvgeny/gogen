package gen

var _ Sequence = (*OrRenderer)(nil)

type OrRenderer struct {
	items []Code
	ctx   Code
}

func Or(items ...Code) *OrRenderer {
	r := &OrRenderer{}
	r.Add(items...)
	return r
}

func (r *OrRenderer) Len() int {
	return len(r.items)
}

func (r *OrRenderer) At(i int) Code {
	return r.items[i]
}

func (r *OrRenderer) Add(items ...Code) {
	r.items = append(r.items, items...)
	for _, item := range items {
		item.SetContext(r)
	}
}

func (r *OrRenderer) GetContext() Code {
	return r.ctx
}

func (r *OrRenderer) SetContext(ctx Code) {
	r.ctx = ctx
	for _, item := range r.items {
		item.SetContext(r)
	}
}

func (r *OrRenderer) Render(w Writer) {
	for i, item := range r.items {
		if i != 0 {
			w.Write(" || ")
		}
		item.Render(w)
	}
}
