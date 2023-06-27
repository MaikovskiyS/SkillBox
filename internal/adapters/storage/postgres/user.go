package postgres

import (
	"context"
	"errors"
	"fmt"
	"skillbox/internal/domain/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, data model.User) (uint64, error)
	MakeFriend(ctx context.Context, source, target uint64) error
	GetUser(ctx context.Context, id uint64) (model.User, error)
	DeleteUser(ctx context.Context, id uint64) error
	DeleteUserFromFriends(ctx context.Context, id uint64) error
	GetFriends(ctx context.Context, id uint64) ([]int, error)
	UpdateUserAge(ctx context.Context, id uint64, age uint64) error
}
type userRepository struct {
	client Client
}

func NewUserRepository(cl *Client) UserRepository {
	return &userRepository{
		client: *cl,
	}
}

//Create user in database.
func (u *userRepository) CreateUser(ctx context.Context, data model.User) (uint64, error) {
	u.client.l.Info("Create User in DB")
	var id uint64
	err := u.client.Connection.QueryRow(ctx, "INSERT INTO users(name,age) values($1,$2) RETURNING ID", data.Name, data.Age).Scan(&id)
	if err != nil {
		return 0, errors.New("cant incert ")
	}
	return id, nil

}

// GetUser request user by id from users table.
func (u *userRepository) GetUser(ctx context.Context, id uint64) (model.User, error) {
	row := u.client.Connection.QueryRow(ctx, "select name,age from users where id=$1", id)
	var user model.User
	err := row.Scan(&user.Name, &user.Age)
	if err != nil {
		return model.User{}, errors.New("cant get user")
	}
	fmt.Println("User in GetUser:", user)
	return user, nil
}

// MakeFriend adds target_id and source_id in friends table.
func (u *userRepository) MakeFriend(ctx context.Context, source, target uint64) error {
	u.client.l.Info("Make firend in DB")
	tag, err := u.client.Connection.Exec(ctx, "INSERT INTO friends(user_id,friend_id) values($1,$2)", target, source)
	if err != nil {
		u.client.l.Info(err, "wrong insert in MakeFriend")
		return errors.New("cant make friend in DB")
	}
	if tag.RowsAffected() < 1 {
		return errors.New("cant make friendship")
	}
	return nil
}

// DeleteUser form users table by id.
func (u *userRepository) DeleteUser(ctx context.Context, id uint64) error {
	u.client.l.Info("Delete user form DB")
	tag, err := u.client.Connection.Exec(ctx, "DELETE FROM users WHERE id=$1", id)
	if err != nil {
		u.client.l.Info("cant delete user")
		return err
	}
	if tag.RowsAffected() < 1 {
		return errors.New("user not found")
	}
	return nil
}

// DeleteUserFromFriends by id.
func (u *userRepository) DeleteUserFromFriends(ctx context.Context, id uint64) error {
	u.client.l.Info("Delete friend form DB")
	tag, err := u.client.Connection.Exec(ctx, "DELETE FROM friends WHERE friend_id=$1 or user_id=$2", id, id)
	if err != nil {
		u.client.l.Info("cant delete friend")
		return err
	}
	if tag.RowsAffected() < 1 {
		return errors.New("friend not found")
	}
	return nil
}

// GetFriends by id.
func (u *userRepository) GetFriends(ctx context.Context, id uint64) ([]int, error) {
	u.client.l.Info("GetFriends in DB")
	rows, err := u.client.Connection.Query(ctx, "select friend_id from friends where user_id=$1", id)
	if err != nil {
		u.client.l.Info("cant get friends")
		return nil, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			u.client.l.Info("cant parse rows")
			return nil, err
		}
		ids = append(ids, id)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	fmt.Println("ids:", ids)
	return ids, nil
}

// UpdageAge at user by id.
func (u *userRepository) UpdateUserAge(ctx context.Context, id, age uint64) error {
	u.client.l.Info("ChangeAge in DB")
	tag, err := u.client.Connection.Exec(ctx, "UPDATE users SET age=$1 WHERE id=$2", age, id)
	if err != nil {
		u.client.l.Info("cant update age")
		return err
	}
	if !tag.Update() {
		return err
	}
	return nil
}
