package controllers

import (
	//"github.com/Kamva/mgm/v2/operator"
	//"github.com/tutunci/sponsorship"
	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber"
	"github.com/tutunci/sponsorship/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllOwneds(ctx *fiber.Ctx) {
	//TODO .implement this
	collection := mgm.Coll(&models.Owned{})
	owneds := []models.Owned{}

	err := collection.SimpleFind(&owneds, bson.D{})
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":     true,
		"owneds": owneds,
	})
}

func GetOwnedByID(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	owned := &models.Owned{}
	collection := mgm.Coll(owned)

	err := collection.FindByID(id, owned)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Owned not found",
		})
		return
	}
	ctx.JSON(fiber.Map{
		"ok":    true,
		"owned": owned,
	})
}

func CreateOwned(ctx *fiber.Ctx) {
	params := new(struct {
		Name        string
		Description string
		Title       string
		Status      string
	})

	ctx.BodyParser(&params)

	if len(params.Name) == 0 || len(params.Description) == 0 {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Name or description not defined",
		})
		return
	}

	owned := models.CreateOwned(params.Name, params.Description, params.Title, params.Status)
	err := mgm.Coll(owned).Create(owned)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":    true,
		"owned": owned,
	})
}

func ToggleOwnedStatus(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	owned := &models.Owned{}
	collection := mgm.Coll(owned)

	err := collection.FindByID(id, owned)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Todo not found.",
		})
		return
	}

	// owned.Status = !owned.Status
	if owned.Status == "active" {
		owned.Status = "paused"
	} else if owned.Status == "paused" {
		owned.Status = "paused"
	}

	err = collection.Update(owned)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": owned,
	})

}
