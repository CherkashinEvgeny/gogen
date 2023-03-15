package gen

var _ Code = (*AssignAndDeclRenderer)(nil)

type AssignAndDeclRenderer struct {
	vars Code
	vals Code
	ctx  Code
}

func AssignAndDecl(vars Code, vals Code) *AssignAndDeclRenderer {
	r := &AssignAndDeclRenderer{}
	r.SetVars(vars)
	r.SetVals(vals)
	return r
}

func (r *AssignAndDeclRenderer) GetVars() Code {
	return r.vars
}

func (r *AssignAndDeclRenderer) SetVars(vars Code) {
	r.vars = vars
	if vars != nil {
		vars.SetContext(r)
	}
}

func (r *AssignAndDeclRenderer) GetVals() Code {
	return r.vals
}

func (r *AssignAndDeclRenderer) SetVals(vals Code) {
	r.vals = vals
	if vals != nil {
		vals.SetContext(r)
	}
}

func (r *AssignAndDeclRenderer) GetContext() Code {
	return r.ctx
}

func (r *AssignAndDeclRenderer) SetContext(ctx Code) {
	r.ctx = ctx
}

func (r *AssignAndDeclRenderer) Render(w Writer) {
	r.vars.Render(w)
	w.Write(" := ")
	r.vals.Render(w)
}
