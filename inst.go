package gen

var _ Code = (*InstRenderer)(nil)

type InstRenderer struct {
	ttype  Code
	values Code
	ctx    Code
}

func Inst(ttype Code, values Code) *InstRenderer {
	r := &InstRenderer{}
	r.SetType(ttype)
	r.SetValues(values)
	return r
}

func (r *InstRenderer) GetType() Code {
	return r.ttype
}

func (r *InstRenderer) SetType(ttype Code) {
	r.ttype = ttype
	if ttype != nil {
		ttype.SetContext(r)
	}
}

func (r *InstRenderer) GetValues() Code {
	return r.values
}

func (r *InstRenderer) SetValues(values Code) {
	r.values = values
	if values != nil {
		values.SetContext(r)
	}
}

func (r *InstRenderer) GetContext() Code {
	return r.ctx
}

func (r *InstRenderer) SetContext(ctx Code) {
	r.ctx = ctx
	if r.ttype != nil {
		r.ttype.SetContext(r)
	}
	if r.values != nil {
		r.values.SetContext(r)
	}
}

func (r *InstRenderer) Render(w Writer) {
	r.ttype.Render(w)
	w.Write("{")
	if !r.isMultiline() {
		r.renderSingleLine(w)
	} else {
		r.renderMultiline(w)
	}
	w.Write("}")
}

func (r *InstRenderer) isMultiline() bool {
	return r.values != nil && IsMultiline(r.values)
}

func (r *InstRenderer) renderSingleLine(w Writer) {
	if r.values == nil || RuneCount(r.values) == 0 {
		return
	}
	w.Write(" ")
	r.values.Render(w)
	w.Write(" ")
}

func (r *InstRenderer) renderMultiline(w Writer) {
	if r.values == nil {
		return
	}
	w.Br()
	w.Indent(1)
	r.values.Render(w)
	w.Indent(-1)
	w.Br()
}

var _ Sequence = (*FieldsRenderer)(nil)

type FieldsRenderer struct {
	items []Code
	ctx   Code
}

func Fields(items ...Code) *FieldsRenderer {
	r := &FieldsRenderer{}
	r.Add(items...)
	return r
}

func (r *FieldsRenderer) Len() int {
	return len(r.items)
}

func (r *FieldsRenderer) At(i int) Code {
	return r.items[i]
}

func (r *FieldsRenderer) Add(items ...Code) {
	r.items = append(r.items, items...)
	for _, item := range items {
		item.SetContext(r)
	}
}

func (r *FieldsRenderer) GetContext() Code {
	return r.ctx
}

func (r *FieldsRenderer) SetContext(ctx Code) {
	r.ctx = ctx
}

func (r *FieldsRenderer) Render(w Writer) {
	if !r.isMultiline() {
		r.renderSingleLine(w)
	} else {
		r.renderMultiline(w)
	}
}

func (r *FieldsRenderer) isMultiline() bool {
	return len(r.items) > 1
}

func (r *FieldsRenderer) renderSingleLine(w Writer) {
	for i, im := range r.items {
		if i != 0 {
			w.Write(", ")
		}
		im.Render(w)
	}
}

func (r *FieldsRenderer) renderMultiline(w Writer) {
	for i, item := range r.items {
		if i != 0 {
			w.Br()
		}
		item.Render(w)
		w.Write(",")
	}
}

type FieldRenderer struct {
	name  string
	ftype Code
	ctx   Code
}

func Field(name string, ftype Code) *FieldRenderer {
	r := &FieldRenderer{}
	r.SetName(name)
	r.SetType(ftype)
	return r
}

func (r *FieldRenderer) GetName() string {
	return r.name
}

func (r *FieldRenderer) SetName(name string) {
	r.name = name
}

func (r *FieldRenderer) GetType() Code {
	return r.ftype
}

func (r *FieldRenderer) SetType(ftype Code) {
	r.ftype = ftype
	if ftype != nil {
		ftype.SetContext(r)
	}
}

func (r *FieldRenderer) GetContext() Code {
	return r.ctx
}

func (r *FieldRenderer) SetContext(ctx Code) {
	r.ctx = ctx
}

func (r *FieldRenderer) Render(w Writer) {
	w.Write(r.name)
	w.Write(": ")
	indent := r.indent()
	for i := 0; i < indent; i++ {
		w.Write(" ")
	}
	r.ftype.Render(w)
}

func (r *FieldRenderer) indent() int {
	fields, ok := r.ctx.(*FieldsRenderer)
	if !ok {
		return 0
	}
	if !fields.isMultiline() {
		return 0
	}
	selfIndex := 0
	for selfIndex < fields.Len() {
		field := fields.At(selfIndex)
		self, ok := field.(*FieldRenderer)
		if ok && r == self {
			break
		}
		selfIndex++
	}
	maxGap := 0
	for i := selfIndex - 1; i >= 0; i-- {
		field := fields.At(i)
		previousField, ok := field.(*FieldRenderer)
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
		nextField, ok := field.(*FieldRenderer)
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
