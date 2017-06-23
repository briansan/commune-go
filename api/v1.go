package api

import (
	"io"
	"os"

	"net/http"

	"github.com/labstack/echo"
	"github.com/mgutz/logxi/v1"

	"github.com/briansan/commune-go/errors"
	model "github.com/briansan/commune-go/model/schema"
	"github.com/briansan/commune-go/model/store"
)

const (
	audioType = "audio"
	imageType = "image"
	textType  = "text"
	videoType = "video"
)

var (
	contentMap = map[string][]string{
		audioType: []string{"audio/wave", "audio/aiff"},
		imageType: []string{"image/png", "image/jpeg"},
		textType:  []string{"text/plain"},
		videoType: []string{},
	}

	logger = log.New("api-v1")
)

func v1FilesPost(c echo.Context) error {
	title := c.FormValue("title")
	description := c.FormValue("description")

	//-----------
	// Read file
	//-----------

	// Get file
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	// Open file
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Detect content type
	data := make([]byte, 512)
	if _, err := src.Read(data); err != nil {
		return err
	}

	ct := http.DetectContentType(data)
	if ct == "application/octet-stream" {
		return c.JSON(http.StatusBadRequest, "no binaries allowed")
	}

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"title":        title,
		"description":  description,
		"content-type": ct,
	})
}

func v1UsersPost(c echo.Context) error {
	u := model.User{}
	c.Bind(&u)

	if err := u.Validate(); err != nil {
		return errors.Response(c, err)
	}

	// Get db connection
	db := store.NewMongoStore()
	defer db.Cleanup()

	// Try to add user
	if err := db.CreateUser(&u); err != nil {
		return errors.Response(c, err)
	}
	return c.JSON(http.StatusOK, u)
}

func initV1(api *echo.Group) error {
	v1 := api.Group("/v1")
	v1.POST("/users", v1UsersPost)
	v1.POST("/files", v1FilesPost)

	return nil
}
