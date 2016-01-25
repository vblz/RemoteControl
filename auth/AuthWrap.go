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
		if wrapper.isOk(r) {
			wrapped(w, r)
		} else {
			w.Header().Set("WWW-Authenticate", "Basic realm=\""+wrapper.Realm+"\"")
			http.Error(w, "Unauthorized", 401)
		}
	}
}

func (wrapper Wrapper) isOk(r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		return false
	}
	if username != wrapper.User {
		return false
	}
	return password == wrapper.Password
}
