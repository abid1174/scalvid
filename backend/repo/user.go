package repo

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserRepository interface {
	CreateUser(user User) (User, error)
	Login(email string, password string) string
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) (User, error) {
	log.Println("Create User ", user)
	rows, err := r.db.NamedQuery(`
		INSERT INTO users (first_name, last_name, email, password)
		VALUES (${FirstName}, ${LastName}, ${Email}, ${Password})
		RETURNING id
	`, user)

	log.Println(rows)

	if err != nil {
		return User{}, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(&user)
	}
	return user, err
}

func (r *userRepository) Login(email string, password string) string {
	accessToken := ""

	return accessToken
}
