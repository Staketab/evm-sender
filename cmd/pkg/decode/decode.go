package decode

import (
	"encoding/hex"
	"errors"
	vars "github.com/staketab/evm-sender/cmd/internal/var"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

func HexToDecSlice(hexValues ...string) ([]int64, error) {
	decValues := make([]int64, len(hexValues))

	for i, hexValue := range hexValues {
		decValue, err := strconv.ParseInt(hexValue[2:], 16, 64)
		if err != nil {
			return nil, err
		}

		decValues[i] = decValue
	}

	return decValues, nil
}

//	func HexToDec(hexValue string) (int64, error) {
//		decValue, err := strconv.ParseInt(hexValue[2:], 16, 64)
//		if err != nil {
//			vars.ErrorLog.Println("failed to convert:", err)
//			return 0, err
//		}
//		return decValue, nil
//	}
func HexToDec(hexValue string) (int64, error) {
	if hexValue == "" {
		return 0, nil
	}

	decValue, err := strconv.ParseInt(hexValue[2:], 16, 64)
	if err != nil {
		return 0, err
	}
	return decValue, nil
}

func HexExtraData(extraDataHex string) (string, error) {
	if strings.HasPrefix(extraDataHex, "0x") {
		extraDataHex = extraDataHex[2:]
	}

	extraDataBytes, err := hex.DecodeString(extraDataHex)
	if err != nil {
		vars.ErrorLog.Println(err)
		return "", err
	}
	validString := strings.ToValidUTF8(string(extraDataBytes), "?")
	reg, err := regexp.Compile("[^a-zA-Z0-9.@-_-+!(){}]+")
	if err != nil {
		return "", err
	}
	processedString := reg.ReplaceAllString(validString, " ")

	return processedString, nil
}

func HexToDecW(hexValue string) (*big.Int, error) {
	hexValue = strings.TrimPrefix(hexValue, "0x")
	decValue := new(big.Int)
	_, success := decValue.SetString(hexValue, 16)
	if !success {
		return nil, errors.New("failed to convert hex value to big.Int")
	}
	return decValue, nil
}

//func HexToStrings(hexValue string) (int64, error) {
//	var methodID string
//	switch {
//	case len(hexValue) >= 8:
//		methodID = hexValue[:8]
//	//case len(hexValue) >= 72:
//	//	toAddress = hexValue[8:72]
//	//	fallthrough
//	//case len(hexValue) >= 136:
//	//	value = hexValue[72:136]
//	//	fallthrough
//	//case len(hexValue) >= 200:
//	//	gasLimit = hexValue[136:200]
//	//	fallthrough
//	//case len(hexValue) >= 264:
//	//	isCreation = hexValue[200:264]
//	//	fallthrough
//	//case len(hexValue) >= 328:
//	//	data = hexValue[264:]
//	default:
//		vars.InfoLog.Println("hexValue is too short")
//	}
//
//	method, err := HexToDec(methodID)
//	if err != nil {
//		vars.ErrorLog.Println("method is too short", err)
//	}
//	return method, nil
//}
