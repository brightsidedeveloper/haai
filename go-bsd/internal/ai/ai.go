package ai

type AI struct {
}

func New() *AI {
	return &AI{}
}

func (ai *AI) Prompt(prompt string) (string, error) {
	return "Woah", nil
}
