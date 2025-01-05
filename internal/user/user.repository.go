package user

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/PontnauGonzalo/go-rest-api/internal/domain"
)

type DB struct {
	Users     []domain.User
	MaxUserID uint64
}

type (
	UserRepository interface {
		Create(ctx context.Context, user *domain.User) error
		Update(ctx context.Context, id uint64, firstName, lastName, email string) error
		GetAll(ctx context.Context) ([]domain.User, error)
		GetById(ctx context.Context, id uint64) (*domain.User, error)
		Delete(ctx context.Context, id uint64) error
	}

	userRepository struct {
		db  *sql.DB
		log *log.Logger
	}
)

func NewRepository(db *sql.DB, logger *log.Logger) UserRepository {
	return &userRepository{
		db:  db,
		log: logger,
	}
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	sqlQuery := "INSERT INTO users(first_name, last_name, email) VALUES(?, ?, ?)"

	res, err := r.db.Exec(sqlQuery, user.FirstName, user.LastName, user.Email)
	if err != nil {
		r.log.Println(err.Error())
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		r.log.Println(err.Error())
		return err
	}

	user.ID = uint64(id)
	r.log.Println("user created with id: ", user.ID)

	return nil
}

func (r *userRepository) Update(ctx context.Context, id uint64, firstName, lastName, email string) error {
	var fields []string
	var values []interface{}

	if email != "" {
		fields = append(fields, "email = ?")
		values = append(values, email)
	}
	if firstName != "" {
		fields = append(fields, "first_name = ?")
		values = append(values, firstName)
	}
	if lastName != "" {
		fields = append(fields, "last_name = ?")
		values = append(values, lastName)
	}

	if len(fields) == 0 {
		r.log.Println()
		return ErrThereArentFields
	}

	values = append(values, id)
	sqlQuery := fmt.Sprintf("UPDATE users SET %s WHERE id = ? ", strings.Join(fields, ", "))

	res, err := r.db.Exec(sqlQuery, values...)
	if err != nil {
		r.log.Println(err.Error())
		return err
	}

	row, err := res.RowsAffected()
	if err != nil {
		r.log.Println(err.Error())
		return err
	}

	if row == 0 {
		err := ErrNotFound{ID: id}
		r.log.Println(err.Error())
		return err
	}

	r.log.Println("user updated id: ", id)
	return nil
}

func (r *userRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	var users []domain.User

	sqlQuery := "SELECT id, email, first_name, last_name FROM users"

	rows, err := r.db.Query(sqlQuery)
	if err != nil {
		r.log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u domain.User

		if err := rows.Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName); err != nil {
			r.log.Println(err.Error())
			return nil, err
		}

		users = append(users, u)
	}
	return users, nil
}

func (r *userRepository) GetById(ctx context.Context, id uint64) (*domain.User, error) {
	var u domain.User

	sqlQuery := "SELECT id, email, first_name, last_name FROM users WHERE id = ?"

	if err := r.db.QueryRow(sqlQuery, id).Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound{ID: id}
		}
		return nil, err
	}

	r.log.Println("user get user with id:", id)

	return &u, nil
}
func (r *userRepository) Delete(ctx context.Context, id uint64) error {
	sqlQuery := "DELETE FROM users WHERE id = ?"
	res, err := r.db.Exec(sqlQuery, id)

	if err != nil {
		return err
	}

	row, err := res.RowsAffected()
	if err != nil {
		r.log.Println(err.Error())
		return err
	}

	if row == 0 {
		err := ErrNotFound{ID: id}
		r.log.Println(err.Error())
		return err
	}

	r.log.Println("user delete with id:", id)

	return nil
}
