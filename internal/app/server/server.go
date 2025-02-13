package server

func Serve() {
	e := registerRoutes()
	e.Logger.Fatal(e.Start(":8080"))
}
