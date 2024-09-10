package protocol

import (
	"backend/pkg/identity/user"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) NewUserHandler(r *fiber.App) {
	admin := r.Group("api/user")

	admin.Post("/", s.createUser)
	admin.Get("/", s.searchUser)
}

func (s *Server) createUser(c *fiber.Ctx) error {
	var cmd user.CreateUserCommand

	err := c.BodyParser(&cmd)
	if err != nil {
		return err
	}

	err = cmd.Validate()
	if err != nil {
		return err
	}

	err = s.Dependencies.UserSvc.Create(c.Context(), &cmd)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) searchUser(c *fiber.Ctx) error {
	var (
		query       user.SearchUserQuery
		queryValues = c.Queries()
	)

	if len(queryValues) > 0 {
		err := c.QueryParser(&query)
		if err != nil {
			return err
		}
	}

	result, err := s.Dependencies.UserSvc.Search(c.Context(), &query)
	if err != nil {
		return err
	}

	return c.JSON(result)
}
