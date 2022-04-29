package style

var Bold = NewTransformer(&bold{})

type bold struct {
}

func (b *bold) LowerFunc(src uint8) []byte {
	return nil
}
func (b *bold) UpperFunc(src uint8) []byte {
	return nil
}
func (b *bold) DigitFunc(src uint8) []byte {
	return nil
}
