package main

import (
	"net/http"
	"net/url"
	"testing"
)

func TestRun(t *testing.T) {
	go api()
}

func TestLoginFail(t *testing.T) {
	go api()

	username := "test"

	resp, err := http.PostForm("http://localhost:8081/login", url.Values{
		"username": {username},
	})
	defer resp.Body.Close()

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Fatal("Unexpected Status Code: ", resp.StatusCode)
	}
}

func TestRegisterAndLogin(t *testing.T) {
	go api()

	username := "test"

	resp, err := http.PostForm("http://localhost:8081/register", url.Values{
		"username": {username},
		"elderly":  {"true"},
	})
	defer resp.Body.Close()

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Fatal("Unexpected Status Code: ", resp.StatusCode)
	}

	resp, err = http.PostForm("http://localhost:8081/login", url.Values{
		"username": {username},
	})
	defer resp.Body.Close()

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Unexpected Status Code: ", resp.StatusCode)
	}
}

func TestRegisterAndGet(t *testing.T) {
	go api()

	username := "test2"

	resp, err := http.PostForm("http://localhost:8081/register", url.Values{
		"username": {username},
		"elderly":  {"true"},
	})
	defer resp.Body.Close()

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Fatal("Unexpected Status Code: ", resp.StatusCode)
	}

	resp, err = http.Get("http://localhost:8081/user/1")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Unexpected Status Code: ", resp.StatusCode)
	}

}
