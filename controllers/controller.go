package controllers

import (
	"framework/framework"
)

type StudentResponse struct {
	Name string `json:"name"`
}

func UsersController(ctx *framework.MyContext) {
	ctx.WriteString("users")
}
func ListsController(ctx *framework.MyContext) {
	ctx.WriteString("lists")
}
func ListItemController(ctx *framework.MyContext) {
	ctx.WriteString("lists_items")
}
func ListNameController(ctx *framework.MyContext) {
	ctx.WriteString("lists_names")
}
func StudentsController(ctx *framework.MyContext) {
	name := ctx.QueryKey("name", "")
	studentResponse := &StudentResponse{
		Name: name,
	}

	ctx.Json(studentResponse)
	ctx.WriteString("students")
}

func ListPicturesController(ctx *framework.MyContext) {
	ctx.WriteString("lists_pictures")
}

