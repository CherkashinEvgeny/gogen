package gen

var _ Code = (*MethodRenderer)(nil)

type MethodRenderer struct {
	receiver  Code
	name      string
	signature Code
	body      Code
	ctx       Code
}

func Method(receiver Code, name string, signature Code, body Code) *MethodRenderer {
	r := &MethodRenderer{}
	r.SetReceiver(receiver)
	r.SetName(name)
	r.SetSignature(signature)
	r.SetBody(body)
	return r
}

func (r *MethodRenderer) GetReceiver() Code {
	return r.receiver
}

func (r *MethodRenderer) SetReceiver(receiver Code) {
	r.receiver = receiver
	if receiver != nil {
		receiver.SetContext(r)
	}
}

func (r *MethodRenderer) GetName() string {
	return r.name
}

func (r *MethodRenderer) SetName(name string) {
	r.name = name
}

func (r *MethodRenderer) GetSignature() Code {
	return r.signature
}

func (r *MethodRenderer) SetSignature(signature Code) {
	r.signature = signature
	if signature != nil {
		signature.SetContext(r)
	}
}

func (r *MethodRenderer) GetBody() Code {
	return r.signature
}

func (r *MethodRenderer) SetBody(body Code) {
	r.body = body
	if body != nil {
		body.SetContext(r)
	}
}

func (r *MethodRenderer) GetContext() Code {
	return r.ctx
}

func (r *MethodRenderer) SetContext(ctx Code) {
	r.ctx = ctx
}

func (r *MethodRenderer) Render(w Writer) {
	w.Write("func ")
	w.Write("(")
	r.receiver.Render(w)
	w.Write(")")
	w.Write(" ")
	w.Write(r.name)
	r.signature.Render(w)
	w.Write(" {")
	w.Br()
	w.Indent(1)
	r.body.Render(w)
	w.Indent(-1)
	w.Br()
	w.Write("}")
}

var _ Code = (*ReceiverRenderer)(nil)

type ReceiverRenderer struct {
	name  string
	ttype Code
	ctx   Code
}

func Receiver(name string, ttype Code) *ReceiverRenderer {
	r := &ReceiverRenderer{}
	r.SetName(name)
	r.SetType(ttype)
	return r
}

func (r *ReceiverRenderer) GetName() string {
	return r.name
}

func (r *ReceiverRenderer) SetName(name string) {
	r.name = name
}

func (r *ReceiverRenderer) GetType() Code {
	return r.ttype
}

func (r *ReceiverRenderer) SetType(ptype Code) {
	r.ttype = ptype
	if ptype != nil {
		ptype.SetContext(r)
	}
}

func (r *ReceiverRenderer) GetContext() Code {
	return r.ctx
}

func (r *ReceiverRenderer) SetContext(ctx Code) {
	r.ctx = ctx
}

func (r *ReceiverRenderer) Render(w Writer) {
	if r.name != "" {
		w.Write(r.name)
		w.Write(" ")
	}
	r.ttype.Render(w)
}
