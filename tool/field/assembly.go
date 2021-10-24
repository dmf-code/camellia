package field

import (
	"encoding/json"
	"fmt"
)

func Assembly(dto *struct{}, data []byte) {
	err := json.Unmarshal(data, &dto)

	if err != nil {
		fmt.Println("error: ", err)
	}
}
