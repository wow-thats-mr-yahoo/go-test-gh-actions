package lib

import "github.com/wow-thats-mr-yahoo/go-test-gh-actions/internal/another"

func UseAnother() (int, error) {
	cfg := another.SomeVeryComplicatedFuncConfig{
		A: 1,
		B: 32,
		C: 32,
	}

	res := another.SomeVeryComplicatedFunc(cfg)

	return res, nil
}
