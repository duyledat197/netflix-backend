package domain

import (
	"context"

	"github.com/duyledat197/netfix-backend/common/errors"
	"github.com/duyledat197/netfix-backend/internal/models"
	"github.com/duyledat197/netfix-backend/internal/repository"
	"github.com/duyledat197/netfix-backend/utils/requestinfo"
	"github.com/duyledat197/netfix-backend/utils/token"
)

// AuthDomain ...
type AuthDomain interface {
	Login(ctx context.Context, email, password string) (*models.User, string, error)
	Register(ctx context.Context, user *models.User) error
}

type authDomain struct {
	userStore     repository.UserRepository
	authenticator token.Authenticator
}

// NewAuthDomain ...
func NewAuthDomain(userStore repository.UserRepository, authenticator token.Authenticator) AuthDomain {
	return &authDomain{
		userStore:     userStore,
		authenticator: authenticator,
	}
}

func (d *authDomain) Login(ctx context.Context, email, password string) (*models.User, string, error) {
	user, err := d.userStore.FindByEmail(email)
	if err != nil {
		return nil, "", err
	}
	if !user.IsCorrectPassword(password) {
		return nil, "", errors.ErrPasswordIsNotCorrect
	}

	tkn, err := d.authenticator.Generate(&requestinfo.Info{
		UserID:   user.ID.Hex(),
		UserName: user.FullName,
	})

	if err != nil {
		return nil, "", err
	}
	return user, tkn.Token, nil
}

func (d *authDomain) Register(ctx context.Context, user *models.User) error {
	_, err := d.userStore.FindByEmail(user.Email)
	if err != nil && err != errors.ErrEmailNotFound {
		return err
	}
	if err == nil {
		return errors.ErrEmailAlreadyExists
	}

	err = d.userStore.Create(user)
	if err != nil {
		return err
	}
	return nil
}
