package handlers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"todo-api-golang/db"
	"todo-api-golang/models"
)

// GetTodos godoc
// @Summary Get all todos
// @Description Retrieves a list of all todos in the database
// @Tags todos
// @Produce json
// @Success 200 {object} []models.Todo
// @Failure 500 {object} gin.H
func GetTodos(c *gin.Context) {
	var todos []models.Todo
	database.DB.Find(&todos)

	c.JSON(200, gin.H{"data": todos})
}

// CreateTodo godoc
// @Summary Create a new todo
// @Description Creates a new todo with the specified title and completed status
// @Tags todos
// @Accept json
// @Param todo body models.Todo true "Todo details"
// @Success 201 {object} models.Todo
// @Failure 400 {object} gin.H
// @Failure 422 {object} gin.H
// @Failure 500 {object} gin.H
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

// GetTodoByID godoc
// @Summary Get a todo by ID
// @Description Retrieves a single todo by its ID
// @Tags todos
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} models.Todo
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
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

// UpdateTodo godoc
// @Summary Update a todo by ID
// @Description Updates an existing todo with the specified ID
// @Tags todos
// @Accept json
// @Param id path int true "Todo ID"
// @Param todo body models.Todo true "Updated todo details"
// @Success 200 {object} models.Todo
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 422 {object} gin.H
// @Failure 500 {object} gin.H
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

// DeleteTodo godoc
// @Summary Delete a todo by ID
// @Description Deletes a single todo by its ID
// @Tags todos
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
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
