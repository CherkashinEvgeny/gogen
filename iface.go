package gen

var _ Code = (*IfaceRenderer)(nil)

type IfaceRenderer struct {
	fields Code
	ctx    Code
}

func Iface(methods ...Code) *IfaceRenderer {
	r := &IfaceRenderer{}
	r.SetFields(Methods(methods...))
	return r
}

func (r *IfaceRenderer) GetFields() Code {
	return r.fields
}

func (r *IfaceRenderer) SetFields(fields Code) {
	r.fields = fields
	if fields != nil {
		fields.SetContext(r)
	}
}

func (r *IfaceRenderer) GetContext() Code {
	return r.ctx
}

func (r *IfaceRenderer) SetContext(ctx Code) {
	r.ctx = ctx
}

func (r *IfaceRenderer) Render(w Writer) {
	w.Write("interface {")
	w.Br()
	w.Indent(1)
	r.fields.Render(w)
	w.Indent(-1)
	w.Br()
	w.Write("}")
}

var _ Sequence = (*MethodDeclsRenderer)(nil)

type MethodDeclsRenderer struct {
	items []Code
	ctx   Code
}

func Methods(methods ...Code) *MethodDeclsRenderer {
	i := &MethodDeclsRenderer{}
	i.Add(methods...)
	return i
}

func (l *MethodDeclsRenderer) Len() int {
	return len(l.items)
}

func (l *MethodDeclsRenderer) At(i int) Code {
	return l.items[i]
}

func (l *MethodDeclsRenderer) Add(items ...Code) {
	l.items = append(l.items, items...)
	for _, item := range items {
		item.SetContext(l)
	}
}

func (l *MethodDeclsRenderer) GetContext() Code {
	return l.ctx
}

func (l *MethodDeclsRenderer) SetContext(ctx Code) {
	l.ctx = ctx
}

func (l *MethodDeclsRenderer) Render(w Writer) {
	for i, m := range l.items {
		if i != 0 {
			w.Br()
		}
		m.Render(w)
	}
}

var _ Code = (*MethodDeclRenderer)(nil)

type MethodDeclRenderer struct {
	name      string
	signature Code
	ctx       Code
}

func MethodDecl(name string, signature Code) *MethodDeclRenderer {
	r := &MethodDeclRenderer{}
	r.SetName(name)
	r.SetSignature(signature)
	return r
}

func (r *MethodDeclRenderer) GetName() string {
	return r.name
}

func (r *MethodDeclRenderer) SetName(name string) {
	r.name = name
}

func (r *MethodDeclRenderer) GetSignature() Code {
	return r.signature
}

func (r *MethodDeclRenderer) SetSignature(signature Code) {
	r.signature = signature
	if signature != nil {
		signature.SetContext(r)
	}
}

func (r *MethodDeclRenderer) GetContext() Code {
	return r.ctx
}

func (r *MethodDeclRenderer) SetContext(ctx Code) {
	r.ctx = ctx
}

func (r *MethodDeclRenderer) Render(w Writer) {
	w.Write(r.name)
	r.signature.Render(w)
}
