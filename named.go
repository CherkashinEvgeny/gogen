package gen

var _ Code = (*NamedRenderer)(nil)

type NamedRenderer struct {
	path string
	name string
	ctx  Code
}

func Named(path string, name string) *NamedRenderer {
	r := &NamedRenderer{}
	r.SetPath(path)
	r.SetName(name)
	return r
}

func (r *NamedRenderer) SetPath(path string) {
	r.path = path
}

func (r *NamedRenderer) SetName(name string) {
	r.name = name
}

func (r *NamedRenderer) GetContext() Code {
	return r.ctx
}

func (r *NamedRenderer) SetContext(ctx Code) {
	r.ctx = ctx
}

func (r *NamedRenderer) Render(w Writer) {
	prefix := r.resolvePrefix()
	if prefix != "" {
		w.Write(prefix)
		w.Write(".")
	}
	w.Write(r.name)
}

func (r *NamedRenderer) resolvePrefix() string {
	im := r.findImport()
	if im == nil {
		return ""
	}
	alias := im.GetAlias()
	if alias == "." || alias == "_" {
		return ""
	}
	if alias != "" {
		return alias
	}
	return im.GetName()
}

func (r *NamedRenderer) findImport() *ImportRenderer {
	imports := r.findImports()
	if imports == nil {
		return nil
	}
	c := imports.Len()
	for i := 0; i < c; i++ {
		importCandidate := imports.At(i)
		im, ok := importCandidate.(*ImportRenderer)
		if ok && r.path == im.GetPath() {
			return im
		}
	}
	return nil
}

func (r *NamedRenderer) findImports() *ImportsRenderer {
	pkg := r.findPackage()
	if pkg == nil {
		return nil
	}
	importsCandidate := pkg.GetImports()
	imports, ok := importsCandidate.(*ImportsRenderer)
	if !ok {
		return nil
	}
	return imports
}

func (r *NamedRenderer) findPackage() *PkgRenderer {
	ctx := r.GetContext()
	for ctx != nil {
		pkg, ok := ctx.(*PkgRenderer)
		if ok {
			return pkg
		}
		ctx = ctx.GetContext()
	}
	return nil
}
