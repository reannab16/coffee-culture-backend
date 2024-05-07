package handler
 
import (
  "net/http"
	"coffee-culture.uk/pkg/serverless"
)
 
// @title Coffee Culture API
// @version 0
// @description This is the REST API for Coffee Culture.
// @host api.coffee-culture.uk
// @BasePath /v0
// @schemes https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func Handler(w http.ResponseWriter, r *http.Request) {
	router, cleanup := serverless.Initialize()
	defer cleanup()
	router.ServeHTTP(w, r)
}