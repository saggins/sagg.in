package main

func initializeRoutes() {
	// INDEX is special... Its a special boi
	router.GET("/", showIndexPage)
}
