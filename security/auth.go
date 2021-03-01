package security

import (
	"context"
	"encoding/json"
	"github.com/barandemirbas/go-jwt-server/config"
	"github.com/barandemirbas/go-jwt-server/database"
	"github.com/barandemirbas/go-jwt-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"regexp"
	"time"

	"github.com/gofiber/fiber/v2"

	jwt "github.com/form3tech-oss/jwt-go"
)

var user models.User
var alnum = regexp.MustCompile(`^\w+$`).MatchString

func Register(c *fiber.Ctx) error {
	collection, err := database.GetMongoDbCollection("user")
	if err != nil {
		c.Status(500)
		return err
	}

	json.Unmarshal([]byte(c.Body()), &user)

	if len(user.Name) < 3 {
		c.Status(500).Send([]byte("name must be +3 characters"))
		user.Password = ""
		return nil
	}

	if !alnum(user.Name) {
		c.Status(500).Send([]byte("username must be alphanumeric, you can use underscores"))
		user.Name = ""
		user.Password = ""
		return nil
	}

	if len(user.Password) < 8 {
		c.Status(500).Send([]byte("password must be +8 characters"))
		user.Name = ""
		return nil
	}

	filter := bson.M{"name": user.Name}

	var isUnique struct {
		Name string
	}

	collection.FindOne(context.Background(), filter).Decode(&isUnique)

	if isUnique.Name != "" {
		c.Status(500).Send([]byte("your username must be unique"))
		user.Name = ""
		user.Password = ""
		return nil
	}

	user.Password = Hmac256(user.Password, config.GetEnv("SECRET_KEY"))
	res, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		c.Status(500)
		return err
	}

	response, _ := json.Marshal(res)
	user.Name = ""
	user.Password = ""
	return c.Send(response)
}

func Login(c *fiber.Ctx) error {
	collection, err := database.GetMongoDbCollection("user")

	if err != nil {
		c.Status(500)
		return err
	}
	json.Unmarshal([]byte(c.Body()), &user)
	user.Password = Hmac256(user.Password, config.GetEnv("SECRET_KEY"))

	var results struct {
		Name     string
		Password string
	}

	filter := bson.M{"name": user.Name}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err = collection.FindOne(ctx, filter).Decode(&results)

	if err != nil {
		c.Status(500)
		return err
	}

	if results.Name == "" && results.Password == "" {
		c.SendStatus(404)
		return err
	}

	if user.Name != results.Name || user.Password != results.Password {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = results.Name
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.GetEnv("SECRET_KEY")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
