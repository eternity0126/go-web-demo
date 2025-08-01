package test

import (
	"fmt"
	"github.com/spf13/viper"
	"gogofly/utils"
	"testing"
)

func initConfig() {
	viper.SetConfigName("settings")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../conf/")
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Sprintf("Load config error: %s\n", err.Error()))
	}

	fmt.Println(viper.GetString("server.port"))
}

func TestJwtUtils(t *testing.T) {
	initConfig()
	token, _ := utils.GenerateToken(1, "zs")
	fmt.Println(token)

	iJwtCustomClaims, err := utils.ParseToken(token)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	fmt.Println(iJwtCustomClaims)
}
