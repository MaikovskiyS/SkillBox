package service

import (
	"context"
	"errors"
	"skillbox/internal/domain/model"
	"skillbox/internal/transport/http/handler"

	"github.com/sirupsen/logrus"
)

//go:generate mockgen -source=user.go -destination=mocks/mock.go

type UserRepository interface {
	CreateUser(ctx context.Context, data model.User) (uint64, error)
	MakeFriend(ctx context.Context, source, target uint64) error
	GetUser(ctx context.Context, id uint64) (model.User, error)
	DeleteUser(ctx context.Context, id uint64) error
	DeleteUserFromFriends(ctx context.Context, id uint64) error
	GetFriends(ctx context.Context, id uint64) ([]int, error)
	UpdateUserAge(ctx context.Context, id uint64, age uint64) error
}

type service struct {
	l    *logrus.Logger
	repo UserRepository
}

//New...
func NewUserService(repo UserRepository, logger *logrus.Logger) handler.UserService {
	return &service{
		l:    logger,
		repo: repo,
	}
}

//Create...
func (s *service) CreateUser(ctx context.Context, data model.User) (uint64, error) {
	id, err := s.repo.CreateUser(ctx, data)
	if err != nil {
		s.l.Info(err)
	}
	return id, nil
}

//MakeFriend...
//TODO проверка на дружбу
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

//Delete...
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

//GetFriends...
func (s *service) GetFriends(ctx context.Context, id uint64) ([]model.User, error) {
	ids, err := s.repo.GetFriends(ctx, id)
	if err != nil {
		//	s.l.Info(err, "GetFriends err in svc from repo")
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

//UpdateAge...
func (s *service) UpdateUserAge(ctx context.Context, id, age uint64) error {
	return s.repo.UpdateUserAge(ctx, id, age)
}
