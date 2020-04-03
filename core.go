package gopkgtest

import "github.com/pkg/errors"

func Hoge(err error) error {
	return errors.Wrap(err, "hoge")
}
