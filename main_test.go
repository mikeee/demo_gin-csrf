package main

import (
	"time"
	"io/ioutil"
	"net/http/cookiejar"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestMain(t *testing.T) {
	go main()

	// Sleep to allow gin time to start
	time.Sleep(time.Second)

	cookies, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: cookies,
	}
	getResp, err := client.Get("http://127.0.0.1:8000/protected")

	if err != nil {
		t.Fatal("unsuccessful get request to the server")
	}

	defer getResp.Body.Close()

	body, err := ioutil.ReadAll(getResp.Body)

	if err != nil {
		t.Fatal("failed to parse body")
	}

	bodyString := string(body)

	sessionCookies := getResp.Cookies()
	

	postReq, _ := http.NewRequest("POST", "http://127.0.0.1:8000/protected", nil)
	cookies.SetCookies(postReq.URL, sessionCookies)
	postReq.Header.Set("X-CSRF-TOKEN", bodyString)

	postResp, err := client.Do(postReq)

	if err != nil {
		t.Fatal("unsuccessful post request to the server")
	}

	assert.Equal(t, http.StatusOK, postResp.StatusCode)
}
