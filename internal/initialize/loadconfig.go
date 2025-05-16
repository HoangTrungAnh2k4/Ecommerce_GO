// initialize/config.go
package initialize

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// Đọc file config
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}

	fmt.Println("Loaded port:", viper.Get("server.port")) // Debug
}
