package another

type SomeVeryComplicatedFuncConfig struct {
	A uint8
	B uint8
	C uint8
	D uint8
}

func SomeVeryComplicatedFunc(cfg SomeVeryComplicatedFuncConfig) int {
	res := int(cfg.A) + int(cfg.B) + int(cfg.C) - int(cfg.D)

	return res
}

const nothing = "nothing"

// DoNothing does nothing
func DoNothing() (string, error) {
	return nothing, nil
}
