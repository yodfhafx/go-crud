package controllers

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type pagingResult struct {
	Page      int   `json:"page"`
	Limit     int   `json:"limit"`
	PrevPage  int   `json:"prevPage"`
	NextPage  int   `json:"nextPage"`
	Count     int64 `json:"count"`
	TotalPage int   `json:"totalPage"`
}

func pagingResource(ctx *gin.Context, query *gorm.DB, records interface{}) *pagingResult {
	// Get limit, page ?limit=10&page=2
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "12"))

	// count records
	var count int64
	query.Model(records).Count(&count)

	// Find Records
	// limit, offset
	// limit => 10
	// page => 1, 1 - 10, offset(skip) => 0
	// page => 2, 11 - 20, offset(skip) => 10
	// page => 3, 21 - 30, offset(skip) => 20
	offset := (page - 1) * limit
	query.Limit(limit).Offset(offset).Find(records)

	// total page
	totalPage := int(math.Ceil(float64(count) / float64(limit)))

	// Find nextPage
	var nextPage int
	if page == totalPage {
		nextPage = totalPage // totalPage == final page
	} else {
		nextPage = page + 1
	}

	// create pagingResult
	return &pagingResult{
		Page:      page,
		Limit:     limit,
		Count:     count,
		PrevPage:  page - 1,
		NextPage:  nextPage,
		TotalPage: totalPage,
	}
}
