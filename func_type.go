package gen

var _ Code = (*FuncTypeRenderer)(nil)

type FuncTypeRenderer struct {
	signature Code
	ctx       Code
}

func FuncType(sign Code) *FuncTypeRenderer {
	r := &FuncTypeRenderer{}
	r.SetSignature(sign)
	return r
}

func (r *FuncTypeRenderer) GetSignature() Code {
	return r.signature
}

func (r *FuncTypeRenderer) SetSignature(signature Code) {
	r.signature = signature
	if signature != nil {
		signature.SetContext(r)
	}
}

func (r *FuncTypeRenderer) GetContext() Code {
	return r.ctx
}

func (r *FuncTypeRenderer) SetContext(ctx Code) {
	r.ctx = ctx
	if r.signature != nil {
		r.signature.SetContext(ctx)
	}
}

func (r *FuncTypeRenderer) Render(w Writer) {
	w.Write("func")
	r.signature.Render(w)
}
