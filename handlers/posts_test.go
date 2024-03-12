package handlers_test

import (
	"fmt"
	"github.com/che-ict/DEV-DT-Microblog/handlers"
	"github.com/che-ict/DEV-DT-Microblog/repositories"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestDeletePostHandler_Allowed(t *testing.T) {
	repositories.CreateUser("bert", "b", "Bert")
	repositories.CreateUser("ernie", "e", "Ernie")
	repositories.CreatePost("Dit is mijn tweet", "bert")
	post := repositories.GetPostsByUser("bert")[0]

	e := echo.New()
	req := httptest.NewRequest("DELETE", fmt.Sprintf("/post/%d", post.ID), nil)
	req.AddCookie(&http.Cookie{Name: "user", Value: "bert", Expires: time.Now().Add(time.Hour * 2)})
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/post/:id")
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprintf("%d", post.ID))

	if assert.NoError(t, handlers.DeletePostHandler(c)) {
		assert.Equal(t, http.StatusFound, rec.Code)
	}
}

func TestDeletePostHandler_NotAllowed(t *testing.T) {
	repositories.CreateUser("elmo", "b", "Elmo")
	repositories.CreateUser("grover", "e", "Grover")
	repositories.CreatePost("Elmo houdt van jou!", "elmo")
	post := repositories.GetPostsByUser("elmo")[0]

	e := echo.New()
	req := httptest.NewRequest("DELETE", fmt.Sprintf("/post/%d", post.ID), nil)
	req.AddCookie(&http.Cookie{Name: "user", Value: "grover", Expires: time.Now().Add(time.Hour * 2)})
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/post/:id")
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprintf("%d", post.ID))

	if assert.NoError(t, handlers.DeletePostHandler(c)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	}
}
