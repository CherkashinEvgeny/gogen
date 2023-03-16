package gen

var _ Code = (*PkgRenderer)(nil)

type PkgRenderer struct {
	comment string
	name    string
	imports Code
	code    Code
}

func Pkg(comment string, name string, imports Code, code Code) *PkgRenderer {
	p := &PkgRenderer{}
	p.SetComment(comment)
	p.SetName(name)
	p.SetImports(imports)
	p.SetCode(code)
	return p
}

func (r *PkgRenderer) GetComment() string {
	return r.comment
}

func (r *PkgRenderer) SetComment(comment string) {
	r.comment = comment
}

func (r *PkgRenderer) GetName() string {
	return r.name
}

func (r *PkgRenderer) SetName(name string) {
	r.name = name
}

func (r *PkgRenderer) GetImports() Code {
	return r.imports
}

func (r *PkgRenderer) SetImports(imports Code) {
	r.imports = imports
	if imports != nil {
		imports.SetContext(r)
	}
}

func (r *PkgRenderer) GetCode() Code {
	return r.code
}

func (r *PkgRenderer) SetCode(code Code) {
	r.code = code
	if code != nil {
		code.SetContext(r)
	}
}

func (r *PkgRenderer) GetContext() Code {
	return nil
}

func (r *PkgRenderer) SetContext(_ Code) {
	if r.imports != nil {
		r.imports.SetContext(r)
	}
	if r.code != nil {
		r.code.SetContext(r)
	}
}

func (r *PkgRenderer) Render(w Writer) {
	if r.hasComment() {
		r.renderComment(w)
		w.Br()
	}
	w.Write("package ")
	w.Write(r.name)
	if r.hasImports() {
		w.Br()
		w.Br()
		r.renderImports(w)
	}
	if r.hasCode() {
		w.Br()
		w.Br()
		r.renderCode(w)
	}
}

func (r *PkgRenderer) hasComment() bool {
	return r.comment != ""
}

func (r *PkgRenderer) renderComment(w Writer) {
	CommentRenderer(r.comment).Render(w)
}

func (r *PkgRenderer) hasImports() bool {
	if r.imports == nil {
		return false
	}
	imports, ok := r.imports.(*ImportsRenderer)
	if !ok {
		return true
	}
	return imports.Len() != 0
}

func (r *PkgRenderer) renderImports(w Writer) {
	r.imports.Render(w)
}

func (r *PkgRenderer) hasCode() bool {
	if r.code == nil {
		return false
	}
	seq, ok := r.code.(Sequence)
	if !ok {
		return true
	}
	return seq.Len() != 0
}

func (r *PkgRenderer) renderCode(w Writer) {
	r.code.Render(w)
}
