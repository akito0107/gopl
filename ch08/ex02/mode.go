package main

type TransferMode int

const (
	ASCII TransferMode = iota
	BIN
	UNKNOWN
)

func (t TransferMode) String() string {
	switch t {
	case ASCII:
		return "Ascii"
	case BIN:
		return "Binary"
	default:
		return "UNKNOWN"
	}
}

func FromCode(code string) TransferMode {
	switch code {
	case "A":
		return ASCII
	case "I":
		return BIN
	default:
		return UNKNOWN
	}
}
