package gen

var _ Code = (*AssignRenderer)(nil)

type AssignRenderer struct {
	vars Code
	vals Code
	ctx  Code
}

func Assign(vars Code, vals Code) *AssignRenderer {
	r := &AssignRenderer{}
	r.SetVars(vars)
	r.SetVals(vals)
	return r
}

func (r *AssignRenderer) GetVars() Code {
	return r.vars
}

func (r *AssignRenderer) SetVars(vars Code) {
	r.vars = vars
	if vars != nil {
		vars.SetContext(r)
	}
}

func (r *AssignRenderer) GetVals() Code {
	return r.vals
}

func (r *AssignRenderer) SetVals(vals Code) {
	r.vals = vals
	if vals != nil {
		vals.SetContext(r)
	}
}

func (r *AssignRenderer) GetContext() Code {
	return r.ctx
}

func (r *AssignRenderer) SetContext(ctx Code) {
	r.ctx = ctx
	if r.vars != nil {
		r.vars.SetContext(r)
	}
	if r.vals != nil {
		r.vals.SetContext(r)
	}
}

func (r *AssignRenderer) Render(w Writer) {
	r.vars.Render(w)
	w.Write(" = ")
	r.vals.Render(w)
}
