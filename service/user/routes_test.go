package user

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/irviner26/ecom/types"
	"github.com/julienschmidt/httprouter"
)

type mockUserStore struct {
}

func (m *mockUserStore) GetUserByEmail(email string, ctx context.Context) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) GetUserByID(id int, ctx context.Context) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(user types.User, ctx context.Context) error {
	return nil
}

func TestUserServiceHandlers(t *testing.T) {
	mockStore := mockUserStore{}
	testHandler := NewHandler(&mockStore)

	t.Run("FAIL IF INVALID PAYLOAD", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "random",
			LastName:  "llll",
			Email:     "aed@mail.com",
			Password:  "12345678",
		}
		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		router := httprouter.New()
		router.POST("/register", testHandler.handleRegister)

		request := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, request)

		response := recorder.Result()
		bytes, err := io.ReadAll(response.Body)
		if err != nil {
			t.Fatal(err)
		}

		if response.StatusCode != http.StatusCreated {
			t.Error("Error with status:", response.StatusCode, response.Status)
			t.Error("Message: ", string(bytes))
		}
	})
}
