package routes

import (
	"evermos/controllers"
	"evermos/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	auth := api.Group("/auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)

	user := api.Group("/user", middleware.JWTMiddleware)
	user.Get("/", controllers.GetUser)
	user.Put("/", controllers.UpdateUser)
	user.Get("/alamat", controllers.GetAlamat)
	user.Post("/alamat", controllers.CreateAlamat)
	user.Put("/alamat/:id", controllers.UpdateAlamat)
	user.Delete("/alamat/:id", controllers.DeleteAlamat)

	toko := api.Group("/toko")
	toko.Get("/my", middleware.JWTMiddleware, controllers.GetMyToko)
	toko.Get("/", controllers.GetAllToko)
	toko.Get("/:id_toko", controllers.GetTokoByID)
	toko.Put("/:id_toko", middleware.JWTMiddleware, controllers.UpdateToko)

	category := api.Group("/category")
	category.Get("/", controllers.GetAllKategori)
	category.Get("/:id", controllers.GetKategoriByID)
	categoryAdmin := category.Group("/", middleware.JWTMiddleware, middleware.AdminMiddleware)
	categoryAdmin.Post("/", controllers.CreateKategori)
	categoryAdmin.Put("/:id", controllers.UpdateKategori)
	categoryAdmin.Delete("/:id", controllers.DeleteKategori)

	product := api.Group("/product")
	product.Get("/", controllers.GetAllProduk)
	product.Get("/:id", controllers.GetProdukByID)
	product.Post("/", middleware.JWTMiddleware, controllers.CreateProduk)
	product.Put("/:id", middleware.JWTMiddleware, controllers.UpdateProduk)
	product.Delete("/:id", middleware.JWTMiddleware, controllers.DeleteProduk)

	trx := api.Group("/trx", middleware.JWTMiddleware)
	trx.Get("/", controllers.GetAllTrx)
	trx.Get("/:id", controllers.GetTrxByID)
	trx.Post("/", controllers.CreateTrx)

	provcity := api.Group("/provcity")
	provcity.Get("/listprovincies", controllers.GetProvinsi)
	provcity.Get("/detailprovince/:prov_id", controllers.GetProvinsiDetail)
	provcity.Get("/listcities/:prov_id", controllers.GetCitiesByProv)
	provcity.Get("/detailcity/:city_id", controllers.GetCityDetail)
}
