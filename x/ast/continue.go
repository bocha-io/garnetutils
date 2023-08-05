package ast

const Continue = "Continue"

func (a *ASTConverter) processContinue(data []byte) (string, error) {
	_ = data
	return "continue", nil
}
