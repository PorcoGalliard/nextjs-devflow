package api

import (
	"os"
	"time"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/fullstack/dev-overflow/db"
	"github.com/fullstack/dev-overflow/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleSignUp(c *fiber.Ctx) error {
	sessionToken := c.Get("Authorization")
	apiKey := os.Getenv("CLERK_SECRET_KEY")

	clerkClient, err := clerk.NewClient(apiKey)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	session, err := clerkClient.Sessions().Verify(sessionToken, "")
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}
	
	clerkUser, err := clerkClient.Users().Read(session.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	userOid := primitive.NewObjectID()

	user := &types.User{
		ID: userOid,
		ClerkID: clerkUser.ID,
		JoinedAt: time.Now(),
	}

	insertedUser, err := h.userStore.CreateUser(c.Context(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "User berhasil ditambahkan, ID nya adalah =>" + insertedUser.ClerkID})
}