package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
)

type RequestParams struct {
	Player1 *string `schema:"player1"`
	Player2 *string `schema:"player2"`
}

func (r RequestParams) IsValid() error {
	if r.Player1 == nil {
		return fmt.Errorf("o parâmetro player1 é obrigatório")
	}

	if r.Player2 == nil {
		return fmt.Errorf("o parâmetro player2 é obrigatório")
	}

	return nil
}

func ValidateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var rp RequestParams

		decoder := schema.NewDecoder()
		err := decoder.Decode(&rp, r.URL.Query())

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Os parâmetros estão incorretos")
			return
		}

		err = rp.IsValid()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
			return
		}
		next.ServeHTTP(w, r)
	})
}
