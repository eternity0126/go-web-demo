package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("settings")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf/")
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Sprintf("Load config error: %s\n", err.Error()))
	}

	fmt.Println(viper.GetString("server.port"))
}
