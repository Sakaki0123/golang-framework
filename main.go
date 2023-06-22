package main

import (
  "framework/controllers"
  "framework/framework"
)

func main() {
  engine := framework.NewEngine()
  router := engine.Router

  router.Get("/lists", controllers.ListsController)
  router.Get("/users", controllers.UsersController)
  router.Get("/students", controllers.StudentsController)
  engine.Run()
}
