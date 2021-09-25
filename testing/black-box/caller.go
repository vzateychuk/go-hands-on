package main

import (
	"fmt"
	"golang.org/x/xerrors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Caller interface {
	Call(req map[string]interface{}) (io.ReadCloser, error)
}

// restCaller is a convenience type for GET requests to REST endpoints.
type restCaller string

// Call implements Caller for the restCaller type.
func (rc restCaller) Call(req map[string]interface{}) (io.ReadCloser, error) {
	var params = make(url.Values)
	for k, v := range req {
		params.Set(k, fmt.Sprint(v))
	}

	url := fmt.Sprintf("%s?%s", string(rc), params.Encode())
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		drainAndClose(res.Body)
		return nil, xerrors.Errorf("unexpected response status code: %d", res.StatusCode)
	}

	return res.Body, nil
}

func drainAndClose(r io.ReadCloser) {
	if r == nil {
		return
	}
	_, _ = io.Copy(ioutil.Discard, r)
	_ = r.Close()
}
