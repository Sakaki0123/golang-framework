package controllers

import (
	"fmt"
	"framework/framework"
	"io/fs"
	"io/ioutil"
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
	listID := ctx.GetParam(":list_id", "")
	pictureID := ctx.GetParam(":picture_id", "")
	type OUTPUT struct {
		ListID    string `json:"list_id"`
		PictureID string `json:"picture_id"`
	}
	output := &OUTPUT{
		ListID:    listID,
		PictureID: pictureID,
	}
	ctx.Json(output)
}

func PostsController(ctx *framework.MyContext) {

	name := ctx.FormKey("name", "defaultName")
	age := ctx.FormKey("age", "20")
	fileInfo, err := ctx.FormFile("file")

	if err != nil {
		ctx.WriteHeader(500)
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s_%s_%s", name, age, fileInfo.Filename), fileInfo.Data, fs.ModePerm)

	if err != nil {
		ctx.WriteHeader(500)
	}
	ctx.WriteString("success")
}

func PostsPageController(ctx *framework.MyContext) {
	ctx.WriteString(`<!DOCTYPE html>
	<html>
		<head>
			<title>form</title>
		</head>
		<body>
			<div>
				<form action="/posts" method="post" enctype="multipart/form-data">
					<div><label>name</label>: <input name="name"/></div>
					<div><label>age</label>: 
					<select name="age">
						<option value="1">1</option>
						<option value="2">2</option>
					</select></div>
					<button type="submit">submit</button>
					<input name="file" type="file"/>
				</form>
			</div>
		</body>
	</html>`)
}
