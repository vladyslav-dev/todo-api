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
	todoId := c.Params("id") // todoId = 6460165c0ce6d4d8aaaaa56c

	dbId, err := primitive.ObjectIDFromHex(todoId) // dbId = ObjectID("6460165c0ce6d4d8aaaaa56c")

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

func GetAllCompletedTodos(c *fiber.Ctx) error {
	collection := database.GetCollection("todos")

	cursor, err := collection.Find(c.Context(), bson.M{"completed": true})

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	completedTodos := make([]models.Todo, 0)

	cursor.All(c.Context(), &completedTodos)

	return c.Status(200).JSON(completedTodos)
}

type CreateTodoDTO struct {
	Title       string `json:"title" bson:"title"`
	Completed   bool   `json:"completed" bson:"completed"`
	Description string `json:"description" bson:"description"`
	Date        string `json:"date" bson:"date"`
	Count       int64  `json:"count" bson:"count"`
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

type UpdateTodoDTO struct {
	Title       string `json:"title" bson:"title"`
	Completed   bool   `json:"completed" bson:"completed"`
	Description string `json:"description" bson:"description"`
	Date        string `json:"date" bson:"date"`
	Count       int64  `json:"count" bson:"count"`
}

type UpdatedTodoResDTO struct {
	UpdatedCount int64 `json:"updated_count" bson:"updated_count"`
}

func UpdateTodo(c *fiber.Ctx) error {
	paramId := c.Params("id")

	dbId, err := primitive.ObjectIDFromHex(paramId)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	uTodo := &UpdateTodoDTO{}

	if err = c.BodyParser(uTodo); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	collection := database.GetCollection("todos")

	update := bson.M{"$set": uTodo}

	res, err := collection.UpdateByID(c.Context(), dbId, update)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"updated_count": res.MatchedCount})

}

type DeletedTodoResDTO struct {
	DeletedCount int64 `json:"deleted_count" bson:"deleted_count"`
}

func DeleteTodo(c *fiber.Ctx) error {
	paramId := c.Params("id")

	dbId, err := primitive.ObjectIDFromHex(paramId)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}

	collection := database.GetCollection("todos")

	res, err := collection.DeleteOne(c.Context(), bson.M{"_id": dbId})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"deletedCount": res.DeletedCount})
}
