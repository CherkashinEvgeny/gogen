package gen

var _ Sequence = (*BlocksRenderer)(nil)

type BlocksRenderer struct {
	items []Code
	ctx   Code
}

func Blocks(items ...Code) *BlocksRenderer {
	r := &BlocksRenderer{}
	r.Add(items...)
	return r
}

func (r *BlocksRenderer) Len() int {
	return len(r.items)
}

func (r *BlocksRenderer) At(i int) Code {
	return r.items[i]
}

func (r *BlocksRenderer) Add(items ...Code) {
	r.items = append(r.items, items...)
	for _, item := range items {
		item.SetContext(r)
	}
}

func (r *BlocksRenderer) GetContext() Code {
	return r.ctx
}

func (r *BlocksRenderer) SetContext(ctx Code) {
	r.ctx = ctx
	for _, item := range r.items {
		item.SetContext(r)
	}
}

func (r *BlocksRenderer) Render(w Writer) {
	for i, item := range r.items {
		if i != 0 {
			w.Br()
			w.Br()
		}
		item.Render(w)
	}
}
