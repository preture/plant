package controllers

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestPost(t *testing.T) {
	v := url.Values{}
	v.Add("login", "admin")
	v.Add("password", "111111")
	client := &http.Client{}
	request, _ := http.NewRequest("POST", "/join", strings.NewReader(v.Encode()))
	response, _ := client.Do(request)
	if response.StatusCode != 403 {
		t.Errorf("Join User(admin,111111) failed. Got %d, expected 403", response.StatusCode)
	}
}
