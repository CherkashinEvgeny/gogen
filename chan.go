package gen

type ChanDir int

const (
	SendAndReceive ChanDir = 0
	Receive        ChanDir = 1
	Send           ChanDir = 2
)

var _ Code = (*ChanRenderer)(nil)

type ChanRenderer struct {
	dir  ChanDir
	elem Code
	ctx  Code
}

func Chan(dir ChanDir, elem Code) *ChanRenderer {
	r := &ChanRenderer{}
	r.SetDir(dir)
	r.SetElem(elem)
	return r
}

func (r *ChanRenderer) GetDir() ChanDir {
	return r.dir
}

func (r *ChanRenderer) SetDir(dir ChanDir) {
	r.dir = dir
}

func (r *ChanRenderer) GetElem() Code {
	return r.elem
}

func (r *ChanRenderer) SetElem(elem Code) {
	r.elem = elem
	if elem != nil {
		elem.SetContext(r)
	}
}

func (r *ChanRenderer) GetContext() Code {
	return r.ctx
}

func (r *ChanRenderer) SetContext(ctx Code) {
	r.ctx = ctx
	if r.elem != nil {
		r.elem.SetContext(r)
	}
}

func (r *ChanRenderer) Render(w Writer) {
	switch r.dir {
	case Receive:
		w.Write("<-chan ")
	case Send:
		w.Write("chan<- ")
	default:
		w.Write("chan ")
	}
	r.elem.Render(w)
}
