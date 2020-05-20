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
