package gen

var _ Code = (*IfElseRenderer)(nil)

type IfElseRenderer struct {
	cond       Code
	ifBranch   Code
	elseBranch Code
	ctx        Code
}

func If(cond Code, code Code) *IfElseRenderer {
	r := &IfElseRenderer{}
	r.SetCond(cond)
	r.SetIfBranch(code)
	return r
}

func IfElse(cond Code, ifBranch Code, elseBranch Code) *IfElseRenderer {
	r := &IfElseRenderer{}
	r.SetCond(cond)
	r.SetIfBranch(ifBranch)
	r.SetElseBranch(elseBranch)
	return r
}

func (r *IfElseRenderer) GetCond() Code {
	return r.cond
}

func (r *IfElseRenderer) SetCond(cond Code) {
	r.cond = cond
	if cond != nil {
		cond.SetContext(r)
	}
}

func (r *IfElseRenderer) GetIfBranch() Code {
	return r.ifBranch
}

func (r *IfElseRenderer) SetIfBranch(ifBranch Code) {
	r.ifBranch = ifBranch
	if ifBranch != nil {
		ifBranch.SetContext(r)
	}
}

func (r *IfElseRenderer) GetElseBranch() Code {
	return r.elseBranch
}

func (r *IfElseRenderer) SetElseBranch(elseBranch Code) {
	r.elseBranch = elseBranch
	if elseBranch != nil {
		elseBranch.SetContext(r)
	}
}

func (r *IfElseRenderer) GetContext() Code {
	return r.ctx
}

func (r *IfElseRenderer) SetContext(ctx Code) {
	r.ctx = ctx
	if r.cond != nil {
		r.cond.SetContext(r)
	}
	if r.ifBranch != nil {
		r.ifBranch.SetContext(r)
	}
	if r.elseBranch != nil {
		r.elseBranch.SetContext(r)
	}
}

func (r *IfElseRenderer) Render(w Writer) {
	w.Write("if ")
	r.cond.Render(w)
	w.Write(" {")
	if r.ifBranch != nil {
		w.Br()
		w.Indent(1)
		r.ifBranch.Render(w)
		w.Indent(-1)
		w.Br()
	}
	w.Write("}")
	if r.elseBranch != nil {
		w.Write(" else ")
		w.Write("{")
		w.Br()
		w.Indent(1)
		r.ifBranch.Render(w)
		w.Indent(-1)
		w.Br()
		w.Write("}")
	}
}
