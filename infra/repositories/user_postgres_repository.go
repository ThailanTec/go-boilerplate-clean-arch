package repositories

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"time"

	"github.com/ThailanTec/challenger/pousada/domain"
	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUsers() ([]*domain.User, error)
	GetUserByData(document string) (*domain.User, error)
	DeleteUser(id uuid.UUID) error
	UpdateUser(id uuid.UUID, user *domain.User) (*domain.User, error)
	GetUserByID(id uuid.UUID) (*domain.User, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) CreateUser(user *domain.User) error {
	user.ID = uuid.New()
	query := `
		INSERT INTO users (id, name, phone, document, created_at, updated_at)
		VALUES (:id, :name, :phone, :document, :created_at, :updated_at)
	`

	_, err := repo.db.NamedExec(query, user)
	return err
}

func (repo *userRepository) GetUsers() ([]*domain.User, error) {
	var users []*domain.User
	query := "SELECT id, name, phone, document, created_at, updated_at FROM users"

	err := repo.db.Select(&users, query)
	return users, err
}

func (repo *userRepository) GetUserByData(document string) (*domain.User, error) {
	var user domain.User
	query := "SELECT id, name, phone, document, created_at, updated_at FROM users WHERE document = $1"

	err := repo.db.Get(&user, query, document)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrGetUserByData
		}
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) GetUserByID(id uuid.UUID) (*domain.User, error) {
	var user domain.User
	query := "SELECT id, name, phone, document, created_at, updated_at FROM users WHERE id = $1"

	err := repo.db.Get(&user, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrIDNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) DeleteUser(id uuid.UUID) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := repo.db.Exec(query, id)
	return err
}

func (repo *userRepository) UpdateUser(id uuid.UUID, user *domain.User) (*domain.User, error) {
	user.UpdatedAt = time.Now()
	query := `
		UPDATE users
		SET name = :name, phone = :phone, document = :document, updated_at = :updated_at
		WHERE id = :id
		RETURNING id, name, phone, document, created_at, updated_at
	`

	user.ID = id // Garante que o ID seja passado corretamente
	rows, err := repo.db.NamedQuery(query, user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.StructScan(user); err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("user not found")
	}

	return user, nil
}
