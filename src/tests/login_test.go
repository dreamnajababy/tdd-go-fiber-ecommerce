package tests

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	app := SetupLoginTest()

	var got map[string]interface{}

	credential := struct {
		username string
		password string
	}{
		username: "dreamnajababy", password: "1234",
	}

	bytes, err := json.Marshal(credential)
	payloadReader := strings.NewReader(string(bytes))

	if err != nil {
		t.Errorf("something wrong, %v", err)
	}

	request := httptest.NewRequest("POST", "/login", payloadReader)
	request.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(request)

	err = json.NewDecoder(resp.Body).Decode(&got)

	if err != nil {
		t.Errorf("cannot parse response body.%v", err)
	}

	assertStatusCode(t, 200, resp.StatusCode) // assert HTTP Response Status Code
	assertTokenAndMessage(t, "login successfully.", got)

}

func assertTokenAndMessage(t testing.TB, msg string, got map[string]interface{}) {
	t.Helper()
	if got["msg"] != msg {
		t.Errorf("expected message: %v, got %v", msg, got["msg"])
	}
	if got["token"] == "" {
		t.Errorf("expect token is not empty but got %v", got["token"])
	}
}
