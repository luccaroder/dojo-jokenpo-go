package middlewares

import (
	"fmt"
	"net/http"

	"github.com/go-playground/form"
)

type RequestParams struct {
	Player1 *int `json:"player1,omitempty"`
	Player2 *int `json:"player2,omitempty"`
}

func (r RequestParams) IsValid() (bool, error) {
	if r.Player1 == nil {
		return false, fmt.Errorf("player1 is missing")
	}

	if r.Player2 == nil {
		return false, fmt.Errorf("player2 is missing")
	}

	return true, nil
}

func ValidateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var rp RequestParams

		decoder := form.NewDecoder()
		err := decoder.Decode(&rp, r.URL.Query())
		fmt.Println(err)
		fmt.Println(rp.IsValid())

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Os parâmetros player1 e player2 são obrigatórios")
			return
		}
		next.ServeHTTP(w, r)
	})
}
