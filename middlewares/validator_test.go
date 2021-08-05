// package middlewares

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestValidatorMiddlewareMissingPlayer1(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/jokenpo", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rr := httptest.NewRecorder()
// 	emptyValidatorMiddleware := func(w http.ResponseWriter, r *http.Request) {
// 		assert.Nil(t, err)
// 		assert.EqualError(t, err, "o parâmetro player1 é obrigatório")
// 	}
// 	ValidateMiddleware(http.HandlerFunc(emptyValidatorMiddleware)).ServeHTTP(rr, req)
// }

// func TestValidatorMiddlewareMissingPlayer2(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/jokenpo?player1=pedra", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rr := httptest.NewRecorder()
// 	emptyValidatorMiddleware := func(w http.ResponseWriter, r *http.Request) {
// 		assert.Nil(t, err)
// 		assert.EqualError(t, err, "o parâmetro player2 é obrigatório")
// 	}
// 	ValidateMiddleware(http.HandlerFunc(emptyValidatorMiddleware)).ServeHTTP(rr, req)
// }

// func TestValidatorMiddlewareIvalidParameter(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/jokenpo?player1=pedra&player2=papel&test=teste", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rr := httptest.NewRecorder()
// 	emptyValidatorMiddleware := func(w http.ResponseWriter, r *http.Request) {
// 		assert.Equal(t, w.Header().Get("Code"), http.StatusOK)
// 	}
// 	ValidateMiddleware(http.HandlerFunc(emptyValidatorMiddleware)).ServeHTTP(rr, req)
// }
