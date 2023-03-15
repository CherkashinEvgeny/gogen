package gen

var _ Code = RawRenderer("")

type RawRenderer string

func Raw(code string) RawRenderer {
	return RawRenderer(code)
}

func (r RawRenderer) GetContext() Code {
	return nil
}

func (r RawRenderer) SetContext(_ Code) {
}

func (r RawRenderer) Render(w Writer) {
	w.Write(string(r))
}
