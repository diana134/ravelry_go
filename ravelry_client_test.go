package main

import "testing"

func TestGetClient(t *testing.T) {

	t.Run("creating a Ravelry client with the given credentials", func(t *testing.T) {
		testCredentials := RavelryCredentials{
			AuthUsername: "foo",
			AuthPassword: "bar",
		}
		wantAuthString := "foo:bar"
		wantAuthHeader := "Basic Zm9vOmJhcg=="

		got := GetRavelryClient(&testCredentials)

		if got.authString != wantAuthString {
			t.Errorf("got %q want %q", got, wantAuthString)
		}

		if got.authHeader != wantAuthHeader {
			t.Errorf("got %q want %q", got, wantAuthHeader)
		}
	})

}

func TestBuildParameterString(t *testing.T) {

	t.Run("construct a parameter string from the given availability type and sort type", func(t *testing.T) {
		availabilityType := Parameter{
			urlKey:    "availability",
			urlValue:  "free",
			tweetText: "free",
		}

		sortType := Parameter{
			urlKey:    "sort",
			urlValue:  "recently-popular",
			tweetText: "hottest",
		}

		got := BuildParameterString(availabilityType, sortType)

		want := "?page_size=1&availability=free&sort=recently-popular"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
