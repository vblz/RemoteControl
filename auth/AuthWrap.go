package auth

import "net/http"

// Wrapper is the wrapper on http.HandleFunc for provide Basic HTTP Auth
type Wrapper struct {
	User     string
	Password string
	Realm    string
}

// Wrap the http.HandleFunc
func (wrapper Wrapper) Wrap(wrapped http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO add the authentification check
		wrapped(w, r)
	}
}
