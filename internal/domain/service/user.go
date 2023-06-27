package service

import (
	"context"
	"errors"
	"skillbox/internal/adapters/storage/postgres"
	"skillbox/internal/domain/model"

	"github.com/sirupsen/logrus"
)

//go:generate mockgen -source=user.go -destination=mocks/mock.go

//go:generate mockgen -source=user.go -destination=mocks/mock.go

type UserService interface {
	CreateUser(ctx context.Context, data *model.User) (uint64, error)
	MakeFriend(ctx context.Context, source, target uint64) (string, string, error)
	GetFriends(ctx context.Context, id uint64) ([]model.User, error)
	UpdateUserAge(ctx context.Context, id, age uint64) error
	DeleteUser(ctx context.Context, id uint64) error
}
type service struct {
	l    *logrus.Logger
	repo postgres.UserRepository
}

// New service.
func NewUserService(repo postgres.UserRepository, logger *logrus.Logger) UserService {
	return &service{
		l:    logger,
		repo: repo,
	}
}

// Create creating new user.
func (s *service) CreateUser(ctx context.Context, data *model.User) (uint64, error) {
	// var user model.User
	// age, err := strconv.Atoi(data.Age)
	// if err != nil {
	// 	s.l.Info(err)
	// 	//c.String(http.StatusBadRequest, "cant parse params", err.Error())
	// }
	// user.Age = uint64(age)
	// user.Name = data.Name
	id, err := s.repo.CreateUser(ctx, *data)
	if err != nil {
		s.l.Info(err)
	}
	return id, nil
}

// MakeFriend checks if there are users with the specified id, if there is, it adds the user with the source id to the list of friends of the target id.
func (s *service) MakeFriend(ctx context.Context, source, target uint64) (string, string, error) {
	targetUser, err := s.repo.GetUser(ctx, target)
	if err != nil {
		return "", "", errors.New("cant find target user")
	}
	sourceUser, err := s.repo.GetUser(ctx, source)
	if err != nil {
		return "", "", errors.New("cant find source user")
	}
	err = s.repo.MakeFriend(ctx, source, target)
	if err != nil {
		return "", "", err
	}
	err = s.repo.MakeFriend(ctx, target, source)
	if err != nil {
		return "", "", err
	}
	return targetUser.Name, sourceUser.Name, nil
}

// Delete user.
func (s *service) DeleteUser(ctx context.Context, id uint64) error {
	err := s.repo.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	err = s.repo.DeleteUserFromFriends(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// GetFriends by ids
func (s *service) GetFriends(ctx context.Context, id uint64) ([]model.User, error) {
	ids, err := s.repo.GetFriends(ctx, id)
	if err != nil {
		s.l.Info(err)
		return nil, err
	}
	if len(ids) == 0 {
		err := errors.New("user dont have friends")
		return []model.User{}, err
	}
	var friends []model.User
	for _, v := range ids {
		user, err := s.repo.GetUser(ctx, uint64(v))
		if err != nil {
			s.l.Info(err, "GetUser err in svc form repo")
			return nil, nil
		}
		friends = append(friends, user)
	}
	return friends, nil
}

// UpdateAge at user.
func (s *service) UpdateUserAge(ctx context.Context, id, age uint64) error {
	return s.repo.UpdateUserAge(ctx, id, age)
}
