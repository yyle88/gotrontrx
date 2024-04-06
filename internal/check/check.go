package check

import (
	"github.com/pkg/errors"
)

func Done(err error) {
	if err != nil {
		panic(errors.WithMessage(err, "wrong"))
	}
}
