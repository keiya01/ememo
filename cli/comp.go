package cli

type CompFlag struct {
	Value string
}

func NewCompFlag(value string) *CompFlag {
	return &CompFlag{Value: value}
}

func (c *CompFlag) FlagAction() error {
	return nil
}
