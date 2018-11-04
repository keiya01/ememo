package cli

import (
	"github.com/keiya01/ememo/input"
)

type CompFlag struct {
	Value string
}

func NewCompFlag(value string) *CompFlag {
	fileName := input.AddExtension(value)

	return &CompFlag{Value: fileName}
}

func (c *CompFlag) FlagAction() error {
	return nil
}
