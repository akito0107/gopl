package params

import (
	"net/http/httptest"
	"testing"
)

func TestUnpack(t *testing.T) {
	t.Run("valid email", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/?email=test@example.com", nil)
		data := struct {
			Email string `http:"email" validate:"email"`
		}{}

		err := Unpack(req, &data)
		if err != nil {
			t.Errorf("valid email")
		}
	})
	t.Run("valid email", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/?email=testexample.com", nil)
		data := struct {
			Email string `http:"email" validate:"email"`
		}{}

		err := Unpack(req, &data)
		if err == nil {
			t.Errorf("invalid email")
		}
	})
}
