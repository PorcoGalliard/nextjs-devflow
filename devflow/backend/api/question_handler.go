package api

import (
	"errors"

	"github.com/fullstack/dev-overflow/db"
	"github.com/fullstack/dev-overflow/types"
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

func (h *QuestionHandler) HandleAskQuestion (ctx *fiber.Ctx) error {
	var params types.AskQuestionParams
	if err := ctx.BodyParser(&params); err != nil {
		return ErrBadRequest()
	}

	if errors := params.Validate(); len(errors) > 0 {
		return ctx.JSON(errors)
	}

	question := &types.Question{
		Title: params.Title,
		Description: params.Description,
		UserID: params.UserID,
		Tags: params.Tags,
		CreatedAt: params.CreatedAt,
	}

	insertedQuestion, err := h.questionStore.AskQuestion(ctx.Context(), question)
	if err != nil {
		return ErrBadRequest()
	}
	return ctx.JSON(insertedQuestion)

}