package util

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"math"
	"sipamit-be/internal/pkg/const"
	"strconv"
	"strings"
)

type CommonQuery struct {
	Q      string `query:"q"`
	Device string `query:"device"`
	Sort   int8   `query:"sort"`

	Page  int `query:"page"`
	Limit int `query:"limit"`
}

func NewCommonQuery(c echo.Context) *CommonQuery {
	q := strings.ToLower(strings.TrimSpace(c.QueryParam("q")))
	device := strings.ToLower(strings.TrimSpace(c.QueryParam("device")))
	page := strings.ToLower(strings.TrimSpace(c.QueryParam("page")))
	limit := strings.ToLower(strings.TrimSpace(c.QueryParam("limit")))
	sort := strings.ToLower(strings.TrimSpace(c.QueryParam("sort")))

	if !_const.ValidDevice(device) {
		device = ""
	}

	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum < 1 {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limitNum < 1 {
		limitNum = 10
	}

	var sortNum int8
	switch sort {
	case "asc":
		sortNum = 1
	case "desc":
		sortNum = -1
	default:
		sortNum = 1
	}

	return &CommonQuery{
		Q:      q,
		Device: device,
		Page:   pageNum,
		Limit:  limitNum,
		Sort:   sortNum,
	}
}

func NilCommonQuery() *CommonQuery {
	return &CommonQuery{
		Q:      "",
		Device: "",
		Page:   1,
		Limit:  math.MaxInt,
		Sort:   1,
	}
}

type PaginationResult struct {
	Result interface{} `json:"result"`
	Total  int64       `json:"total"`
	Page   int         `json:"page"`
	Pages  int         `json:"pages"`
	Limit  int         `json:"limit"`
}

func BuildPaginationAndOrderOptionByField(sortParam bson.M, page, limit int) (*options.FindOptionsBuilder, error) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	skipped := (page - 1) * limit
	findOptions := options.Find()
	findOptions.SetSort(sortParam)
	findOptions.SetSkip(int64(skipped))
	findOptions.SetLimit(int64(limit))

	return findOptions, nil
}

func CalculateTotalPages(totalData, limit, page int) (int, bool) {
	return int(math.Ceil(float64(totalData) / float64(limit))),
		page > int(math.Ceil(float64(totalData)/float64(limit))) && int(math.Ceil(float64(totalData)/float64(limit))) != 0
}

func MakeResult(data interface{}, totalData int64, page, limit int) *PaginationResult {
	totalPages, pageOutOfRange := CalculateTotalPages(int(totalData), limit, page)
	if pageOutOfRange {
		return &PaginationResult{
			Result: nil,
			Total:  totalData,
			Pages:  totalPages,
			Page:   page,
			Limit:  limit,
		}
	}

	return &PaginationResult{
		Result: data,
		Total:  totalData,
		Pages:  totalPages,
		Page:   page,
		Limit:  limit,
	}
}
