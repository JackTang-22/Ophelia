package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/10/7 00:10
 * @Description:
 */

type State struct {
	Name     string
	Area     int
	Water    float64
	Senators []interface{}
}

func TestJson(t *testing.T) {

	ma := []byte("{\"name\": \"golang\", \"area\": 27336, \"water\": 26.7, \"senators\":[\"Tom\", \"Mike\"]}")
	var state State
	if err := json.Unmarshal(ma, &state); err != nil {
        fmt.Println(err)
	} else {
		fmt.Println(state)
	}
}
