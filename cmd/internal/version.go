package internal

import (
	"encoding/json"
	"fmt"
	vars "github.com/staketab/evm-sender/cmd/internal/var"
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
