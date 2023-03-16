package gen

var _ Code = (*MapRenderer)(nil)

type MapRenderer struct {
	key   Code
	value Code
	ctx   Code
}

func Map(key Code, value Code) *MapRenderer {
	r := &MapRenderer{}
	r.SetKey(key)
	r.SetValue(value)
	return r
}

func (r *MapRenderer) GetKey() Code {
	return r.key
}

func (r *MapRenderer) SetKey(key Code) {
	r.key = key
	if key != nil {
		key.SetContext(r)
	}
}

func (r *MapRenderer) GetValue() Code {
	return r.value
}

func (r *MapRenderer) SetValue(value Code) {
	r.value = value
	if value != nil {
		value.SetContext(r)
	}
}

func (r *MapRenderer) GetContext() Code {
	return r.ctx
}

func (r *MapRenderer) SetContext(ctx Code) {
	r.ctx = ctx
	if r.key != nil {
		r.key.SetContext(r)
	}
	if r.value != nil {
		r.value.SetContext(r)
	}
}

func (r *MapRenderer) Render(w Writer) {
	w.Write("map[")
	r.key.Render(w)
	w.Write("]")
	r.value.Render(w)
}
