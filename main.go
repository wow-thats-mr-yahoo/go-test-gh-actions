package main

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/wow-thats-mr-yahoo/go-test-gh-actions/internal/lib"
)

func main() {
	err := fmt.Errorf("OKOK")

	err = errors.Wrap(err, "something")

	logrus.WithError(err).Error("something has gone sidewise")

	a := lib.DoStuff()

	b, err := lib.UseAnother()
	if err != nil {
		logrus.WithError(err).Fatalln("fail")
	}

	logrus.WithFields(logrus.Fields{
		"A": a,
		"B": b,
	}).Info("metrics")
}
