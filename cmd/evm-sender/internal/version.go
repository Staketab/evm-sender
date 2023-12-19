package internal

import (
	"encoding/json"
	"fmt"
	"github.com/staketab/evm-sender/cmd/evm-sender/internal/var"
)

func getVersion() {
	data := struct {
		Version string `json:"version"`
	}{
		Version: vars.Version,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		vars.ErrorLog.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}
