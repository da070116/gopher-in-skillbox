package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"skillbox-test/pkg"
	"strconv"
)

// addUser endpoint handler
func (h *Handler) addCity(ctx *gin.Context) {
	var input pkg.City
	if err := ctx.BindJSON(&input); err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	stringValues := pkg.CityWithIdToString(input)
	city, err := h.services.CreateCity(stringValues)
	displayError(ctx, err)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"name": city.Name,
	})
}

// getUsers endpoint handler
func (h *Handler) getCities(ctx *gin.Context) {
	allCities, err := h.services.GetAllCities()
	displayError(ctx, err)

	if allCities == nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{"result": "no cities yet"})
	} else {
		ctx.JSON(http.StatusOK, allCities)
	}
}

// patchUser endpoint handler
func (h *Handler) patchCity(ctx *gin.Context) {
	var editId int

	editId, err := strconv.Atoi(ctx.Param("id"))
	displayError(ctx, err)

	var newData pkg.CityPopulation
	displayError(ctx, err)

	err = h.services.UpdateCity(editId, newData)
	displayError(ctx, err)
	ctx.JSON(http.StatusOK, map[string]interface{}{"result": "City data changed"})
}

// deleteUser endpoint handler
func (h *Handler) deleteCity(ctx *gin.Context) {
	var deleteId int

	deleteId, err := strconv.Atoi(ctx.Param("id"))
	displayError(ctx, err)

	err = h.services.DeleteCity(deleteId)
	displayError(ctx, err)
	ctx.JSON(http.StatusNoContent, map[string]interface{}{"result": "City deleted"})
}

// displayError - return StatusBadRequest
func displayError(ctx *gin.Context, err error) {
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusBadRequest)
	}
}
