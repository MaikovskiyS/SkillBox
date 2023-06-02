package postgres

import (
	"context"
	"errors"
	"fmt"
	"skillbox/internal/domain/model"
	"skillbox/internal/domain/service"
)

type userRepository struct {
	client Client
}

func NewUserRepository(cl *Client) service.UserRepository {
	return &userRepository{
		client: *cl,
	}
}

//Create...
func (u *userRepository) CreateUser(ctx context.Context, data model.User) (uint64, error) {
	u.client.Logger.Info("Create User in DB")
	var id uint64
	err := u.client.Connection.QueryRow(ctx, "INSERT INTO users(name,age) values($1,$2) RETURNING ID", data.Name, data.Age).Scan(&id)
	if err != nil {
		return 0, errors.New("cant incert ")
	}
	return id, nil

}

//GetUser...
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

//MakeFriend...
func (u *userRepository) MakeFriend(ctx context.Context, source, target uint64) error {
	u.client.Logger.Info("Make firend in DB")
	tag, err := u.client.Connection.Exec(ctx, "INSERT INTO friends(user_id,friend_id) values($1,$2)", target, source)
	if err != nil {
		u.client.Logger.Info(err, "problems with insert in MakeFriend")
		return errors.New("cant make friend in DB")
	}
	if tag.RowsAffected() < 1 {
		return errors.New("cant make friendship")
	}
	return nil
}

//DeleteUser...
func (u *userRepository) DeleteUser(ctx context.Context, id uint64) error {
	u.client.Logger.Info("Delete user form DB")
	fmt.Println("id", id)
	tag, err := u.client.Connection.Exec(ctx, "DELETE FROM users WHERE id=$1", id)
	if err != nil {
		u.client.Logger.Info("cant delete user")
		return err
	}
	if tag.RowsAffected() < 1 {
		return errors.New("user not found")
	}
	return nil
}

//DeleteUserFromFriends...
func (u *userRepository) DeleteUserFromFriends(ctx context.Context, id uint64) error {
	u.client.Logger.Info("Delete friend form DB")
	fmt.Println("id", id)
	tag, err := u.client.Connection.Exec(ctx, "DELETE FROM friends WHERE friend_id=$1 or user_id=$2", id, id)
	if err != nil {
		u.client.Logger.Info("cant delete friend")
		return err
	}
	if tag.RowsAffected() < 1 {
		return errors.New("friend not found")
	}
	return nil
}

//GetFriends...
func (u *userRepository) GetFriends(ctx context.Context, id uint64) ([]int, error) {
	u.client.Logger.Info("GetFriends in DB")
	rows, err := u.client.Connection.Query(ctx, "select friend_id from friends where user_id=$1", id)
	if err != nil {
		u.client.Logger.Info("cant get friends")
		return nil, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			u.client.Logger.Info("cant parse rows")
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

//UpdageAge...
func (u *userRepository) UpdateUserAge(ctx context.Context, id, age uint64) error {
	u.client.Logger.Info("ChangeAge in DB")
	tag, err := u.client.Connection.Exec(ctx, "UPDATE users SET age=$1 WHERE id=$2", age, id)
	if err != nil {
		u.client.Logger.Info("cant update age")
		return err
	}
	if !tag.Update() {
		return err
	}
	return nil
}
