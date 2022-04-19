package module

import (
	"github.com/pkg/errors"
)

func (m module) Tidy() error {
	if output, err := m.execute("tidy"); err != nil {
		return errors.Wrap(err, output.Combined)
	}
	return nil
}
