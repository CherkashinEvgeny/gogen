package gen

import (
	"strconv"
	"strings"
)

var _ Sequence = (*ImportsRenderer)(nil)

type ImportsRenderer struct {
	items   []Code
	pathMap map[string]struct{}
	nameMap map[string]struct{}
	ctx     Code
}

func Imports(items ...Code) *ImportsRenderer {
	r := &ImportsRenderer{}
	r.init()
	r.Add(items...)
	return r
}

func (r *ImportsRenderer) init() {
	if r.pathMap == nil {
		r.pathMap = map[string]struct{}{}
	}
	if r.nameMap == nil {
		r.nameMap = map[string]struct{}{}
	}
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
	im, ok := item.(*ImportRenderer)
	if !ok {
		r.items = append(r.items, item)
		item.SetContext(r)
		return
	}
	_, found := r.pathMap[im.GetPath()]
	if found {
		return
	}
	r.pathMap[im.GetPath()] = struct{}{}
	if im.GetAlias() == "." || im.GetAlias() == "_" {
		r.items = append(r.items, item)
		item.SetContext(r)
		return
	}
	var name string
	if im.GetName() != "" {
		name = im.GetName()
	}
	if im.GetAlias() != "" {
		name = im.GetAlias()
	}
	if name == "" {
		r.items = append(r.items, item)
		item.SetContext(r)
		return
	}
	alias := name
	counter := 1
	_, found = r.nameMap[alias]
	for found {
		counter++
		alias = name + strconv.Itoa(counter)
	}
	r.nameMap[alias] = struct{}{}
	if name != alias {
		im.SetAlias(alias)
	}
	r.items = append(r.items, item)
	item.SetContext(r)
}

func (r *ImportsRenderer) GetContext() Code {
	return r.ctx
}

func (r *ImportsRenderer) SetContext(ctx Code) {
	r.ctx = ctx
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

var _ Code = (*ImportRenderer)(nil)

type ImportRenderer struct {
	alias string
	name  string
	path  string
}

func Import(name string, alias string, path string) *ImportRenderer {
	r := &ImportRenderer{}
	r.SetName(name)
	r.SetAlias(alias)
	r.SetPath(path)
	return r
}

func (r *ImportRenderer) GetName() string {
	if r.name == "" {
		chunks := strings.Split(r.path, "/")
		return chunks[len(chunks)-1]
	}
	return r.name
}

func (r *ImportRenderer) SetName(name string) {
	r.name = name
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
