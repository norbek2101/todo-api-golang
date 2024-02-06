package routes

import (
    "github.com/gin-gonic/gin"
    "todo-api-golang/handlers"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    v1 := router.Group("/v1")
    {
        v1.GET("/todos", handlers.GetTodos)
        v1.POST("/todos", handlers.CreateTodo)
        v1.GET("/todos/:id", handlers.GetTodoByID)
        v1.PUT("/todos/:id", handlers.UpdateTodo)
        v1.DELETE("/todos/:id", handlers.DeleteTodo)
    }

    return router
}
