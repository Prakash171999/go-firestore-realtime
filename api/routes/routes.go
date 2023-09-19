package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		// v1.GET("todo", Controllers.GetTodo)
		// v1.GET("todo/:id", Controllers.GetSingleTodo)
		// v1.POST("todo/add", Controllers.CreateATodo)
		// v1.PUT("todo/update/:id", Controllers.UpdateTodo)
		// v1.DELETE("todo/delete/:id", Controllers.DeleteATodo)
		// v1.DELETE("todo/delete-field/:id", Controllers.DeleteAField)
	}

	return r
}
