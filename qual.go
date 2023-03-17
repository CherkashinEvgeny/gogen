package gen

var _ Code = (*SmartQualRenderer)(nil)

type SmartQualRenderer struct {
	pkgName string
	pkgPath string
	name    string
	ctx     Code
}

func SmartQual(pkgName string, pkgPath string, name string) *SmartQualRenderer {
	r := &SmartQualRenderer{}
	r.SetPkgName(pkgName)
	r.SetPkgPath(pkgPath)
	r.SetName(name)
	return r
}

func (r *SmartQualRenderer) GetPkgName() string {
	return r.pkgName
}

func (r *SmartQualRenderer) SetPkgName(name string) {
	r.pkgName = name
}

func (r *SmartQualRenderer) GetPkgPath() string {
	return r.pkgPath
}

func (r *SmartQualRenderer) SetPkgPath(path string) {
	r.pkgPath = path
}

func (r *SmartQualRenderer) GetName() string {
	return r.name
}

func (r *SmartQualRenderer) SetName(name string) {
	r.name = name
}

func (r *SmartQualRenderer) GetContext() Code {
	return r.ctx
}

func (r *SmartQualRenderer) SetContext(ctx Code) {
	r.ctx = ctx
	r.addImport()
}

func (r *SmartQualRenderer) addImport() {
	if r.pkgPath == "" {
		return
	}
	imports := r.findImports()
	if imports == nil {
		return
	}
	imports.Add(SmartImport(r.pkgName, "", r.pkgPath))
}

func (r *SmartQualRenderer) Render(w Writer) {
	prefix := r.resolvePkgAlias()
	if prefix != "" && prefix != "." {
		w.Write(prefix)
		w.Write(".")
	}
	w.Write(r.name)
}

func (r *SmartQualRenderer) resolvePkgAlias() string {
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

func (r *SmartQualRenderer) findImports() *ImportsRenderer {
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

func (r *SmartQualRenderer) findPackage() *PkgRenderer {
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

var _ Code = (*QualRenderer)(nil)

type QualRenderer struct {
	pkg  string
	name string
	ctx  Code
}

func Qual(pkg string, name string) *QualRenderer {
	r := &QualRenderer{}
	r.SetPkg(pkg)
	r.SetName(name)
	return r
}

func (r *QualRenderer) GetPkg() string {
	return r.pkg
}

func (r *QualRenderer) SetPkg(name string) {
	r.pkg = name
}

func (r *QualRenderer) GetName() string {
	return r.name
}

func (r *QualRenderer) SetName(name string) {
	r.name = name
}

func (r *QualRenderer) GetContext() Code {
	return r.ctx
}

func (r *QualRenderer) SetContext(ctx Code) {
	r.ctx = ctx
}

func (r *QualRenderer) Render(w Writer) {
	if r.pkg != "" && r.pkg != "." {
		w.Write(r.pkg)
		w.Write(".")
	}
	w.Write(r.name)
}
