package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Router struct {
	UserController *controllers.UserController
	TodoController *controllers.TodoController
	AuthController *controllers.AuthController
	Fiber          *fiber.App
}

func NewRouter(userController *controllers.UserController, todoController *controllers.TodoController, authController *controllers.AuthController) *Router {
	return &Router{
		UserController: userController,
		TodoController: todoController,
		AuthController: authController,
		Fiber:          fiber.New(),
	}
}

func (r *Router) SetupRoutes() {
	r.Fiber.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowCredentials: true,
	}))
	r.Fiber.Get("/:id", r.getUser)
}

func (r *Router) Start() {
	r.Fiber.Listen(":8080")
}

func (r *Router) getUser(c *fiber.Ctx) error {
}
