package gen

var _ Sequence = (*MultRenderer)(nil)

type MultRenderer struct {
	items []Code
	ctx   Code
}

func Mult(items ...Code) *MultRenderer {
	r := &MultRenderer{}
	r.Add(items...)
	return r
}

func (r *MultRenderer) Len() int {
	return len(r.items)
}

func (r *MultRenderer) At(i int) Code {
	return r.items[i]
}

func (r *MultRenderer) Add(items ...Code) {
	r.items = append(r.items, items...)
	for _, item := range items {
		item.SetContext(r)
	}
}

func (r *MultRenderer) GetContext() Code {
	return r.ctx
}

func (r *MultRenderer) SetContext(ctx Code) {
	r.ctx = ctx
	for _, item := range r.items {
		item.SetContext(r)
	}
}

func (r *MultRenderer) Render(w Writer) {
	for i, item := range r.items {
		if i != 0 {
			w.Write(" * ")
		}
		item.Render(w)
	}
}
