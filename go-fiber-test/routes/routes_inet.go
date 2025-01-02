package routes

import (
	c "go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v3 := api.Group("/v3")

	v1.Get("/usersprofile", c.GetUsersProfiles)
	v1.Get("/usersprofile/generations", c.GetGenerations)
	v1.Get("usersprofile/search", c.GetUserProfile)

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"testgo": "23012023",
		},
	}))

	v1.Get("/fact/:number", c.Factorial)
	v1.Post("/register", c.Register)
	v3.Get("/champ", c.NickName)

	//CRUD dogs
	dog := v1.Group("/dog")
	dog.Get("", c.GetDogs)
	dog.Get("/filter", c.GetDog)
	dog.Get("/json", c.GetDogsJson)
	dog.Post("/", c.AddDog)
	dog.Put("/:id", c.UpdateDog)
	dog.Delete("/:id", c.RemoveDog)
	dog.Get("/removedogs", c.GetRemoveDogs)
	dog.Get("dogsinrange", c.GetDogsInRange)
	dog.Get("/dogcount", c.GetDogCount)

	//CRUD company
	company := v1.Group("/company")
	company.Get("", c.GetCompanies)
	company.Get("/filter", c.GetCompanyId)
	company.Post("/", c.AddCompany)
	company.Put("/:id", c.UpdateCompany)
	company.Delete("/:id", c.RemoveCompany)

	//CRUD usersprofile
	usersprofile := v1.Group("/usersprofile")
	usersprofile.Post("/", c.AddUsersProfile)
	usersprofile.Put("/:id", c.UpdateUsersProfile)
	usersprofile.Delete("/:id", c.RemoveUsersProfile)

	// v1.Get("/", controllers.HelloTest)
	// v1.Post("/", controllers.BodyParser)
	// v1.Get("/user/:name", controllers.ParamsTest)
	// v1.Post("/inet", controllers.QueryTest)
	// v1.Post("/valid", controllers.ValidTest)
}
