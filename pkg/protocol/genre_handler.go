package protocol

import (
	"backend/pkg/identity/genre"
	"backend/pkg/infra/api/response"
	"backend/pkg/infra/api/routing"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) NewGenreHandler(r *routing.Router) {
	admin := r.Group("api/genre")

	admin.POST("/", s.createGenre)
	admin.GET("/", s.searchGenre)
	admin.GET("/:id", s.getGenreDetail)
	admin.PUT("/:id", s.updateGenre)
}

func (s *Server) createGenre(c *fiber.Ctx) error {
	var cmd genre.CreateGenre

	err := c.BodyParser(&cmd)
	if err != nil {
		return response.SendError(c, fiber.StatusBadRequest, err)
	}

	err = cmd.Validate()
	if err != nil {
		return response.SendError(c, fiber.StatusBadRequest, err)
	}

	err = s.Dependencies.GenreSvc.Create(c.Context(), &cmd)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, err)
	}

	return response.SuccessMessage(c, "Created successfully")
}

func (s *Server) searchGenre(c *fiber.Ctx) error {
	var (
		query       genre.SearchGenreQuery
		queryValues = c.Queries()
	)

	if len(queryValues) > 0 {
		err := c.QueryParser(&query)
		if err != nil {
			return response.SendError(c, fiber.StatusBadRequest, err)
		}
	}

	result, err := s.Dependencies.GenreSvc.Search(c.Context(), &query)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, err)
	}

	return response.Result(c, result)
}

func (s *Server) getGenreDetail(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "id is not valid")
	}

	result, err := s.Dependencies.GenreSvc.GetByID(c.Context(), id)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, err)
	}

	return response.Result(c, result)
}

func (s *Server) updateGenre(c *fiber.Ctx) error {
	var cmd genre.UpdateGenre

	idStr := c.Params("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "id is not valid")
	}

	err = c.BodyParser(&cmd)
	if err != nil {
		return response.SendError(c, fiber.StatusBadRequest, err)
	}

	cmd.ID = id
	err = cmd.Validate()
	if err != nil {
		return response.SendError(c, fiber.StatusBadRequest, err)
	}

	err = s.Dependencies.GenreSvc.Update(c.Context(), &cmd)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, err)
	}

	return response.SuccessMessage(c, "Updated successfully")
}
