package api

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type User struct {
	ID       int64  `db:"id"`
	Email    string `db:"email"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type UserRepository interface {
	Create(user *User) error
	//FindByEmail(email string) (*User, error)
	FindByUsername(username string) (*User, error)
}

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (repo *UserRepo) Create(user *User) error {
	query := "INSERT INTO users(email, username, password) VALUES($1, $2, $3) RETURNING id"
	row := repo.db.QueryRow(query, user.Email, user.Username, user.Password)

	err := row.Scan(&user.ID)
	if err != nil {
		fmt.Printf("Could not find user's ID.\n Err message:\n, %s\n", err)
		return err
	}

	return nil
}

/*
func CreateUser(userRepo *UserRepo) echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.FormValue("email")
		username := c.FormValue("username")
		password := c.FormValue("password")
		var id int

		query := "INSERT INTO users(email, username, password) VALUES($1, $2, $3) RETURNING id"
		userRepo.db.QueryRow(query, email, username, password)

	}
}*/

func (repo *UserRepo) FindByUsername(username string) (*User, error) {
	query := `SELECT * FROM users WHERE username = $1`

	user := &User{}
	err := repo.db.Get(user, query, username)

	/*
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
	*/

	if err != nil {
		fmt.Printf("%s\n", err)
	}

	return user, nil
}

func CreateUser(userRepo UserRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		// parse request body and create a User object
		user := &User{
			//ID:       uuid.New().String(), // generate a unique ID using the uuid package
			Email:    c.FormValue("email"),
			Username: c.FormValue("username"),
			Password: c.FormValue("password"),
		}

		// call the Create method on the UserRepository to save the user to the database
		err := userRepo.Create(user)
		if err != nil {
			// handle error, e.g. return an HTTP 500 error
			fmt.Println("User's creating failed!")
		}

		// return a success response with the created user object
		//return c.JSON(http.StatusOK, user)
		fmt.Println("Successfully signedup!")
		return c.String(http.StatusOK, "Successfully signedup!")

	}
}

func LoginHandler(userRepo UserRepository) echo.HandlerFunc {
	return func(c echo.Context) error {

		username := c.FormValue("username")
		password := c.FormValue("password")
		fmt.Printf("%s,%s\n", username, password)
		user, err := userRepo.FindByUsername(username)
		/*if err != nil || user == nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
			return echo.ErrUnauthorized
		}*/
		if err != nil {
			//fmt.Printf("%s\n", err)
			fmt.Println("Fetching a user from db failed!")
		}

		if user.Password != password {
			return echo.ErrUnauthorized
		}
		fmt.Println("Successfully logged in!")
		return c.String(http.StatusOK, "Successfully logged in!")
	}
}
