package handles

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)



func Test_RootGET(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	Status(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	if !strings.Contains(string(body), "COUNT: 1") {
		t.Fatalf("GET on root failed")
	}

}
