package controllers

import (
	"fmt"
	"framework/framework"
	"io/fs"
	"io/ioutil"
)

type PostsPageForm struct {
	Name string
}

func PostsPageController(ctx *framework.MyContext) {
	postsPageForm := &PostsPageForm{
		Name: "defaultName",
	}
	ctx.RenderHtml("./htmls/post_page.html", postsPageForm)
}

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

type UserPost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func UserPostsController(ctx *framework.MyContext) {
	userPost := &UserPost{}

	if err := ctx.BindJson(userPost); err != nil {
		ctx.WriteHeader(500)
		return
	}

	ctx.Json(userPost)
}

func JsonPTestController(ctx *framework.MyContext) {
	queryKey := ctx.QueryKey("callback", "cb")

	ctx.JsonP(queryKey, "bbb")

}
