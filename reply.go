package xhttp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ReplyWrapper struct {
	ID      string `json:"id"`
	Context string `json:"context"`

	Meta interface{} `json:"meta,omitempty"`
}

func JSONResponseMessage(w http.ResponseWriter, statusCode int, replyData ReplyWrapper, message string) error {

	if codeText := http.StatusText(statusCode); codeText == "" {
		return fmt.Errorf("%w: expected valid Status Code integer, but got %d", ErrInvalidStatusCode, statusCode)
	}

	reply := struct {
		ReplyWrapper
		Message string `json:"message"`
	}{
		ReplyWrapper: replyData,
		Message:      fmt.Sprintf("%.256s", message),
	}

	reply.Context = fmt.Sprintf("%.256s", replyData.Context)
	if reply.Message == "" {
		reply.Message = http.StatusText(statusCode)
	}

	jsonReply, err := json.Marshal(reply)
	if err != nil {
		return fmt.Errorf("%w: expected json, but got %q", ErrFailedToMarshalJSON, err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	w.WriteHeader(statusCode)
	_, err = w.Write(jsonReply)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrWritingResponse, err)
	}

	return nil
}

func JSONResponseData(w http.ResponseWriter, statusCode int, replyData ReplyWrapper, data interface{}) error {

	if codeText := http.StatusText(statusCode); codeText == "" {
		return fmt.Errorf("%w: expected valid Status Code integer, but got %d", ErrInvalidStatusCode, statusCode)
	}

	if data == nil {
		return fmt.Errorf("%w: expected non-nil data value, but got nil", ErrDataOmitted)
	}

	reply := struct {
		ReplyWrapper
		Data interface{} `json:"data"`
	}{
		ReplyWrapper: replyData,
		Data:         data,
	}

	reply.Context = fmt.Sprintf("%.256s", replyData.Context)

	jsonReply, err := json.Marshal(reply)
	if err != nil {
		return fmt.Errorf("error converting to json: %s", err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	w.WriteHeader(statusCode)
	_, err = w.Write(jsonReply)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrWritingResponse, err)
	}

	return nil
}
