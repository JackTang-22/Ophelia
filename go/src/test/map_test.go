package test

import (
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/9/15 09:14
 * @Description:
 */

// map在64位机器上传递需要8字节，在32位机器上需要4字节

type PersonInfo struct {
	Id      string
	Name    string
	Address string
}

func TestMap(t *testing.T) {
	var personInfo map[string]PersonInfo
	personInfo = make(map[string]PersonInfo)
	personInfo["1"] = PersonInfo{Id: "1", Name: "name", Address: "address"}
	if res, ok := personInfo["1"]; ok {
		t.Fatal(res)
	} else {
		t.Fatal("Not Found")
	}

}
