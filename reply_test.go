package xhttp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_JSONResponseMessage_tooLongMessage(t *testing.T) {
	respR := httptest.NewRecorder()

	var replyData = ReplyWrapper{
		ID:      "id1",
		Context: "da context",
	}

	message := `These are details.  These are details. These are details. These are
details. These are details.  These are details.  These are details. These
are details. These are details. These are details. These are details.
These are details. These are details. These are details. These are details.
These are details.  These are details. These are details. These are
details. These are details.  These are details.  These are details. These
are details. These are details. These are details. These are details.
These are details. These are details. These are details. These are details.
These are details.`

	err := JSONResponseMessage(respR, http.StatusNotFound, replyData, message)
	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	var respData struct {
		ReplyWrapper
		Message string
	}

	fmt.Println(respR.Body.String())
	err = json.Unmarshal(respR.Body.Bytes(), &respData)
	if err != nil {
		t.Errorf("error unmarshaling data: %s", err)
	}

	if len(respData.Message) > 256 {
		t.Errorf("error, details are too long; expected 256 characters or less, but got %d", len(respData.Message))
	}
}
