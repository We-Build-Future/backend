package protocol

import (
	"backend/pkg/identity/user"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) NewUserHandler(r *fiber.App) {
	admin := r.Group("api/user")

	admin.Post("/", s.createUser)
	admin.Get("/", s.searchUser)
	admin.Get("/:id", s.getUserDetail)
	admin.Put("/:id", s.updateUser)
	admin.Put("/:id/status", s.updateStatus)
	admin.Put("/:id/password", s.updatePassword)
}

func (s *Server) createUser(c *fiber.Ctx) error {
	var cmd user.CreateUserCommand

	err := c.BodyParser(&cmd)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = cmd.Validate()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
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
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	}

	result, err := s.Dependencies.UserSvc.Search(c.Context(), &query)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (s *Server) getUserDetail(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "id is not valid")
	}

	result, err := s.Dependencies.UserSvc.GetByID(c.Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (s *Server) updateUser(c *fiber.Ctx) error {
	var cmd user.UpdateUserCommand

	idStr := c.Params("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "id is not valid")
	}

	err = c.BodyParser(&cmd)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	cmd.ID = id
	err = cmd.Validate()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = s.Dependencies.UserSvc.Update(c.Context(), &cmd)
	if err != nil {
		return err
	}

	return c.SendString("Updated successfully")
}

func (s *Server) updateStatus(c *fiber.Ctx) error {
	var cmd user.UpdateStatusCommand

	idStr := c.Params("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "id is not valid")
	}

	err = c.BodyParser(&cmd)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	cmd.ID = id

	err = cmd.Validate()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = s.Dependencies.UserSvc.UpdateStatus(c.Context(), &cmd)
	if err != nil {
		return err
	}

	return c.SendString("Updated successfully")
}

func (s *Server) updatePassword(c *fiber.Ctx) error {
	var cmd user.UpdatePasswordCommand

	idStr := c.Params("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "id is not valid")
	}

	err = c.BodyParser(&cmd)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	cmd.ID = id

	err = cmd.Validate()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = s.Dependencies.UserSvc.UpdatePassword(c.Context(), &cmd)
	if err != nil {
		return err
	}

	return c.SendString("Updated successfully")
}
