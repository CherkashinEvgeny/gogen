package gen

var _ Code = (*FuncRenderer)(nil)

type FuncRenderer struct {
	name      string
	signature Code
	body      Code
	ctx       Code
}

func Func(name string, signature Code, body Code) *FuncRenderer {
	r := &FuncRenderer{}
	r.SetName(name)
	r.SetSignature(signature)
	r.SetBody(body)
	return r
}

func (r *FuncRenderer) GetName() string {
	return r.name
}

func (r *FuncRenderer) SetName(name string) {
	r.name = name
}

func (r *FuncRenderer) GetSignature() Code {
	return r.signature
}

func (r *FuncRenderer) SetSignature(signature Code) {
	r.signature = signature
	if signature != nil {
		signature.SetContext(r)
	}
}

func (r *FuncRenderer) GetBody() Code {
	return r.body
}

func (r *FuncRenderer) SetBody(body Code) {
	r.body = body
	if body != nil {
		body.SetContext(r)
	}
}

func (r *FuncRenderer) GetContext() Code {
	return r.ctx
}

func (r *FuncRenderer) SetContext(ctx Code) {
	r.ctx = ctx
	if r.signature != nil {
		r.signature.SetContext(r)
	}
	if r.body != nil {
		r.body.SetContext(r)
	}
}

func (r *FuncRenderer) Render(w Writer) {
	w.Write("func ")
	w.Write(r.name)
	r.signature.Render(w)
	w.Write(" {")
	if r.body != nil {
		w.Br()
		w.Indent(1)
		r.body.Render(w)
		w.Indent(-1)
		w.Br()
	}
	w.Write("}")
}
