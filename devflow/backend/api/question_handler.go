package api

import (
	"errors"

	"github.com/fullstack/dev-overflow/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuestionHandler struct {
	questionStore db.QuestionStore
}

func NewQuestionHandler(questionStore db.QuestionStore) *QuestionHandler {
	return &QuestionHandler{
		questionStore: questionStore,
	}
}

func (h *QuestionHandler) HandleGetQuestionByID(ctx *fiber.Ctx) error {
	var (
		id = ctx.Params("id")
	)
	
	question, err := h.questionStore.GetQuestionByID(ctx.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ctx.JSON(map[string]string{"error": "Not Found"})
		}
		return err
	}

	return ctx.JSON(question)
}

