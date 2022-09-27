package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"skillbox-test/pkg"
	"strconv"
)

// addCity endpoint handler
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

// getCities endpoint handler
func (h *Handler) getCities(ctx *gin.Context) {
	allCities, err := h.services.GetAllCities()
	displayError(ctx, err)

	if allCities == nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{"result": "no cities yet"})
	} else {
		ctx.JSON(http.StatusOK, allCities)
	}
}

// getCities endpoint handler
func (h *Handler) filterCities(ctx *gin.Context) {

	parameter := ctx.Param("param")
	switch parameter {
	case "region":
		value, ok := ctx.GetQuery("value")
		if !ok {
			displayError(ctx, errors.New("no value provided"))
			return
		}
		filteredCities, err := h.services.FilterCitiesByRegion(value)
		if err != nil {
			displayError(ctx, err)
			return
		}
		if filteredCities != nil {
			ctx.JSON(http.StatusOK, filteredCities)

		} else {
			ctx.JSON(http.StatusOK, map[string]interface{}{"result": "no related cities yet"})
		}
	case "district":
		value, ok := ctx.GetQuery("value")
		if !ok {
			displayError(ctx, errors.New("no value provided"))
			return
		}
		filteredCities, err := h.services.FilterCitiesByDistrict(value)
		if err != nil {
			displayError(ctx, err)
			return
		}
		if filteredCities != nil {
			ctx.JSON(http.StatusOK, filteredCities)

		} else {
			ctx.JSON(http.StatusOK, map[string]interface{}{"result": "no related cities yet"})
		}
	case "population":
		min, ok := ctx.GetQuery("min")
		if !ok {
			displayError(ctx, errors.New("no value provided"))
			return
		}
		minVal, err := strconv.Atoi(min)
		if err != nil {
			displayError(ctx, err)
			return
		}
		max, ok := ctx.GetQuery("max")
		if !ok {
			displayError(ctx, errors.New("no value provided"))
			return
		}
		maxVal, err := strconv.Atoi(max)
		if err != nil {
			displayError(ctx, err)
			return
		}
		filteredCities, err := h.services.FilterCitiesByPopulation(minVal, maxVal)
		if err != nil {
			displayError(ctx, err)
			return
		}
		if filteredCities != nil {
			ctx.JSON(http.StatusOK, filteredCities)

		} else {
			ctx.JSON(http.StatusOK, map[string]interface{}{"result": "no related cities yet"})
		}

	case "foundation":
		min, ok := ctx.GetQuery("min")
		if !ok {
			displayError(ctx, errors.New("no value provided"))
			return
		}
		minVal, err := strconv.Atoi(min)
		if err != nil {
			displayError(ctx, err)
			return
		}
		max, ok := ctx.GetQuery("max")
		if !ok {
			displayError(ctx, errors.New("no value provided"))
			return
		}
		maxVal, err := strconv.Atoi(max)
		if err != nil {
			displayError(ctx, err)
			return
		}
		filteredCities, err := h.services.FilterCitiesByFoundation(minVal, maxVal)
		if err != nil {
			displayError(ctx, err)
			return
		}
		if filteredCities != nil {
			ctx.JSON(http.StatusOK, filteredCities)

		} else {
			ctx.JSON(http.StatusOK, map[string]interface{}{"result": "no related cities yet"})
		}

	default:
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"result": "Bad request"})

	}
}

// patchCity endpoint handler
func (h *Handler) patchCity(ctx *gin.Context) {
	var editId int

	editId, err := strconv.Atoi(ctx.Param("id"))
	displayError(ctx, err)

	var newData pkg.CityPopulation
	err = ctx.BindJSON(&newData)
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
