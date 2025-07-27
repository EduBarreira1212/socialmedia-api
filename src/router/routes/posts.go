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
		URI:      "/posts/{publicacaoId}",
		Method:   http.MethodGet,
		Function: controllers.GetPost,
		Auth:     true,
	},
	{
		URI:      "/posts/{publicacaoId}",
		Method:   http.MethodPut,
		Function: controllers.UpdatePost,
		Auth:     true,
	},
	{
		URI:      "/posts/{publicacaoId}",
		Method:   http.MethodDelete,
		Function: controllers.DeletePost,
		Auth:     true,
	},
	{
		URI:      "/usuarios/{usuarioId}/posts",
		Method:   http.MethodGet,
		Function: controllers.GetPostsByUserId,
		Auth:     true,
	},
	{
		URI:      "/posts/{publicacaoId}/curtir",
		Method:   http.MethodPost,
		Function: controllers.LikePost,
		Auth:     true,
	},
	{
		URI:      "/posts/{publicacaoId}/descurtir",
		Method:   http.MethodPost,
		Function: controllers.UnlikePost,
		Auth:     true,
	},
}
