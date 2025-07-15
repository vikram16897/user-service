package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/user-service/model"
	"github.com/user-service/service"
	dto "github.com/user-service/users"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(s service.UserService) *UserController {
	return &UserController{s}
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	if err := ctrl.userService.CreateUser(&user); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Failed to create user"})
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (ctrl *UserController) GetAllUsers(c *gin.Context) {
	users, err := ctrl.userService.GetUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
	}

	c.JSON(http.StatusOK, users)
}

func (ctrl *UserController) GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	users, err := ctrl.userService.GetUsersByID((id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
	}

	c.JSON(http.StatusOK, users)
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	if err := ctrl.userService.UpdateUser(&user); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Failed to create user"})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) GetUsersByEmail(c *gin.Context) {
	var req dto.GetUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	users, err := ctrl.userService.GetUsersByEmail((req.Email))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
	}

	c.JSON(http.StatusOK, users)
}
