package repository

import (
	"context"
	"database/sql"

	"github.com/FahemHakikiKhaya/crtvns-backend-go/helper"
	"github.com/FahemHakikiKhaya/crtvns-backend-go/model"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

// Delete implements UserRepository.
func (u *UserRepositoryImpl) Delete(ctx context.Context, userId int) {
	tx, err := u.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	SQL := "delete from users where id = $1"
	_, err = tx.ExecContext(ctx, SQL, userId)
	helper.PanicIfError(err)
}

// FindAll implements UserRepository.
func (u *UserRepositoryImpl) FindAll(ctx context.Context) []model.User {
	tx, err := u.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	SQL := "select * from users"

	result, errQuery := tx.QueryContext(ctx, SQL)

	helper.PanicIfError(errQuery)
	
	defer result.Close()

	var users []model.User

	for result.Next() {
		user := model.User{}
		err := result.Scan(&user.Id,&user.Email,&user.Name,&user.Password, &user.CreatedAt)
		helper.PanicIfError(err)

		users = append(users, user)
	}

	return users
}

// FindById implements UserRepository.
func (u *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (model.User, error) {
	tx, err := u.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	SQL := "SELECT id, email, password, name, created_at FROM users WHERE email = $1"

	result := tx.QueryRowContext(ctx,SQL, email)

	var user model.User

	err = result.Scan(&user.Id,&user.Email,&user.Password,&user.Name, &user.CreatedAt)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return user,nil
		} else {
			helper.PanicIfError(err)
			return user, err
		}
	}
	
	return user, err
}

// FindById implements UserRepository.
func (u *UserRepositoryImpl) FindById(ctx context.Context, userId int) (model.User, error) {
	tx, err := u.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	SQL := "SELECT * FROM users where id = $1"

	result := tx.QueryRowContext(ctx,SQL, userId)

	var user model.User

	err = result.Scan(&user.Id,&user.Email,&user.Password,&user.Name, &user.CreatedAt)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return user,nil
		} else {
			helper.PanicIfError(err)
			return user, err
		}
	}
	
	return user, err
}

// Save implements UserRepository.
func (u *UserRepositoryImpl) Save(ctx context.Context, user model.User) (model.User, error) {
	tx, err := u.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "insert into users(name, email, password) values($1, $2, $3)"
	row := tx.QueryRowContext(ctx, SQL, user.Name, user.Email, user.Password)

	var newUser model.User
	err = row.Scan(&newUser.Id, &newUser.Name, &newUser.Email)

	if err != nil {
		return model.User{}, err
	}

	return newUser, nil
}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(ctx context.Context, user model.User) {
	tx, err := u.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "update users set name=$1 where id=$2"

	_,err = tx.ExecContext(ctx, SQL, user.Name, user.Id)
	helper.PanicIfError(err)
}

func NewUserRepository(Db *sql.DB) UserRepository {
	return &UserRepositoryImpl{DB: Db}
}
