package rest

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func HasValidContent(resp *http.Response) bool {
	if resp != nil {
		code := resp.StatusCode
		return codeOK(code) && Is2xx(resp) &&
			code != http.StatusNoContent &&
			resp.Body != nil
	}

	return false
}

func Is4xx(resp *http.Response) (bool, error) {
	code := resp.StatusCode
	if code >= http.StatusBadRequest && code <= 499 {
		return true, errors.New(fmt.Sprintf("Received HTTP%d: %s", code, readerToString(resp.Body)))
	}

	return false, nil
}

func Is5xx(resp *http.Response) (bool, error) {
	code := resp.StatusCode
	if code >= http.StatusInternalServerError && code <= 599 {
		return true, errors.New(fmt.Sprintf("Received HTTP%d: %s", code, readerToString(resp.Body)))
	}

	return false, nil
}

func Is2xx(resp *http.Response) bool {
	code := resp.StatusCode
	return code >= http.StatusOK && code <= 299
}

func codeOK(code int) bool {
	return code != 0 && code <= 599
}

func readerToString(r io.ReadCloser) string {
	buffer := new(bytes.Buffer)
	if _, err := buffer.ReadFrom(r); err != nil {
		return "Error converting io.ReadCloser to string"
	} else {
		return buffer.String()
	}
}
