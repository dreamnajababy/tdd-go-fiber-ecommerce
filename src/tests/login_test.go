package tests

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	app := SetupLoginTest()

	var got JwtResponse

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

	if err != nil {
		t.Errorf("cannot parse response body.%v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("expect status code 200, got %v", resp.StatusCode)
	}

	if got.Msg != expectedMessage {
		t.Errorf("expected message: %v, got %v", expected.Msg, got.Msg)
	}
	if got.Token == "" {
		t.Errorf("expect token is not empty but got %v", got.Token)
	}
}
