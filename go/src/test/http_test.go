package test

import (
	"io/ioutil"
	"net/http"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/9/19 15:48
 * @Description:
 */

func TestHttp(t *testing.T) {
	response, err := http.Get("http://www.baidu.com")
	if err != nil {
		t.Log(err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	t.Log(string(body))
}
