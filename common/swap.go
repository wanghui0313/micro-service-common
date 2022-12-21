package common

import (
	"encoding/json"
)

func SwapTo(req, res any) (err error) {
	dataByte, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(dataByte, res)
	return
}
