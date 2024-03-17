package user

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateUser inserts a new user into the database.
func CreateUser(db *sql.DB, u User) error {
    const query = `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`
    _, err := db.Exec(query, u.Username, u.Email, u.Password)
    return err
}

// GetUser retrieves a single user by ID.
func GetUser(db *sql.DB, id int64) (User, error) {
    const query = `SELECT id, username, email FROM users WHERE id = $1`
    var u User
    err := db.QueryRow(query, id).Scan(&u.ID, &u.Username, &u.Email)
    return u, err
}

// UpdateUser updates a user's information.
func UpdateUser(db *sql.DB, u User) error {
    const query = `UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4`
    _, err := db.Exec(query, u.Username, u.Email, u.Password, u.ID)
    return err
}

// DeleteUser removes a user from the database.
func DeleteUser(db *sql.DB, id int64) error {
    const query = `DELETE FROM users WHERE id = $1`
    _, err := db.Exec(query, id)
    return err
}

//HANDLERS
func CreateUserHandler(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var u User
        if err := c.ShouldBindJSON(&u); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if err := CreateUser(db, u); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"status": "user created"})
    }
}

func GetUserHandler(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, err := strconv.ParseInt(c.Param("id"), 10, 64)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
            return
        }
        u, err := GetUser(db, id)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, u)
    }
}

func UpdateUserHandler(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var u User
        if err := c.ShouldBindJSON(&u); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if err := UpdateUser(db, u); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"status": "user updated"})
    }
}

func DeleteUserHandler(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, err := strconv.ParseInt(c.Param("id"), 10, 64)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
            return
        }
        if err := DeleteUser(db, id); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"status": "user deleted"})
    }
}

func RegisterUserRoutes(g *gin.RouterGroup, db *sql.DB) {
    g.POST("/users", CreateUserHandler(db))
    g.GET("/users/:id", GetUserHandler(db))
    g.PUT("/users", UpdateUserHandler(db))
    g.DELETE("/users/:id", DeleteUserHandler(db))
}
