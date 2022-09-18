package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	gopherinskillbox "skillbox-test"
	"strconv"
)

// addUser endpoint handler
func (h *Handler) addUser(ctx *gin.Context) {
	var input gopherinskillbox.User
	if err := ctx.BindJSON(&input); err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.services.CreateUser(input)
	displayError(ctx, err)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"name": user.Name,
		"age":  user.Age,
	})
}

// getUsers endpoint handler
func (h *Handler) getUsers(ctx *gin.Context) {
	allUsers, err := h.services.GetAllUsers()
	displayError(ctx, err)

	if allUsers == nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{"result": "no users yet"})
	} else {
		ctx.JSON(http.StatusOK, allUsers)
	}
}

// patchUser endpoint handler
func (h *Handler) patchUser(ctx *gin.Context) {
	var editId int

	editId, err := strconv.Atoi(ctx.Param("id"))
	displayError(ctx, err)

	var newUserData gopherinskillbox.UpdateUserData
	displayError(ctx, err)

	err = h.services.UpdateUser(editId, newUserData)
	displayError(ctx, err)
	ctx.JSON(http.StatusOK, map[string]interface{}{"result": "User data changed"})
}

// deleteUser endpoint handler
func (h *Handler) deleteUser(ctx *gin.Context) {
	var deleteId int

	deleteId, err := strconv.Atoi(ctx.Param("id"))
	displayError(ctx, err)

	err = h.services.DeleteUser(deleteId)
	displayError(ctx, err)
	ctx.JSON(http.StatusNoContent, map[string]interface{}{"result": "User deleted"})
}

// addFriend endpoint handler
func (h *Handler) addFriend(ctx *gin.Context) {
	var friendData gopherinskillbox.UserFriendData
	var ownerId int

	ownerId, err := strconv.Atoi(ctx.Param("id"))
	displayError(ctx, err)

	err = ctx.BindJSON(&friendData)
	displayError(ctx, err)

	err = h.services.AddFriend(ownerId, friendData)
	displayError(ctx, err)
	ctx.JSON(http.StatusOK, map[string]interface{}{"status": "User was added as a friend"})
}

// displayError - return StatusBadRequest
func displayError(ctx *gin.Context, err error) {
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusBadRequest)
	}
}
