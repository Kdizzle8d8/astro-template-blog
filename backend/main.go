package main

import (
	"backend/prisma/db"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type BlogRequest struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func main() {
	godotenv.Load()
	e := echo.New()
	client := db.NewClient()
	ctx := context.Background()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}

	e.POST("/blogs", func(c echo.Context) error {
		req := new(BlogRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		fmt.Println(req)
		createdBlog, err := client.Blog.CreateOne(
			db.Blog.Title.Set(req.Title),
			db.Blog.Author.Set(req.Author),
			db.Blog.Content.Set(req.Content),
			db.Blog.Description.Set(req.Description),
		).Exec(ctx)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, createdBlog)
	})

	e.GET("/blogs", func(c echo.Context) error {
		blogs, err := client.Blog.FindMany().Exec(ctx)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, blogs)
	})

	e.GET("blogs/:id", func(c echo.Context) error {
		id := c.Param("id")
		blog, err := client.Blog.FindUnique(
			db.Blog.ID.Equals(id),
		).Exec(ctx)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, blog)
	})

	e.Logger.Fatal(e.Start(":" + os.Getenv("SERVER_PORT")))
}
