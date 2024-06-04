package internals

import (
	"ETaalim/internals/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RunServer() error {
	r := gin.Default()

	// all app routes
	routes.AuthRoutes(r)
	routes.StudentRoutes(r)
	routes.UserRoutes(r)

	err := r.Run(":8080")
	if err != nil {
		return fmt.Errorf("could not run server: %v", err)
	}

	// Applying migrations
	// if err := migrate.MakeMigrations(); err != nil {
	// 	panic(err)
	// }

	return nil
}
