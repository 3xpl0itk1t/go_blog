package handlers

import (
	"context"
	"go_blog/config"
	"go_blog/models"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func SignupUser(c *fiber.Ctx) error {
	var usr models.SignupModel
	err := c.BodyParser(&usr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid body",
		})
	}
	// Check if user with the same email already exists
	existingUser := bson.M{"email": usr.Email}
	count, err := user_collection.CountDocuments(context.Background(), existingUser)
	if err != nil {
		log.Fatal(err)
	}

	if count > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "user with this email already exists",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	newuser := bson.M{
		"first_name": usr.First_Name,
		"last_name":  usr.Last_Name,
		"email":      usr.Email,
		"password":   string(hashedPassword),
	}

	_, err = user_collection.InsertOne(context.Background(), newuser)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(fiber.Map{
		"message": "User Created Successfully",
	})
}

func Login(c *fiber.Ctx) error {
	var login models.LoginModel

	err := c.BodyParser(&login)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid body",
		})
	}
	storedPassword, userID := getPasswordAndID(login.Email)

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(login.Password))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid credentials",
		})
	}
	token := generateToken(userID.Hex())
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login Successful",
		"token":   token,
	})

}

func getPasswordAndID(email string) (string, primitive.ObjectID) {

	var result bson.M
	filter := bson.M{"email": email}
	err := user_collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	hashedPassword := result["password"].(string)
	userID := result["_id"].(primitive.ObjectID)
	return hashedPassword, userID
}

func generateToken(userID string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["_id"] = userID

	tokenstring, err := token.SignedString([]byte(config.SECRET))
	if err != nil {
		log.Fatal(err)
	}
	return tokenstring
}
