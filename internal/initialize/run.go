package initialize

func Run() {
	LoadConfig()

	InitLogger()

	InitMySQL()

	r := InitRouter()

	r.Run("3002")
}
