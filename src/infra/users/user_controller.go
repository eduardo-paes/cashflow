package users

import (
	"fmt"
	"github.com/eduardo-paes/cashflow/core/users"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	UseCase users.UserUseCases
}

// NewUserController returns contract implementation of UserService
func NewUserController(usecase users.UserUseCases) users.UserService {
	return &UserController{
		UseCase: usecase,
	}
}

// @Summary		Login
// @Description	Login with user credentials
// @Tags			auth
// @Accept			json
// @Produce		json
// @Param			request	body		users.AuthInput	true	"User data"
// @Success		200		{object}	users.AuthOutput
// @Failure		400		{string}	string
// @Failure		500		{string}	string
// @Router			/auth/login [post]
func (s *UserController) Login(c *gin.Context) {
	authRequest, err := users.FromJSONAuthInput(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authOutput, err := s.UseCase.Login(authRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, authOutput)
}

// @Summary		Create a new user
// @Description	Create a new user
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			request	body		users.UserInput	true	"User data"
// @Success		200		{object}	users.User
// @Failure		400		{string}	string
// @Failure		500		{string}	string
// @Router			/user [post]
func (s *UserController) Create(c *gin.Context) {
	userRequest, err := users.FromJSONCreateUser(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.UseCase.Create(userRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("User created created successfully. ID: %v", user.ID))
}

// @Summary		Delete an user by ID
// @Description	Delete an user by ID
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"User ID"
// @Success		200	{object}	users.User
// @Failure		400	{string}	string
// @Failure		500	{string}	string
// @Router			/user/{id} [delete]
func (s *UserController) Delete(c *gin.Context) {
	// Extract the "id" parameter from the URL path
	idStr := c.Param("id")

	// Parse the "id" string to an int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		// Handle the case where "id" is not a valid integer
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	_, err = s.UseCase.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// @Summary		Get one user
// @Description	Get one user
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"User ID"
// @Success		200	{array}		users.User
// @Failure		400	{string}	string
// @Failure		500	{string}	string
// @Router			/user [get]
func (s *UserController) GetOne(c *gin.Context) {
	// Extract the "id" parameter from the URL path
	idStr := c.Param("id")

	// Parse the "id" string to an int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		// Handle the case where "id" is not a valid integer
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	// If an id is provided, fetch a specific user
	user, err := s.UseCase.GetOne(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If no user is found, return a 404
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary		Update an user by ID
// @Description	Update an user by ID
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id		path		int				true	"User ID"
// @Param			request	body		users.UserInput	true	"User data"
// @Success		200		{object}	users.User
// @Failure		400		{string}	string
// @Failure		500		{string}	string
// @Router			/user/{id} [put]
func (s *UserController) Update(c *gin.Context) {
	// Extract the "id" parameter from the URL path
	idStr := c.Param("id")

	// Parse the "id" string to an int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		// Handle the case where "id" is not a valid integer
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	// Extract the request body
	var userRequest users.UserInput
	if err := c.BindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	_, err = s.UseCase.Update(id, &userRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "User updated successfully")
}
