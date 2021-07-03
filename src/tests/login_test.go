package tests

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	models "github.com/dreamnajababy/go-ecom/src/models"
)

func handleBodyParsing(t testing.TB, err error) {
	if err != nil {
		t.Errorf("cannot parse response body.%v", err)
	}
}
func TestLogin(t *testing.T) {
	app := SetupLoginTest()

	t.Run("user login success with correct username and password.", func(t *testing.T) {
		var got map[string]interface{}

		bytes, _ := json.Marshal(models.CorrectUser)
		payloadReader := strings.NewReader(string(bytes))

		request := httptest.NewRequest("POST", "/login", payloadReader)
		request.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(request)

		err := json.NewDecoder(resp.Body).Decode(&got)
		handleBodyParsing(t, err)

		assertStatusCode(t, 200, resp.StatusCode) // assert HTTP Response Status Code
		assertMessage(t, "login successfully.", got)
		assertToken(t, got)
	})

	t.Run("user login fail with correct username but incorrect password.", func(t *testing.T) {
		var got map[string]interface{}

		bytes, _ := json.Marshal(models.IncorrectUser)
		payloadReader := strings.NewReader(string(bytes))

		request := httptest.NewRequest("POST", "/login", payloadReader)
		request.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(request)

		err := json.NewDecoder(resp.Body).Decode(&got)
		handleBodyParsing(t, err)

		assertStatusCode(t, 401, resp.StatusCode) // assert HTTP Response Status Code
		assertMessage(t, "login unsuccessfully.", got)
	})

}

func assertMessage(t testing.TB, msg string, got map[string]interface{}) {
	t.Helper()
	if got["msg"] != msg {
		t.Errorf("expected message: %v, got %v", msg, got["msg"])
	}
}
func assertToken(t testing.TB, got map[string]interface{}) {
	t.Helper()
	if got["token"] == "" {
		t.Errorf("expect token is not empty but got %v", got["token"])
	}
}
