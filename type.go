package gen

var _ Code = (*TypeRenderer)(nil)

type TypeRenderer struct {
	name    string
	ttype   Code
	methods Code
	ctx     Code
}

func Type(name string, ttype Code) *TypeRenderer {
	r := &TypeRenderer{}
	r.SetName(name)
	r.SetType(ttype)
	return r
}

func (r *TypeRenderer) GetName() string {
	return r.name
}

func (r *TypeRenderer) SetName(name string) {
	r.name = name
}

func (r *TypeRenderer) GetType() Code {
	return r.ttype
}

func (r *TypeRenderer) SetType(ttype Code) {
	r.ttype = ttype
	if ttype != nil {
		ttype.SetContext(r)
	}
}

func (r *TypeRenderer) GetContext() Code {
	return r.ctx
}

func (r *TypeRenderer) SetContext(ctx Code) {
	r.ctx = ctx
}

func (r *TypeRenderer) Render(w Writer) {
	w.Write("type")
	w.Write(" ")
	w.Write(r.name)
	w.Write(" ")
	r.ttype.Render(w)
	if r.methods != nil {
		w.Br()
		w.Br()
		r.methods.Render(w)
	}
}
