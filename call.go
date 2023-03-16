package gen

var _ Code = (*CallRenderer)(nil)

type CallRenderer struct {
	ffunc Code
	args  Code
	ctx   Code
}

func Call(ffunc Code, args Code) *CallRenderer {
	r := &CallRenderer{}
	r.SetFunc(ffunc)
	r.SetArgs(args)
	return r
}

func (r *CallRenderer) GetFunc() Code {
	return r.ffunc
}

func (r *CallRenderer) SetFunc(ffunc Code) {
	r.ffunc = ffunc
	if ffunc != nil {
		ffunc.SetContext(r)
	}
}

func (r *CallRenderer) GetArgs() Code {
	return r.args
}

func (r *CallRenderer) SetArgs(args Code) {
	r.args = args
	if args != nil {
		args.SetContext(r)
	}
}

func (r *CallRenderer) GetContext() Code {
	return r.ctx
}

func (r *CallRenderer) SetContext(ctx Code) {
	r.ctx = ctx
}

func (r *CallRenderer) Render(w Writer) {
	r.ffunc.Render(w)
	w.Write("(")
	r.args.Render(w)
	w.Write(")")
}
