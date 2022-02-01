package api

import "github.com/gofiber/fiber/v2"

func buyItemHandler(ctx *fiber.Ctx) error {
	summary := BuySummary{
		Summary:     "compra al 01/2021",
		TotalAmount: 2500,
	}
	return ctx.JSON(summary)
}

type BuySummary struct {
	Summary     string `json:summary`
	TotalAmount int    `json:total_amount`
}
