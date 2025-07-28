package routes

import (
	"net/http"
	"socialmedia-api/src/controllers"
)

var postsRoutes = []Route{
	{
		URI:      "/posts",
		Method:   http.MethodPost,
		Function: controllers.CreatePost,
		Auth:     true,
	},
	{
		URI:      "/posts",
		Method:   http.MethodGet,
		Function: controllers.Getposts,
		Auth:     true,
	},
	{
		URI:      "/posts/{postID}",
		Method:   http.MethodGet,
		Function: controllers.GetPost,
		Auth:     true,
	},
	{
		URI:      "/posts/{postID}",
		Method:   http.MethodPut,
		Function: controllers.UpdatePost,
		Auth:     true,
	},
	{
		URI:      "/posts/{postID}",
		Method:   http.MethodDelete,
		Function: controllers.DeletePost,
		Auth:     true,
	},
	{
		URI:      "/users/{userID}/posts",
		Method:   http.MethodGet,
		Function: controllers.GetPostsByUserId,
		Auth:     true,
	},
	{
		URI:      "/posts/{postID}/curtir",
		Method:   http.MethodPost,
		Function: controllers.LikePost,
		Auth:     true,
	},
	{
		URI:      "/posts/{postID}/descurtir",
		Method:   http.MethodPost,
		Function: controllers.UnlikePost,
		Auth:     true,
	},
}
