package protocol

import "github.com/gofiber/fiber/v2"

func (s *Server) NewMarketHandler(r *fiber.App) {
	admin := r.Group("api/market")

	admin.Get("/:id", s.getMarketDetail)
}

func (s *Server) getMarketDetail(c *fiber.Ctx) error {
	marketID, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	result, err := s.Dependencies.MarketSvc.GetByID(c.Context(), int64(marketID))
	if err != nil {
		return err
	}

	return c.JSON(result)
}
