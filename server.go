package main

import (
	"log"
	"os"

	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber"
	"github.com/tutunci/sponsorship/controllers"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {

	connectionString := os.Getenv("MONGODB_CONNECTION_STRING")
	if len(connectionString) == 0 {
		connectionString = "mongodb+srv://supervisor:AvumZkqOMUs0otCi@bucksense-ox2sg.mongodb.net/test?retryWrites=true&w=majority"
	}

	err := mgm.SetDefaultConfig(nil, "sponsorship", options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	app := fiber.New()

	app.Get("/api/owned", controllers.GetAllOwneds)
	app.Get("/api/todos/:id", controllers.GetOwnedByID)
	app.Post("/api/todos", controllers.CreateOwned)
	app.Patch("/api/todos/:id", controllers.ToggleOwnedStatus)
	//app.Delete("/api/todos/:id", controllers.DeleteTodo)

	app.Listen(3000)
}
