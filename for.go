package gen

var _ Code = (*ForRenderer)(nil)

type ForRenderer struct {
	cond Code
	body Code
	ctx  Code
}

func For(cond Code, body Code) *ForRenderer {
	r := &ForRenderer{}
	r.SetCond(cond)
	r.SetBody(body)
	return r
}

func (r *ForRenderer) GetCond() Code {
	return r.cond
}

func (r *ForRenderer) SetCond(cond Code) {
	r.cond = cond
	if cond != nil {
		cond.SetContext(r)
	}
}

func (r *ForRenderer) GetBody() Code {
	return r.body
}

func (r *ForRenderer) SetBody(body Code) {
	r.body = body
	if body != nil {
		body.SetContext(r)
	}
}

func (r *ForRenderer) GetContext() Code {
	return r.ctx
}

func (r *ForRenderer) SetContext(ctx Code) {
	r.ctx = ctx
	if r.cond != nil {
		r.cond.SetContext(r)
	}
	if r.body != nil {
		r.body.SetContext(r)
	}
}

func (r *ForRenderer) Render(w Writer) {
	w.Write("for ")
	if r.cond != nil {
		r.cond.Render(w)
		w.Write(" ")
	}
	w.Write("{")
	if r.body != nil {
		w.Br()
		w.Indent(1)
		r.body.Render(w)
		w.Indent(-1)
		w.Br()
	}
	w.Write("}")
}
