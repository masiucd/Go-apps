package lib

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

// In a real app the secret would be stored in an environment variable - os.Getenv("SECRET")
var Store = sessions.NewCookieStore([]byte("secret"))

type SessionData struct {
	ID    string
	Email string
}

func GetSessionData(w http.ResponseWriter, r *http.Request) (*SessionData, error) {
	session, err := Store.Get(r, "session")
	if err != nil {
		return nil, err
	}
	email, ok := session.Values["email"]
	if !ok {
		return nil, fmt.Errorf("no valid session")
	}
	id, ok := session.Values["ID"]
	if !ok {
		return nil, fmt.Errorf("no valid session")
	}

	return &SessionData{
		ID:    id.(string),
		Email: email.(string),
	}, nil
}
