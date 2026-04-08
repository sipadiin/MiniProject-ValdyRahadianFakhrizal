package utils

import "github.com/gofiber/fiber/v2"

type PaginationResponse struct {
	Data  interface{} `json:"data"`
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
	Total int64       `json:"total"`
}

func BuildPagination(data interface{}, page, limit int, total int64) PaginationResponse {
	return PaginationResponse{
		Data:  data,
		Page:  page,
		Limit: limit,
		Total: total,
	}
}

func GetPaginationParams(c *fiber.Ctx) (page int, limit int, offset int) {
	page = c.QueryInt("page", 1)
	limit = c.QueryInt("limit", 10)

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset = (page - 1) * limit
	return page, limit, offset
}
