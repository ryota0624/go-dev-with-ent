package user

import (
	"context"
	"github.com/go-faster/errors"
	"github.com/google/uuid"
	"github.com/ryota0624/go-dev-with-ent/ent"
)

type User interface {
	ID() uuid.UUID
	Age() int
	Name() string
}

type Factory interface {
	Create(ctx context.Context, groupId uuid.UUID, age int, name string) (User, error)
}

type userImpl struct {
	entUser ent.User
}

func (u *userImpl) Age() int {
	return u.entUser.Age
}

func (u *userImpl) Name() string {
	return u.entUser.Name
}

func (u userImpl) ID() uuid.UUID {
	return u.entUser.ID
}

type factoryImpl struct{}

var (
	ErrGroupIdShouldExists = errors.New("group id should exists")
)

var (
	EntClientContextKey = "ENTCLIENT"
)

func WithEntClient(ctx context.Context, entClient *ent.Client) context.Context {
	return context.WithValue(ctx, EntClientContextKey, entClient)
}

func GetEntClient(ctx context.Context) *ent.Client {
	return ctx.Value(EntClientContextKey).(*ent.Client)
}

func (factory *factoryImpl) Create(ctx context.Context, groupId uuid.UUID, age int, name string) (User, error) {
	entClient := GetEntClient(ctx)
	_, err := entClient.Group.Get(ctx, groupId)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return nil, ErrGroupIdShouldExists
		}
		return nil, err
	}
	id := uuid.New()

	return &userImpl{
		entUser: ent.User{
			ID:   id,
			Age:  age,
			Name: name,
		},
	}, nil
}
