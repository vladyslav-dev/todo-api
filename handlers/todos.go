package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vladyslav-dev/todo-api/database"
	"github.com/vladyslav-dev/todo-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllTodos(c *fiber.Ctx) error {
	collection := database.GetCollection("todos")

	cursor, err := collection.Find(c.Context(), bson.M{})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	todos := make([]models.Todo, 0)

	if err = cursor.All(c.Context(), &todos); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(todos)
}

func GetOneTodo(c *fiber.Ctx) error {
	todoId := c.Params("id")

	dbId, err := primitive.ObjectIDFromHex(todoId)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	collection := database.GetCollection("todos")

	var todo models.Todo

	err = collection.FindOne(c.Context(), bson.M{"_id": dbId}).Decode(&todo)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(todo)
}

type CreateTodoDTO struct {
	Title       string `json:"title" bson:"title"`
	Completed   bool   `json:"completed" bson:"completed"`
	Description string `json:"description" bson:"description"`
	Date        string `json:"date" bson:"date"`
}

func CreateTodo(c *fiber.Ctx) error {
	newTodo := &CreateTodoDTO{}

	if err := c.BodyParser(newTodo); err != nil {
		return c.Status(400).JSON(fiber.Map{"erorr": err.Error()})
	}

	collection := database.GetCollection("todos")
	res, err := collection.InsertOne(c.Context(), newTodo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"inserted_id": res.InsertedID})
}

func UpdateTodo(c *fiber.Ctx) error {
	return nil
}

func DeleteTodo(c *fiber.Ctx) error {
	return nil
}
