package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const baseURL = "https://www.emsifa.com/api-wilayah-indonesia/api"

func fetchExternal(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}

func GetProvinsi(c *fiber.Ctx) error {
	data, _ := fetchExternal(fmt.Sprintf("%s/provinces.json", baseURL))
	return c.JSON(data)
}

func GetProvinsiDetail(c *fiber.Ctx) error {
	id := c.Params("prov_id")
	data, _ := fetchExternal(fmt.Sprintf("%s/province/%s.json", baseURL, id))
	return c.JSON(data)
}

func GetCitiesByProv(c *fiber.Ctx) error {
	id := c.Params("prov_id")
	data, _ := fetchExternal(fmt.Sprintf("%s/regencies/%s.json", baseURL, id))
	return c.JSON(data)
}

func GetCityDetail(c *fiber.Ctx) error {
	id := c.Params("city_id")
	data, _ := fetchExternal(fmt.Sprintf("%s/regency/%s.json", baseURL, id))
	return c.JSON(data)
}
