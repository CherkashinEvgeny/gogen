package gen

import (
	"strconv"
	"strings"
)

var _ Sequence = (*ImportsRenderer)(nil)

type ImportsRenderer struct {
	items    []Code
	pathMap  map[string]string
	aliasMap map[string]struct{}
	ctx      Code
}

func Imports(items ...Code) *ImportsRenderer {
	r := &ImportsRenderer{}
	r.init()
	r.Add(items...)
	return r
}

func (r *ImportsRenderer) Resolve(path string) (string, bool) {
	r.init()
	alias, found := r.pathMap[path]
	return alias, found
}

func (r *ImportsRenderer) Len() int {
	return len(r.items)
}

func (r *ImportsRenderer) At(i int) Code {
	return r.items[i]
}

func (r *ImportsRenderer) Add(items ...Code) {
	r.init()
	for _, item := range items {
		r.add(item)
	}
}

func (r *ImportsRenderer) add(item Code) {
	im, ok := item.(*SmartImportRenderer)
	if !ok {
		r.items = append(r.items, item)
		item.SetContext(r)
		return
	}
	alias, found := r.pathMap[im.GetPath()]
	if found && alias != "_" {
		return
	}
	preferredAlias := im.GetName()
	if im.GetAlias() != "" {
		preferredAlias = im.GetAlias()
	}
	alias = r.generatePkgAlias(preferredAlias)
	if preferredAlias != alias {
		im.SetAlias(alias)
	}
	if !found {
		r.items = append(r.items, item)
		item.SetContext(r)
	}
}

func (r *ImportsRenderer) generatePkgAlias(alias string) string {
	if alias == "" || alias == "." || alias == "_" {
		return alias
	}
	freeAlias := alias
	counter := 1
	_, found := r.aliasMap[freeAlias]
	for found {
		counter++
		freeAlias = alias + strconv.Itoa(counter)
	}
	r.aliasMap[freeAlias] = struct{}{}
	return freeAlias
}

func (r *ImportsRenderer) GetContext() Code {
	return r.ctx
}

func (r *ImportsRenderer) SetContext(ctx Code) {
	r.ctx = ctx
	for _, item := range r.items {
		item.SetContext(r)
	}
}

func (r *ImportsRenderer) Render(w Writer) {
	if len(r.items) == 0 {
		return
	}
	if len(r.items) == 1 {
		w.Write("import ")
		r.items[0].Render(w)
		return
	}
	w.Write("import (")
	w.Br()
	w.Indent(1)
	for i, im := range r.items {
		if i != 0 {
			w.Br()
		}
		im.Render(w)
	}
	w.Indent(-1)
	w.Br()
	w.Write(")")
}

func (r *ImportsRenderer) init() {
	if r.pathMap == nil {
		r.pathMap = map[string]string{}
	}
	if r.aliasMap == nil {
		r.aliasMap = map[string]struct{}{}
	}
}

var _ Code = (*SmartImportRenderer)(nil)

type SmartImportRenderer struct {
	alias string
	name  string
	path  string
}

func SmartImport(name string, alias string, path string) *SmartImportRenderer {
	r := &SmartImportRenderer{}
	r.SetName(name)
	r.SetAlias(alias)
	r.SetPath(path)
	return r
}

func (r *SmartImportRenderer) GetName() string {
	if r.name == "" {
		chunks := strings.Split(r.path, "/")
		return chunks[len(chunks)-1]
	}
	return r.name
}

func (r *SmartImportRenderer) SetName(name string) {
	r.name = name
}

func (r *SmartImportRenderer) GetAlias() string {
	return r.alias
}

func (r *SmartImportRenderer) SetAlias(alias string) {
	r.alias = alias
}

func (r *SmartImportRenderer) GetPath() string {
	return r.path
}

func (r *SmartImportRenderer) SetPath(path string) {
	r.path = path
}

func (r *SmartImportRenderer) GetContext() Code {
	return nil
}

func (r *SmartImportRenderer) SetContext(_ Code) {
}

func (r *SmartImportRenderer) Render(w Writer) {
	if r.alias != "" {
		w.Write(r.alias)
		w.Write(" ")
	}
	w.Write("\"")
	w.Write(r.path)
	w.Write("\"")
}

var _ Code = (*ImportRenderer)(nil)

type ImportRenderer struct {
	alias string
	path  string
}

func Import(alias string, path string) *ImportRenderer {
	r := &ImportRenderer{}
	r.SetAlias(alias)
	r.SetPath(path)
	return r
}

func (r *ImportRenderer) GetAlias() string {
	return r.alias
}

func (r *ImportRenderer) SetAlias(alias string) {
	r.alias = alias
}

func (r *ImportRenderer) GetPath() string {
	return r.path
}

func (r *ImportRenderer) SetPath(path string) {
	r.path = path
}

func (r *ImportRenderer) GetContext() Code {
	return nil
}

func (r *ImportRenderer) SetContext(_ Code) {
}

func (r *ImportRenderer) Render(w Writer) {
	if r.alias != "" {
		w.Write(r.alias)
		w.Write(" ")
	}
	w.Write("\"")
	w.Write(r.path)
	w.Write("\"")
}
