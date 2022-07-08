package another

import "fmt"

type SomeVeryComplicatedFuncConfig struct {
	A uint8
	B uint8
	C uint8
	D uint8
}

// SomeVeryComplicatedFunc does domething
func SomeVeryComplicatedFunc(cfg SomeVeryComplicatedFuncConfig) int {
	res := int(cfg.A) + int(cfg.B) + int(cfg.C) - int(cfg.D)

	return res
}

// DoNothing does nothing
func DoNothing() (string, error) {
	res := fmt.Sprintf("%+#v", &SomeVeryComplicatedFuncConfig{})

	return res, nil
}
