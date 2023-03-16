package gen

var _ Code = (*SmartNamedRenderer)(nil)

type SmartNamedRenderer struct {
	pkgName string
	pkgPath string
	name    string
	ctx     Code
}

func SmartNamed(pkgName string, pkgPath string, name string) *SmartNamedRenderer {
	r := &SmartNamedRenderer{}
	r.SetPkgName(pkgName)
	r.SetPkgPath(pkgPath)
	r.SetName(name)
	return r
}

func (r *SmartNamedRenderer) GetPkgName() string {
	return r.pkgName
}

func (r *SmartNamedRenderer) SetPkgName(name string) {
	r.pkgName = name
}

func (r *SmartNamedRenderer) GetPkgPath() string {
	return r.pkgPath
}

func (r *SmartNamedRenderer) SetPkgPath(path string) {
	r.pkgPath = path
}

func (r *SmartNamedRenderer) GetName() string {
	return r.name
}

func (r *SmartNamedRenderer) SetName(name string) {
	r.name = name
}

func (r *SmartNamedRenderer) GetContext() Code {
	return r.ctx
}

func (r *SmartNamedRenderer) SetContext(ctx Code) {
	r.ctx = ctx
	r.addImport()
}

func (r *SmartNamedRenderer) addImport() {
	if r.pkgPath == "" {
		return
	}
	imports := r.findImports()
	if imports == nil {
		return
	}
	imports.Add(SmartImport(r.pkgName, "", r.pkgPath))
}

func (r *SmartNamedRenderer) Render(w Writer) {
	prefix := r.resolvePkgAlias()
	if prefix != "" && prefix != "." {
		w.Write(prefix)
		w.Write(".")
	}
	w.Write(r.name)
}

func (r *SmartNamedRenderer) resolvePkgAlias() string {
	if r.pkgPath == "" {
		return r.pkgName
	}
	imports := r.findImports()
	if imports == nil {
		return r.pkgName
	}
	alias, found := imports.Resolve(r.pkgPath)
	if !found {
		return r.pkgName
	}
	return alias
}

func (r *SmartNamedRenderer) findImports() *ImportsRenderer {
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

func (r *SmartNamedRenderer) findPackage() *PkgRenderer {
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

var _ Code = (*NamedRenderer)(nil)

type NamedRenderer struct {
	pkg  string
	name string
	ctx  Code
}

func Named(pkg string, name string) *NamedRenderer {
	r := &NamedRenderer{}
	r.SetPkg(pkg)
	r.SetName(name)
	return r
}

func (r *NamedRenderer) GetPkg() string {
	return r.pkg
}

func (r *NamedRenderer) SetPkg(name string) {
	r.pkg = name
}

func (r *NamedRenderer) GetName() string {
	return r.name
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
	if r.pkg != "" && r.pkg != "." {
		w.Write(r.pkg)
		w.Write(".")
	}
	w.Write(r.name)
}
