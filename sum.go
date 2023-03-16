package gen

var _ Sequence = (*SumRenderer)(nil)

type SumRenderer struct {
	items []Code
	ctx   Code
}

func Sum(items ...Code) *SumRenderer {
	r := &SumRenderer{}
	r.Add(items...)
	return r
}

func (r *SumRenderer) Len() int {
	return len(r.items)
}

func (r *SumRenderer) At(i int) Code {
	return r.items[i]
}

func (r *SumRenderer) Add(items ...Code) {
	r.items = append(r.items, items...)
	for _, item := range items {
		item.SetContext(r)
	}
}

func (r *SumRenderer) GetContext() Code {
	return r.ctx
}

func (r *SumRenderer) SetContext(ctx Code) {
	r.ctx = ctx
	for _, item := range r.items {
		item.SetContext(r)
	}
}

func (r *SumRenderer) Render(w Writer) {
	for i, item := range r.items {
		if i != 0 {
			w.Write(" + ")
		}
		item.Render(w)
	}
}
