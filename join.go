package gen

var _ Sequence = (*JoinRenderer)(nil)

type JoinRenderer struct {
	items []Code
	ctx   Code
}

func Concat(strs ...string) *JoinRenderer {
	codes := make([]Code, 0, len(strs))
	for _, str := range strs {
		codes = append(codes, Raw(str))
	}
	return Join(codes...)
}

func Join(codes ...Code) *JoinRenderer {
	r := &JoinRenderer{}
	r.Add(codes...)
	return r
}

func (r *JoinRenderer) Len() int {
	return len(r.items)
}

func (r *JoinRenderer) At(i int) Code {
	return r.items[i]
}

func (r *JoinRenderer) Add(items ...Code) {
	r.items = append(r.items, items...)
	for _, item := range items {
		item.SetContext(r)
	}
}

func (r *JoinRenderer) GetContext() Code {
	return r.ctx
}

func (r *JoinRenderer) SetContext(ctx Code) {
	r.ctx = ctx
}

func (r *JoinRenderer) Render(w Writer) {
	for _, item := range r.items {
		item.Render(w)
	}
}
