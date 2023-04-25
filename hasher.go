package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func Hash(path string) (string, error) {
	_, e := url.Parse(path)
	if e != nil {
		return "", e
	}
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	rsp, e := client.Get(path)
	if e != nil {
		return "", e
	}
	if rsp.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("not valid response statusCode: %d", rsp.StatusCode))
	}
	defer rsp.Body.Close()
	h := md5.New()
	if _, e = io.Copy(h, rsp.Body); e != nil {
		return "", e
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
