package routes

import (
	"aimeechat_api/user"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db *sql.DB) {
	api := r.Group("/api")
{
	user.RegisterUserRoutes(api, db);
}
}