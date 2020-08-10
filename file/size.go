package file

import "fmt"

type Size int

const (
	Byte  Size = 1
	KByte      = 1024 * Byte
	MByte      = 1024 * KByte
	GByte      = 1024 * MByte
	TByte      = 1024 * GByte
)

func ParseSize(s string) (size Size, err error) {

	oneSize := 0
	unit := Size(1)

	for i := 0; i < len(s); i++ {

		if s[i] >= '0' && s[i] <= '9' {
			oneSize = oneSize*10 + int(s[i]-'0')
			if i+1 == len(s) {
				size += Size(oneSize * int(unit))
				return
			}
			continue
		}

		switch {
		case s[i] == 'B' || s[i] == 'b':
			unit = Byte
		case s[i] == 'K' || s[i] == 'k':
			unit = KByte
		case s[i] == 'M' || s[i] == 'm':
			unit = MByte
		case s[i] == 'G' || s[i] == 'g':
			unit = GByte
		case s[i] == 'T' || s[i] == 't':
			unit = TByte
		default:
			return Size(0), fmt.Errorf("Invalid value of function arg:%s", s)
		}

		if int(unit) > 0 {
			// peek next element
			if i+1 == len(s) {
				size += Size(oneSize * int(unit))
				return
			}

			if s[i+1] == 'B' || s[i+1] == 'b' {
				i++
				size += Size(oneSize * int(unit))
				oneSize = 0
				unit = Byte
				continue
			}

			if s[i+1] >= '0' && s[i+1] <= '9' {
				size += Size(oneSize * int(unit))
				oneSize = 0
				unit = Byte
			}
		}
	}

	return
}
