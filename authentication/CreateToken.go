package authentication

import (
	"net/http"
	"bp-auth-interceptor/helpers"
	"fmt"
)


// CreateToken generate un nuovo token e lo manda all'utente come stringa
func CreateToken(w http.ResponseWriter, r *http.Request)  {
	token, err := helpers.GenerateToken("mauro")
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, token)
}