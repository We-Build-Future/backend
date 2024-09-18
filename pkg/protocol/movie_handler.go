package protocol

import (
	"backend/pkg/identity/movie"
	"backend/pkg/infra/api/response"
	"backend/pkg/infra/api/routing"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) NewMovieHandler(r *routing.Router) {
	admin := r.Group("api/movie")

	admin.POST("/", s.createMovie)
	admin.GET("/", s.searchMovie)
	admin.GET("/:id", s.getMovieDetail)
	admin.PUT("/:id", s.updateMovie)
}

func (s *Server) createMovie(c *fiber.Ctx) error {
	var cmd movie.CreateMovie

	err := c.BodyParser(&cmd)
	if err != nil {
		return response.SendError(c, fiber.StatusBadRequest, err)
	}

	err = cmd.Validate()
	if err != nil {
		return response.SendError(c, fiber.StatusBadRequest, err)
	}

	err = s.Dependencies.MovieSvc.Create(c.Context(), &cmd)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, err)
	}

	return response.SuccessMessage(c, "Movie created successfully")
}

func (s *Server) searchMovie(c *fiber.Ctx) error {
	var (
		query       movie.SearchMovieQuery
		queryValues = c.Queries()
	)

	if len(queryValues) > 0 {
		err := c.QueryParser(&query)
		if err != nil {
			return response.SendError(c, fiber.StatusBadRequest, err)
		}
	}

	result, err := s.Dependencies.MovieSvc.Search(c.Context(), &query)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, err)
	}

	return response.Result(c, result)
}

func (s *Server) getMovieDetail(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "id is not valid")
	}

	result, err := s.Dependencies.MovieSvc.GetByID(c.Context(), id)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, err)
	}

	return response.Result(c, result)
}

func (s *Server) updateMovie(c *fiber.Ctx) error {
	var cmd movie.UpdateMovie

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

	err = s.Dependencies.MovieSvc.Update(c.Context(), &cmd)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, err)
	}

	return response.SuccessMessage(c, "Updated successfully")
}
