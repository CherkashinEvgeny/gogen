package gen

var _ Code = (*StructRenderer)(nil)

type StructRenderer struct {
	fields Code
	ctx    Code
}

func Struct(fields Code) *StructRenderer {
	r := &StructRenderer{}
	r.SetFields(fields)
	return r
}

func (r *StructRenderer) GetFields() Code {
	return r.fields
}

func (r *StructRenderer) SetFields(fields Code) {
	r.fields = fields
	if fields != nil {
		fields.SetContext(r)
	}
}

func (r *StructRenderer) GetContext() Code {
	return r.ctx
}

func (r *StructRenderer) SetContext(ctx Code) {
	r.ctx = ctx
}

func (r *StructRenderer) Render(w Writer) {
	multiline := r.isMultiline()
	w.Write("struct")
	if multiline {
		w.Write(" ")
	}
	w.Write("{")
	if !multiline {
		r.renderSingleLine(w)
	} else {
		r.renderMultiline(w)
	}
	w.Write("}")
}

func (r *StructRenderer) isMultiline() bool {
	return r.fields != nil && IsMultiline(r.fields)
}

func (r *StructRenderer) renderSingleLine(w Writer) {
	if r.fields == nil || RuneCount(r.fields) == 0 {
		return
	}
	w.Write(" ")
	r.fields.Render(w)
	w.Write(" ")
}

func (r *StructRenderer) renderMultiline(w Writer) {
	if r.fields == nil {
		return
	}
	w.Br()
	w.Indent(1)
	r.fields.Render(w)
	w.Indent(-1)
	w.Br()
}

var _ Sequence = (*FieldDeclsRenderer)(nil)

type FieldDeclsRenderer struct {
	items []Code
	ctx   Code
}

func FieldDecls(items ...Code) *FieldDeclsRenderer {
	i := &FieldDeclsRenderer{}
	i.Add(items...)
	return i
}

func (r *FieldDeclsRenderer) Len() int {
	return len(r.items)
}

func (r *FieldDeclsRenderer) At(i int) Code {
	return r.items[i]
}

func (r *FieldDeclsRenderer) Add(items ...Code) {
	r.items = append(r.items, items...)
	for _, item := range items {
		item.SetContext(r)
	}
}

func (r *FieldDeclsRenderer) GetContext() Code {
	return r.ctx
}

func (r *FieldDeclsRenderer) SetContext(ctx Code) {
	r.ctx = ctx
}

func (r *FieldDeclsRenderer) Render(w Writer) {
	if !r.isMultiline() {
		r.renderSingleLine(w)
	} else {
		r.renderMultiline(w)
	}
}

func (r *FieldDeclsRenderer) isMultiline() bool {
	return len(r.items) > 1
}

func (r *FieldDeclsRenderer) renderSingleLine(w Writer) {
	for i, im := range r.items {
		if i != 0 {
			w.Write("; ")
		}
		im.Render(w)
	}
}

func (r *FieldDeclsRenderer) renderMultiline(w Writer) {
	for i, im := range r.items {
		if i != 0 {
			w.Br()
		}
		im.Render(w)
	}
}

var _ Code = (*FieldDeclRenderer)(nil)

type FieldDeclRenderer struct {
	name  string
	ttype Code
	ctx   Code
}

func FieldDecl(name string, ftype Code) *FieldDeclRenderer {
	r := &FieldDeclRenderer{}
	r.SetName(name)
	r.SetType(ftype)
	return r
}

func (r *FieldDeclRenderer) GetName() string {
	return r.name
}

func (r *FieldDeclRenderer) SetName(name string) {
	r.name = name
}

func (r *FieldDeclRenderer) GetType() Code {
	return r.ttype
}

func (r *FieldDeclRenderer) SetType(ftype Code) {
	r.ttype = ftype
	if ftype != nil {
		ftype.SetContext(r)
	}
}

func (r *FieldDeclRenderer) GetContext() Code {
	return r.ctx
}

func (r *FieldDeclRenderer) SetContext(ctx Code) {
	r.ctx = ctx
}

func (r *FieldDeclRenderer) Render(w Writer) {
	w.Write(r.name)
	w.Write(" ")
	indent := r.indent()
	for i := 0; i < indent; i++ {
		w.Write(" ")
	}
	r.ttype.Render(w)
}

func (r *FieldDeclRenderer) indent() int {
	fields, ok := r.ctx.(*FieldDeclsRenderer)
	if !ok {
		return 0
	}
	if !fields.isMultiline() {
		return 0
	}
	selfIndex := 0
	for selfIndex < fields.Len() {
		field := fields.At(selfIndex)
		self, ok := field.(*FieldDeclRenderer)
		if ok && r == self {
			break
		}
		selfIndex++
	}
	maxGap := 0
	for i := selfIndex - 1; i >= 0; i-- {
		field := fields.At(i)
		previousField, ok := field.(*FieldDeclRenderer)
		if !ok {
			break
		}
		gap := len(previousField.GetName()) - len(r.GetName())
		if gap > maxGap {
			maxGap = gap
		}
	}
	for i := selfIndex + 1; i < fields.Len(); i++ {
		field := fields.At(i)
		nextField, ok := field.(*FieldDeclRenderer)
		if !ok {
			break
		}
		gap := len(nextField.GetName()) - len(r.GetName())
		if gap > maxGap {
			maxGap = gap
		}
	}
	return maxGap
}
