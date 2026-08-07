package service

import (
	"bufio"
	"bytes"
	"codewars-pretty-stats/internal/config"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func makeRequest(size string, username string) *http.Response {
	cfg := &config.AxiomConfig{
		AxiomURL:   "invalid",
		AxiomToken: "invalid",
	}

	rec := httptest.NewRecorder()

	params := url.Values{}
	params.Add("size", size)
	params.Add("username", username)

	target := "/?" + params.Encode()

	req := httptest.NewRequest(http.MethodGet, target, nil)

	Svg(cfg)(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	return res
}

func containsSVG(resp *http.Response) bool {
	reader := bufio.NewReader(resp.Body)

	peekBytes, err := reader.Peek(512)
	if err != nil && err != io.EOF {
		return false
	}

	trimmed := bytes.TrimSpace(peekBytes)
	hasSVG := bytes.HasPrefix(bytes.ToLower(trimmed), []byte("<?xml"))

	resp.Body = struct {
		io.Reader
		io.Closer
	}{
		Reader: reader,
		Closer: resp.Body,
	}

	return hasSVG
}

// ---

func TestSVGValid(t *testing.T) {
	res := makeRequest("1", "Selector0073")

	if res.StatusCode != http.StatusOK {
		t.Error("Valid request returned error. Code:" + res.Status)
	}

	if containsSVG(res) != true {
		t.Error("SVG not found. Code:" + res.Status)
	}
}

func TestSVGInvalidSize(t *testing.T) {
	res := makeRequest("one", "Selector0073")

	if res.StatusCode == http.StatusOK {
		t.Error("Invalid size returned OK. Code:" + res.Status)
	}

	if containsSVG(res) == true {
		t.Error("SVG not found. Code:" + res.Status)
	}
}

func TestSVGInvalidUser(t *testing.T) {
	res := makeRequest("1", "_unexisted^user_")

	if res.StatusCode == http.StatusOK {
		t.Error("Invalid user returned OK. Code:" + res.Status)
	}

	if containsSVG(res) == true {
		t.Error("SVG not found. Code:" + res.Status)
	}
}

func TestSVGInvalidBoth(t *testing.T) {
	res := makeRequest("one", "_unexisted^user_")

	if res.StatusCode == http.StatusOK {
		t.Error("Invalid user returned OK. Code:" + res.Status)
	}

	if containsSVG(res) == true {
		t.Error("SVG not found. Code:" + res.Status)
	}
}
