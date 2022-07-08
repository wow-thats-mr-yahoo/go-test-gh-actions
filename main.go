package main

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func main() {
	err := fmt.Errorf("OKOK")

	err = errors.Wrap(err, "something")

	logrus.WithError(err).Info("something has gone sidewise")
}
