package postgresql

import (
	"database/sql"
	"go-tests/integration"
)

var _ integration.UsersRepo = (*UsersPostgreSqlRepo)(nil)

type UsersPostgreSqlRepo struct {
	db *sql.DB
}

func NewPostgresqlRepo() (*UsersPostgreSqlRepo, error) {
	db, err := sql.Open("postgres", "user=user password=password dbname=postgres sslmode=disable")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT)")
	if err != nil {
		return nil, err
	}

	return &UsersPostgreSqlRepo{db: db}, nil
}

func (r *UsersPostgreSqlRepo) CreateUsers(users []integration.User) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO users (id,name) VALUES ($1,$2)")
	if err != nil {
		return err
	}

	for _, user := range users {
		if _, err := stmt.Exec(user.ID, user.Name); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *UsersPostgreSqlRepo) FindAllUsers() ([]integration.User, error) {
	rows, err := r.db.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []integration.User
	for rows.Next() {
		var user integration.User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
