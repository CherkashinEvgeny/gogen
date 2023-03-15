package gen

var _ Sequence = (*SubRenderer)(nil)

type SubRenderer struct {
	items []Code
	ctx   Code
}

func Sub(items ...Code) *SubRenderer {
	r := &SubRenderer{}
	r.Add(items...)
	return r
}

func (r *SubRenderer) Len() int {
	return len(r.items)
}

func (r *SubRenderer) At(i int) Code {
	return r.items[i]
}

func (r *SubRenderer) Add(items ...Code) {
	r.items = append(r.items, items...)
	for _, item := range items {
		item.SetContext(r)
	}
}

func (r *SubRenderer) GetContext() Code {
	return r.ctx
}

func (r *SubRenderer) SetContext(ctx Code) {
	r.ctx = ctx
}

func (r *SubRenderer) Render(w Writer) {
	for i, item := range r.items {
		if i != 0 {
			w.Write(" - ")
		}
		item.Render(w)
	}
}
