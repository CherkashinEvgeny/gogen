package gen

var _ Sequence = (*IdsRenderer)(nil)

type IdsRenderer struct {
	items []Code
	ctx   Code
}

func Ids(vars ...string) *IdsRenderer {
	r := &IdsRenderer{}
	varsCode := make([]Code, 0, len(vars))
	for _, v := range vars {
		varsCode = append(varsCode, Id(v))
	}
	r.Add(varsCode...)
	return r
}

func (r *IdsRenderer) Len() int {
	return len(r.items)
}

func (r *IdsRenderer) At(i int) Code {
	return r.items[i]
}

func (r *IdsRenderer) Add(items ...Code) {
	r.items = append(r.items, items...)
	for _, item := range items {
		item.SetContext(r)
	}
}

func (r *IdsRenderer) GetContext() Code {
	return r.ctx
}

func (r *IdsRenderer) SetContext(ctx Code) {
	r.ctx = ctx
	for _, item := range r.items {
		item.SetContext(r)
	}
}

func (r *IdsRenderer) Render(w Writer) {
	for i, item := range r.items {
		if i != 0 {
			w.Write(", ")
		}
		item.Render(w)
	}
}

var _ Code = IdRenderer("")

type IdRenderer string

func Id(v string) IdRenderer {
	return IdRenderer(v)
}

func (r IdRenderer) GetContext() Code {
	return nil
}

func (r IdRenderer) SetContext(ctx Code) {
}

func (r IdRenderer) Render(w Writer) {
	w.Write(string(r))
}
