package encode

import "fmt"

func DecToHex(decValue int64) string {
	hexValue := fmt.Sprintf("%X", decValue)
	return "0x" + hexValue
}
