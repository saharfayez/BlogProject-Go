package server

func Serve() {
	e := RegisterRoutes()
	e.Logger.Fatal(e.Start(":8080"))
}
