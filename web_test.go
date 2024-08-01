package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPProbe(t *testing.T) {
	t.Run("Return default status", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		HTTPProbe(response, request)

		got := response.Result().StatusCode
		want := 200

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
	t.Run("Return status 200", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/200", nil)
		response := httptest.NewRecorder()

		HTTPProbe(response, request)

		got := response.Result().StatusCode
		want := 200

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

	t.Run("Return status 404", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/404", nil)
		response := httptest.NewRecorder()

		HTTPProbe(response, request)

		got := response.Result().StatusCode
		want := 404

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

	t.Run("Return status non-status", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/asdf", nil)
		response := httptest.NewRecorder()

		HTTPProbe(response, request)

		got := response.Result().StatusCode
		want := 400

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}

	})

	t.Run("PATCH method", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPatch, "/asdf", nil)
		response := httptest.NewRecorder()

		HTTPProbe(response, request)

		got := response.Result().StatusCode
		want := 400

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}

	})
}
