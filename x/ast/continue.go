package ast

const Continue = "Continue"

func (a *Converter) processContinue(data []byte) (string, error) {
	_ = data
	return "continue", nil
}
