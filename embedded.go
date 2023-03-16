package gen

var _ Code = (*EmbeddedRenderer)(nil)

type EmbeddedRenderer struct {
	ftype Code
	ctx   Code
}

func Embedded(ftype Code) *EmbeddedRenderer {
	r := &EmbeddedRenderer{}
	r.SetType(ftype)
	return r
}

func (r *EmbeddedRenderer) GetType() Code {
	return r.ftype
}

func (r *EmbeddedRenderer) SetType(ftype Code) {
	r.ftype = ftype
	if ftype != nil {
		ftype.SetContext(r)
	}
}

func (r *EmbeddedRenderer) GetContext() Code {
	return r.ctx
}

func (r *EmbeddedRenderer) SetContext(ctx Code) {
	r.ctx = ctx
	if r.ftype != nil {
		r.ftype.SetContext(r)
	}
}

func (r *EmbeddedRenderer) Render(w Writer) {
	r.ftype.Render(w)
}
