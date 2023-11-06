package ast

const Break = "Break"

func (a *Converter) processBreak(data []byte) (string, error) {
	_ = data
	return "break", nil
}
