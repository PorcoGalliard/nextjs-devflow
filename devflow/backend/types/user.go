package types

import (
	"fmt"
	"time"

	"net/mail"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost = 12
	minFirstNameLen = 2
	minLastNameLen = 2
	minPasswordLen = 7
)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ClerkID string `bson:"clerkID" json:"clerkID"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName string `bson:"lastName" json:"lastName"`
	Bio *string `bson:"bio" json:"bio"`
	Picture *string `bson:"picture" json:"picture"` 
	Email string `bson:"email" json:"email"`
	EncryptedPassword string `bson:"password" json:"-"`
	Location *string `bson:"location" json:"location"`
	PortfolioWebsite *string `bson:"portfolioWebsite" json:"portfolioWebsite"`
	IsAdmin bool `bson:"isAdmin" json:"isAdmin"`
	Reputation *int `bson:"reputation" json:"reputation"`
	Saved []Question `bson:"saved" json:"saved"`
	JoinedAt time.Time `bson:"joinedAt" json:"joinedAt"`
}

type CreateUserParam struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	ClerkID string `json:"clerkID"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserParam struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

func NewUserFromParams(params CreateUserParam) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName: params.FirstName,
		LastName: params.LastName,
		ClerkID: params.ClerkID,
		Email: params.Email,
		EncryptedPassword: string(encpw),
	}, nil
}

func isValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsValidPassword(encpw, pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encpw), []byte(pw)) == nil
}

func (params CreateUserParam) Validate() map[string]string {
	errors := map[string]string{}

	if len(params.FirstName) < minFirstNameLen {
		errors["FirstName"] = fmt.Sprintf("First Name must be at least %d characters", minFirstNameLen)
	}
	
	if len(params.LastName) < minLastNameLen {
		errors["LastName"] = fmt.Sprintf("Last Name must be at least %d characters", minLastNameLen)
	}

	if !isValid(params.Email) {
		errors["Email"] = fmt.Sprintf("Your email %s is not a valid email", params.Email)
	}

	if len(params.Password) < minPasswordLen {
		errors["Password"] = fmt.Sprintf("Password must be at least %d characters", minPasswordLen)
	}

	return errors
}