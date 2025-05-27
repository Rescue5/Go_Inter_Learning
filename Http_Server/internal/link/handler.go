package link

import (
	"HttpServer/configs"
	"HttpServer/extra"
	"HttpServer/pkg/req"
	"HttpServer/pkg/res"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Handler struct {
	HandlerDeps
}

type HandlerDeps struct {
	*configs.Config
	*Repository
}

func NewHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := Handler{deps}
	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("GET /{hash}", handler.GoTo())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
}

func (handler *Handler) Delete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		idString := request.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		err = handler.HandlerDeps.Repository.Delete(uint(id))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(writer, "Удалено", 200)
	}
}

func (handler *Handler) Create() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		body, err := req.HandleBody[CreateRequest](writer, request)
		if err != nil {
			return
		}

		link := NewLink(body.Url)
		createdLink, err := handler.HandlerDeps.Repository.Create(link)

		if err != nil {
			res.Json(writer, extra.WrapError("Ошибка создания ссылки:", err), 400)
			return
		}
		res.Json(writer, createdLink, 201)
	}
}

func (handler *Handler) GoTo() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		hash := request.PathValue("hash")
		link, err := handler.HandlerDeps.Repository.GetByHash(hash)
		if err != nil {
			res.Json(writer, extra.WrapError("Ошибка получения ссылки по хэшу:",
				err), http.StatusNotFound)
			return
		}
		http.Redirect(writer, request, link.Url, http.StatusTemporaryRedirect)
	}
}

func (handler *Handler) Update() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		body, err := req.HandleBody[UpdateRequest](writer, request)
		if err != nil {
			return
		}
		idString := request.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		}

		updatedLink, err := handler.HandlerDeps.Repository.Update(&Link{
			Model: gorm.Model{ID: uint(id)},
			Url:   body.Url,
			Hash:  body.Hash,
		})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		}
		res.Json(writer, updatedLink, 200)
	}
}
