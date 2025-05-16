package initialize

func Run() {
	LoadConfig()
	// port := viper.GetString("server.port")
	// fmt.Println("Server starting at port", port)

	InitLogger()

	r := InitRouter()

	r.Run()
}
