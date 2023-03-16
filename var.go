package gen

var _ Code = (*VarRenderer)(nil)

type VarRenderer struct {
	ids   Code
	ttype Code
	ctx   Code
}

func Var(ids Code, ttype Code) *VarRenderer {
	r := &VarRenderer{}
	r.SetIds(ids)
	r.SetType(ttype)
	return r
}

func (r *VarRenderer) GetIds() Code {
	return r.ids
}

func (r *VarRenderer) SetIds(ids Code) {
	r.ids = ids
	if ids != nil {
		ids.SetContext(r)
	}
}

func (r *VarRenderer) GetType() Code {
	return r.ids
}

func (r *VarRenderer) SetType(ttype Code) {
	r.ttype = ttype
	if ttype != nil {
		ttype.SetContext(r)
	}
}

func (r *VarRenderer) GetContext() Code {
	return r.ctx
}

func (r *VarRenderer) SetContext(ctx Code) {
	r.ctx = ctx
	if r.ids != nil {
		r.ids.SetContext(r)
	}
	if r.ttype != nil {
		r.ttype.SetContext(r)
	}
}

func (r *VarRenderer) Render(w Writer) {
	w.Write("var ")
	r.ids.Render(w)
	if r.ttype != nil {
		r.ttype.Render(w)
	}
}
