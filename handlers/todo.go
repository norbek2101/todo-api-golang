package handlers

import (
	"github.com/gin-gonic/gin"
	"todo-api-golang/db"
	"todo-api-golang/models"
	"strconv"
)

// @Summary Get all todos
// @Description Get a list of all todos
// @Tags todos
// @Success 200 {object} gin.H "Success response"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /todos [get]
func GetTodos(c *gin.Context) {
    var todos []models.Todo
    database.DB.Find(&todos)

    c.JSON(200, gin.H{"data": todos})
}


// CreateTodo creates a new todo
func CreateTodo(c *gin.Context) {
	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{Title: input.Title, Completed: input.Completed}
	database.DB.Create(&todo)

	c.JSON(201, gin.H{"data": todo})
}

// GetTodoByID returns a todo by ID
func GetTodoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid todo ID"})
		return
	}

	var todo models.Todo
	result := database.DB.First(&todo, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(200, gin.H{"data": todo})
}

// UpdateTodo updates a todo by ID
func UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid todo ID"})
		return
	}

	var existingTodo models.Todo
	result := database.DB.First(&existingTodo, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}

	var updatedTodo models.Todo
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	existingTodo.Title = updatedTodo.Title
	existingTodo.Completed = updatedTodo.Completed

	database.DB.Save(&existingTodo)

	c.JSON(200, gin.H{"data": existingTodo})
}

// DeleteTodo deletes a todo by ID
func DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid todo ID"})
		return
	}

	var todo models.Todo
	result := database.DB.First(&todo, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}

	database.DB.Delete(&todo)

	c.JSON(200, gin.H{"data": "Todo deleted successfully"})
}
