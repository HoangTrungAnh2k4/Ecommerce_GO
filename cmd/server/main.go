package main

import (
	"github.com/HoangTrungAnh2k4/Ecommerce_GO/internal/routers"
)

func main() {
	r := routers.NewRouter()
	r.Run()
}
