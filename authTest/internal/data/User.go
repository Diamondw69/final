package data

import (
	"authTest/internal/validator"
	pb "authTest/pkg/proto"
	"context"
	"crypto/sha256"
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserModel struct {
	DB *sql.DB
}

var (
	ErrDuplicateEmail = errors.New("duplicate email")
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

func (m UserModel) Insert(user *pb.User) error {
	query := `
INSERT INTO users (name, email, password_hash,role,balance)
VALUES ($1, $2, $3, $4,$5)
RETURNING id`
	args := []any{user.Name, user.Email, user.Password.Hash, user.Role, user.Balance}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.Id)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}
	return nil
}

func (m UserModel) GetByEmail(email string) (*pb.User, error) {
	query := `
SELECT id, name, email, password_hash,  balance
FROM users
WHERE email = $1`
	var user pb.User
	password := pb.Password{
		PlainText: "",
		Hash:      []byte("auvhjkdfbhjlvadfjlhkgadskjlhadshjkasdkjhasdkhjb;asdjbkasdjkhasdj"),
	}
	user.Password = &password
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, email).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password.Hash,
		&user.Balance,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &user, nil
}

func (m UserModel) Update(user *pb.User) error {
	query := `
UPDATE users
SET name = $1, email = $2, password_hash = $3,role=$5,balance=$6
WHERE id = $5
RETURNING id`
	args := []any{
		user.Name,
		user.Email,
		user.Password.Hash,
		user.Role,
		user.Id,
		user.Balance,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.Id)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}

func (m UserModel) GetForToken(tokenScope, tokenPlaintext string) (*pb.User, error) {

	tokenHash := sha256.Sum256([]byte(tokenPlaintext))

	query := `
SELECT users.id, users.name, users.email, users.password_hash,users.role,users.balance
FROM users
INNER JOIN tokens
ON users.id = tokens.user_id
WHERE tokens.hash = $1
AND tokens.scope = $2
AND tokens.expiry > $3`

	args := []interface{}{tokenHash[:], tokenScope, time.Now()}
	var user pb.User
	password := pb.Password{
		PlainText: tokenPlaintext,
		Hash:      []byte("auvhjkdfbhjlvadfjlhkgadskjlhadshjkasdkjhasdkhjb;asdjbkasdjkhasdj")}
	user.Password = &password
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password.Hash,
		&user.Role,
		&user.Balance,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func Matches(plaintextPassword string, hash []byte) (bool, error) {
	p := pb.Password{}
	p.Hash = hash
	err := bcrypt.CompareHashAndPassword(p.Hash, []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "username must be provided")
	v.Check(validator.Matches(email, validator.EmailRX), "email", "must be a valid email address")
}
func ValidatePasswordPlaintext(v *validator.Validator, password string) {
	v.Check(password != "", "password", "password must be provided")
	v.Check(len(password) >= 8, "password", "password must be at least 8 bytes long")
	v.Check(len(password) <= 72, "password", "password must not be more than 72 bytes long")
}
func ValidateUser(v *validator.Validator, user *pb.User) {
	v.Check(user.Name != "", "name", "must be provided")
	v.Check(len(user.Name) <= 500, "name", "username must not be more than 500 bytes long")
	v.Check(len(user.Name) >= 5, "name", "username must  be more than 5 bytes long")

	ValidateEmail(v, user.Email)

	if user.Password.PlainText != "" {
		ValidatePasswordPlaintext(v, user.Password.PlainText)
	}

	if user.Password.Hash == nil {
		panic("missing password hash for user")
	}
}
