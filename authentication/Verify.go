package authentication

import (
	"net/http"
	"fmt"
	"bp-auth-interceptor/helpers"
)

// Verify verifica se un la richiesta contiene un jwt valido, ritorna 200 in caso di successo, 401
func Verify (w http.ResponseWriter, r *http.Request) {
	token := helpers.ExtractToken(r)
	success, err := helpers.VerifyToken(token)


	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, err.Error());
		return;
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, success)
}