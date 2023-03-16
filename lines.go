package gen

var _ Sequence = (*LinesRenderer)(nil)

type LinesRenderer struct {
	items []Code
	ctx   Code
}

func Lines(codes ...Code) *LinesRenderer {
	r := &LinesRenderer{}
	r.Add(codes...)
	return r
}

func (r *LinesRenderer) Len() int {
	return len(r.items)
}

func (r *LinesRenderer) At(i int) Code {
	return r.items[i]
}

func (r *LinesRenderer) Add(items ...Code) {
	r.items = append(r.items, items...)
	for _, item := range items {
		item.SetContext(r)
	}
}

func (r *LinesRenderer) GetContext() Code {
	return r.ctx
}

func (r *LinesRenderer) SetContext(ctx Code) {
	r.ctx = ctx
	for _, item := range r.items {
		item.SetContext(r)
	}
}

func (r *LinesRenderer) Render(w Writer) {
	for i, item := range r.items {
		if i != 0 {
			w.Br()
		}
		item.Render(w)
	}
}
