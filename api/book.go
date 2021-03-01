package api

import (
	"context"
	"encoding/json"
	"github.com/barandemirbas/go-jwt-server/database"
	"github.com/barandemirbas/go-jwt-server/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var book models.Book

func GetBook(c *fiber.Ctx) error {
	collection, err := database.GetMongoDbCollection("book")
	if err != nil {
		c.Status(500)
		return err
	}

	var filter bson.M = bson.M{}

	if c.Params("id") != "" {
		id := c.Params("id")
		objID, _ := primitive.ObjectIDFromHex(id)
		filter = bson.M{"_id": objID}
	}

	var results []bson.M
	cur, err := collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())

	if err != nil {
		c.Status(500)
		return err
	}

	cur.All(context.Background(), &results)

	if results == nil {
		c.SendStatus(404)
		return err
	}

	json, _ := json.Marshal(results)
	return c.Send(json)
}

func AddBook(c *fiber.Ctx) error {
	collection, err := database.GetMongoDbCollection("book")
	if err != nil {
		c.Status(500)
		return err
	}

	json.Unmarshal([]byte(c.Body()), &book)

	if len(book.Name) < 3 {
		c.Status(500).Send([]byte("name cannot be less than 3 characters"))
		book.Author = ""
		book.Rating = 0
		return nil
	}

	if len(book.Author) < 3 {
		c.Status(500).Send([]byte("author cannot be less than 3 characters"))
		book.Name = ""
		book.Rating = 0
		return nil
	}

	if book.Rating == 0 || book.Rating > 5 {
		c.Status(500).Send([]byte("rating must be between 1-5"))
		book.Name = ""
		book.Author = ""
		return nil
	}

	res, err := collection.InsertOne(context.Background(), book)
	if err != nil {
		c.Status(500)
		return err
	}

	response, _ := json.Marshal(res)
	book.Name = ""
	book.Author = ""
	book.Rating = 0
	return c.Send(response)
}

func UpdateBook(c *fiber.Ctx) error {
	collection, err := database.GetMongoDbCollection("book")

	if err != nil {
		c.Status(500)
		return err
	}

	json.Unmarshal([]byte(c.Body()), &book)

	update := bson.D{{"$set",
		bson.D{
			{"name", book.Name},
			{"author", book.Author},
			{"rating", book.Rating},
		},
	}}

	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	res, err := collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)

	if err != nil {
		c.Status(500)
		return err
	}

	response, _ := json.Marshal(res)
	return c.Send(response)
}

func DeleteBook(c *fiber.Ctx) error {
	collection, err := database.GetMongoDbCollection("book")

	if err != nil {
		c.Status(500)
		return err
	}

	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	res, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID})

	if err != nil {
		c.Status(500)
		return err
	}

	jsonResponse, _ := json.Marshal(res)
	return c.Send(jsonResponse)
}
