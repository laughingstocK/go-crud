package http

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/laughingstocK/go-crud/author"
	"github.com/laughingstocK/go-crud/models"
)

type ResponseError struct {
	Message string `json:"message"`
}

type AuthorHandler struct {
	AuthorUsecase author.Usecase
}

func NewAuthorHandler(e *echo.Echo, AuthorUsecase author.Usecase) {
	handler := &AuthorHandler{
		AuthorUsecase,
	}

	e.GET("/author/:id", handler.GetByID)
	e.POST("/author", handler.Create)
}

func (a *AuthorHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := a.AuthorUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func (a *AuthorHandler) Create(c echo.Context) error {
	var author models.Author
	err := c.Bind(&author)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	fmt.Printf("Author: %+v\n", author)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := a.AuthorUsecase.Create(ctx, &author)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

// func (a *ArticleHandler) Store(c echo.Context) error {
// 	var article models.Article
// 	err := c.Bind(&article)
// 	if err != nil {
// 		return c.JSON(http.StatusUnprocessableEntity, err.Error())
// 	}

// 	if ok, err := isRequestValid(&article); !ok {
// 		return c.JSON(http.StatusBadRequest, err.Error())
// 	}
// 	ctx := c.Request().Context()
// 	if ctx == nil {
// 		ctx = context.Background()
// 	}

// 	err = a.AUsecase.Store(ctx, &article)

// 	if err != nil {
// 		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
// 	}
// 	return c.JSON(http.StatusCreated, article)
// }

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
