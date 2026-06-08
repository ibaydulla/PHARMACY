func main() {

	connStr := fmt.Sprintf(
		"user=%s dbname=%s host=%s password=%s sslmode=disable connect_timeout=5",
		"postgres",
		"pharmacy_db",
		"localhost",
		"5432",
	)

	utils.ConnectDB(connStr)
	defer utils.GetDB().Close(context.Background())

	r := gin.Default()

	rg := r.Group("/api")

	// Auth
	controllers.AuthRoute(rg)

	// Existing Routes
	controllers.UserRoute(rg)
	controllers.PharmacyRoute(rg)
	controllers.MedicinesRoute(rg)
	controllers.OrderRoute(rg)
	controllers.CategoryRoute(rg)

	if err := r.Run(":8080"); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}