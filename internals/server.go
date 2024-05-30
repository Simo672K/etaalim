package internals

import (
	"ETaalim/internals/data/migrate"
)

func RunServer() error {
	//r := gin.Default()
	//
	//// all app routes
	//routes.StudentRoutes(&r.RouterGroup)
	//
	//err := r.Run(":8080")
	//if err != nil {
	//	return fmt.Errorf("could not run server: %v", err)
	//}

	// Applying migrations
	if err := migrate.MakeMigrations(); err != nil {
		panic(err)
	}

	return nil
}
