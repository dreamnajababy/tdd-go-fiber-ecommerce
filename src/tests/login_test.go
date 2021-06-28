package tests

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
)

type HttpResponse struct {
	msg    string
	status int
}

func TestLogin(t *testing.T) {
	app := SetupLoginTest()
	var got HttpResponse
	expected := HttpResponse{
		msg:    "login successfully.",
		status: 200,
	}
	credential := struct {
		username string
		password string
	}{
		username: "dreamnajababy",
		password: "1234",
	}

	bytes, err := json.Marshal(credential)
	payloadReader := strings.NewReader(string(bytes))

	if err != nil {
		t.Errorf("something wrong, %v", err)
	}

	request := httptest.NewRequest("POST", "/login", payloadReader)
	request.Header.Set("Content-Type", "application/json") // need to set header for using json body parser
	resp, _ := app.Test(request)

	err = json.NewDecoder(resp.Body).Decode(&got)
	if resp.StatusCode != 200 {
		t.Errorf("expect status code 200, got %v", resp.StatusCode)
	}
	assertStruct(t, expected, got, err)
}
