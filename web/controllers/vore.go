package controllers

import (
	"elsenova/query"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VoreController struct{}

func (v *VoreController) Mount(baseGroup gin.IRouter) {
	us := baseGroup.Group("vore")
	us.GET("/", v.All)
	us.GET("/:id", v.GetOne)
	us.GET("/stats", v.Stats)
}

func (v *VoreController) All(c ctx) {
	all, err := query.Vore.Find()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, all)
}

func (v *VoreController) GetOne(c ctx) {
	Vore := query.Vore
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	entry, err := Vore.Where(Vore.ID.Eq(uint(id))).First()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, entry)
}

func (v *VoreController) Stats(c ctx) {
	var results []struct {
		UserID string
		Total  int
	}

	Vore := query.Vore
	err := Vore.
		Select(Vore.UserID, Vore.UserID.Count().As("total")).
		Group(Vore.UserID).
		Order(Vore.UserID.Count().Desc()).
		Scan(&results)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, results)
}
