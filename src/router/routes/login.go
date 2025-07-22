package routes

import (
	"net/http"
	"socialmedia-api/src/controllers"
)

var loginRoute = Route{
	URI:      "/login",
	Method:   http.MethodPost,
	Function: controllers.Login,
	Auth:     false,
}
