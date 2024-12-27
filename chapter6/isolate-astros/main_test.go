package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testWebRequest struct {
}

// type MockHttpClient struct {
// 	MockDo func(req *http.Request) (*http.Response, error)
// }

// func (c *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
// 	return c.MockDo(req)
// }

func (testWebRequest) FetchBytes(url string) ([]byte, error) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"number": 2}`))
	}))
	defer mockServer.Close()

	req, err := http.NewRequest(http.MethodGet, mockServer.URL, nil)
	if err != nil {
		return []byte{}, err
	}
	req.Header.Set("User-Agent", "test")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

func TestGetAstronauts(t *testing.T) {
	want := 3
	got, err := getAstronauts(testWebRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("People in space, want: %d, got: %d.", want, got)
	}
}
