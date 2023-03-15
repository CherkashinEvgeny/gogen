package gen

var _ Code = (*SignRenderer)(nil)

type SignRenderer struct {
	in  Code
	out Code
	ctx Code
}

func Sign(in Code, out Code) *SignRenderer {
	r := &SignRenderer{}
	r.SetIn(in)
	r.SetOut(out)
	return r
}

func (r *SignRenderer) GetIn() Code {
	return r.in
}

func (r *SignRenderer) SetIn(in Code) {
	r.in = in
	if in != nil {
		in.SetContext(r)
	}
}

func (r *SignRenderer) GetOut() Code {
	return r.out
}

func (r *SignRenderer) SetOut(out Code) {
	r.out = out
	if out != nil {
		out.SetContext(r)
	}
}

func (r *SignRenderer) GetContext() Code {
	return r.ctx
}

func (r *SignRenderer) SetContext(ctx Code) {
	r.ctx = ctx
}

func (r *SignRenderer) Render(w Writer) {
	r.renderIn(w)
	if r.hasOut() {
		w.Write(" ")
		r.renderOut(w)
	}
}

func (r *SignRenderer) renderIn(w Writer) {
	w.Write("(")
	if r.in != nil {
		multiline := IsMultiline(r.in)
		if multiline {
			w.Br()
		}
		r.in.Render(w)
		if multiline {
			w.Br()
		}
	}
	w.Write(")")
}

func (r *SignRenderer) hasOut() bool {
	return RuneCount(r.out) > 0
}

func (r *SignRenderer) renderOut(w Writer) {
	bracket := r.needOutBrackets()
	if bracket {
		w.Write("(")
	}
	if r.out != nil {
		multiline := IsMultiline(r.out)
		if multiline {
			w.Br()
		}
		r.out.Render(w)
		if multiline {
			w.Br()
		}
	}
	if bracket {
		w.Write(")")
	}
}

func (r *SignRenderer) needOutBrackets() bool {
	if r.out == nil {
		return false
	}
	params, ok := r.out.(*ParamsRenderer)
	if !ok {
		return true
	}
	if params.Len() == 0 {
		return false
	}
	if params.Len() > 1 {
		return true
	}
	param, ok := params.At(0).(*ParamRenderer)
	if !ok {
		return true
	}
	return param.GetName() != ""
}

var _ Sequence = (*ParamsRenderer)(nil)

type ParamsRenderer struct {
	items []Code
	ctx   Code
}

func In(items ...Code) *ParamsRenderer {
	return Params(items...)
}

func Out(items ...Code) *ParamsRenderer {
	return Params(items...)
}

func Params(items ...Code) *ParamsRenderer {
	r := &ParamsRenderer{}
	r.Add(items...)
	return r
}

func (r *ParamsRenderer) Len() int {
	return len(r.items)
}

func (r *ParamsRenderer) At(i int) Code {
	return r.items[i]
}

func (r *ParamsRenderer) Add(items ...Code) {
	r.items = append(r.items, items...)
	for _, item := range items {
		item.SetContext(r)
	}
}

func (r *ParamsRenderer) GetContext() Code {
	return r.ctx
}

func (r *ParamsRenderer) SetContext(ctx Code) {
	r.ctx = ctx
}

func (r *ParamsRenderer) Render(w Writer) {
	if !r.isMultiline() {
		r.renderSingleLine(w)
	} else {
		r.renderMultiline(w)
	}
}

func (r *ParamsRenderer) isMultiline() bool {
	return len(r.items) > 3
}

func (r *ParamsRenderer) renderSingleLine(w Writer) {
	for i, item := range r.items {
		if i != 0 {
			w.Write(", ")
		}
		item.Render(w)
	}
}

func (r *ParamsRenderer) renderMultiline(w Writer) {
	for _, item := range r.items {
		item.Render(w)
		w.Write(",")
		w.Br()
	}
}

var _ Code = (*ParamRenderer)(nil)

type ParamRenderer struct {
	name     string
	ttype    Code
	variadic bool
	ctx      Code
}

func Param(name string, ttype Code, variadic bool) *ParamRenderer {
	r := &ParamRenderer{}
	r.SetName(name)
	r.SetType(ttype)
	r.SetVariadic(variadic)
	return r
}

func (r *ParamRenderer) GetName() string {
	return r.name
}

func (r *ParamRenderer) SetName(name string) {
	r.name = name
}

func (r *ParamRenderer) GetType() Code {
	return r.ttype
}

func (r *ParamRenderer) SetType(ptype Code) {
	r.ttype = ptype
	if ptype != nil {
		ptype.SetContext(r)
	}
}

func (r *ParamRenderer) GetVariadic() bool {
	return r.variadic
}

func (r *ParamRenderer) SetVariadic(variadic bool) {
	r.variadic = variadic
}

func (r *ParamRenderer) GetContext() Code {
	return r.ctx
}

func (r *ParamRenderer) SetContext(ctx Code) {
	r.ctx = ctx
}

func (r *ParamRenderer) Render(w Writer) {
	if r.name != "" {
		w.Write(r.name)
		w.Write(" ")
	}
	if r.variadic {
		w.Write("...")
	}
	r.ttype.Render(w)
}
