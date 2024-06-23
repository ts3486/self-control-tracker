package routes

import (
	"database/sql"
	"sct-api/user"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db *sql.DB) {
	api := r.Group("/api")
{
	user.RegisterUserRoutes(api, db);
}
}