package routes

import (
	"net/http"
	"socialmedia-api/src/controllers"
)

var usersRoutes = []Route{
	{
		URI:      "/users",
		Method:   http.MethodPost,
		Function: controllers.CreateUser,
		Auth:     false,
	},
	{
		URI:      "/users",
		Method:   http.MethodGet,
		Function: controllers.GetUsers,
		Auth:     true,
	},
	{
		URI:      "/users/{userID}",
		Method:   http.MethodGet,
		Function: controllers.GetUser,
		Auth:     true,
	},
	{
		URI:      "/users/{userID}",
		Method:   http.MethodPut,
		Function: controllers.UpdateUser,
		Auth:     true,
	},
	{
		URI:      "/users/{userID}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteUser,
		Auth:     true,
	},
	{
		URI:      "/users/{userID}/follow",
		Method:   http.MethodPost,
		Function: controllers.ToFollowUser,
		Auth:     true,
	},
	{
		URI:      "/users/{userID}/stop-following",
		Method:   http.MethodDelete,
		Function: controllers.StopFollowingUser,
		Auth:     true,
	},
	{
		URI:      "/users/{userID}/followers",
		Method:   http.MethodGet,
		Function: controllers.GetFollowers,
		Auth:     true,
	},
	{
		URI:      "/users/{userID}/following",
		Method:   http.MethodGet,
		Function: controllers.GetFollowing,
		Auth:     true,
	},
}
