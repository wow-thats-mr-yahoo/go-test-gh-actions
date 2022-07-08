package another

type SomeVeryComplicatedFuncConfig struct {
	A uint8
	B uint8
	C uint8
}

func SomeVeryComplicatedFunc(cfg SomeVeryComplicatedFuncConfig) int {
	res := int(cfg.A) + int(cfg.B) + int(cfg.C)

	return res
}
