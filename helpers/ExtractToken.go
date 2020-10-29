package helpers

import "net/http"

// ExtractToken estrae un token dalla richiesta (header: access_token o query: token)
func ExtractToken(r *http.Request) string {
	token := r.Header.Get("access_token")
	if len(token) < 1 {
		token = r.URL.Query().Get("token")
	}
	
	return token;
}
