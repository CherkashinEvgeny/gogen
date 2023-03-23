package gen

var _ Sequence = (*DotRenderer)(nil)

type DotRenderer struct {
	items []Code
	ctx   Code
}

func Dot(items ...Code) *DotRenderer {
	r := &DotRenderer{}
	r.Add(items...)
	return r
}

func (r *DotRenderer) Len() int {
	return len(r.items)
}

func (r *DotRenderer) At(i int) Code {
	return r.items[i]
}

func (r *DotRenderer) Add(items ...Code) {
	r.items = append(r.items, items...)
	for _, item := range items {
		item.SetContext(r)
	}
}

func (r *DotRenderer) GetContext() Code {
	return r.ctx
}

func (r *DotRenderer) SetContext(ctx Code) {
	r.ctx = ctx
	for _, item := range r.items {
		item.SetContext(r)
	}
}

func (r *DotRenderer) Render(w Writer) {
	for i, item := range r.items {
		if i != 0 {
			w.Write(".")
		}
		item.Render(w)
	}
}
