package user

import (
	"context"
	"database/sql"
	"net/http"
	userMethods "sct-api/user/methods"
	"strconv"

	"github.com/gin-gonic/gin"
)

//HANDLERS
func CreateUserHandler(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var u User
        if err := c.ShouldBindJSON(&u); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        q := userMethods.New(db) // create a new Queries instance

        params := userMethods.CreateUserParams{
            Username: u.Username,
            Email:    u.Email,
            Password: u.Password,
        }

        if err := q.CreateUser(context.Background(), params); err != nil {
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

        q := userMethods.New(db) // create a new Queries instance

        u, err := q.GetUser(c.Request.Context(), int32(id))
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, u)
    }
}

func UpdateUserHandler(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var u userMethods.UpdateUserParams
        if err := c.ShouldBindJSON(&u); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        q := userMethods.New(db)

        if err := q.UpdateUser(c.Request.Context(), u); err != nil {
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

        q := userMethods.New(db)

        if err := q.DeleteUser(c.Request.Context(), int32(id)); err != nil {
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
