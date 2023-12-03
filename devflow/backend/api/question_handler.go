package api

import (
	"errors"
	"fmt"
	"time"

	"github.com/fullstack/dev-overflow/db"
	"github.com/fullstack/dev-overflow/types"
	"github.com/fullstack/dev-overflow/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuestionHandler struct {
	questionStore db.QuestionStore
	userStore db.UserStore
	tagStore db.TagStore
}

func NewQuestionHandler(questionStore db.QuestionStore, userStore db.UserStore, tagStore db.TagStore) *QuestionHandler {
	return &QuestionHandler{
		questionStore: questionStore,
		userStore: userStore,
		tagStore: tagStore,
	}
}

func (h *QuestionHandler) HandleGetQuestionByID(ctx *fiber.Ctx) error {
	var (
		id = ctx.Params("_id")
	)
	
	question, err := h.questionStore.GetQuestionByID(ctx.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ErrResourceNotFound(id)
		}
		return err
	}

	return ctx.JSON(question)
}

func (h *QuestionHandler) HandleAskQuestion(ctx *fiber.Ctx) error {
	var params types.AskQuestionParams
	// fmt.Println(string(ctx.Body()))
	// fmt.Println(params)
	if err := ctx.BodyParser(&params); err != nil {
		return ErrBadRequest()
	}

	if errors := params.Validate(); len(errors) > 0 {
		return ctx.JSON(errors)
	}

	for i, tag := range params.Tags {
		tag = utils.FormatTag(tag)
		params.Tags[i] = tag
	}

	user, err := h.userStore.GetUserByID(ctx.Context(), params.ClerkID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ErrResourceNotFound(params.ClerkID)
		}
	}

	tags := make([]types.Tag, len(params.Tags))
	for i, tagName := range params.Tags {

		tag, err := h.tagStore.GetTagByName(ctx.Context(), tagName)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				tag = &types.Tag{
					Name: tagName,
				}

				insertedTag, err := h.tagStore.CreateTag(ctx.Context(), tag)
				if err != nil {
					return ErrBadRequest()
				}

				tag = insertedTag
			} 
		}
		tags[i] = *tag
	}

	question := &types.Question{
		Title: params.Title,
		Description: params.Description,
		UserID: user.ID,
		Tags: tags,
		CreatedAt: time.Now().UTC(),
	}


	insertedQuestion, err := h.questionStore.AskQuestion(ctx.Context(), question)
	fmt.Println(insertedQuestion)
	if err != nil {
			return ErrBadRequest()
		}

	for _, tag := range tags {
		tagFromDB, err := h.tagStore.GetTagByID(ctx.Context(), tag.ID.Hex())
		if err != nil {
			return ErrBadRequest()
		}

		tagFromDB.Questions = append(tagFromDB.Questions, insertedQuestion.ID)
		tagFromDB.Followers = append(tagFromDB.Followers, user.ID)

		if err := h.tagStore.UpdateTag(ctx.Context(), db.Map{"_id":tag.ID}, &types.UpdateTagQuestionAndFollowers{Questions: tagFromDB.Questions, Followers: tagFromDB.Followers}); err != nil {
			return ErrBadRequest()
		}
	}

	return ctx.JSON(insertedQuestion)
}