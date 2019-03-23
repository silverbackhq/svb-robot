// Copyright 2019 Silverbackhq. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/silverbackhq/svb-robot/internal/app/cmd"
	"github.com/silverbackhq/svb-robot/internal/app/controller"
	"github.com/silverbackhq/svb-robot/internal/app/middleware"
	"github.com/silverbackhq/svb-robot/internal/app/pkg/utils"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {

	var exec string
	var configFile string

	flag.StringVar(&exec, "exec", "", "exec")
	flag.StringVar(&configFile, "config", "config.prod.yml", "config")
	flag.Parse()

	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Sprintf(
			"Error while loading config file [%s]: %s",
			configFile,
			err.Error(),
		))
	}

	if exec != "" {
		switch exec {
		case "health":
			cmd.HealthStatus()
		default:
			utils.PrintCommands()
		}
		return
	}

	if viper.GetString("app.mode") == "prod" {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
		f, _ := os.Create(fmt.Sprintf("%s/gin.log", viper.GetString("log.path")))
		gin.DefaultWriter = io.MultiWriter(f)
	}

	r := gin.Default()
	r.Use(middleware.Correlation())
	r.Use(middleware.Logger())
	r.GET("/", controller.Index)
	r.GET("/_healthcheck", controller.HealthCheck)
	r.GET("/login", controller.Login)
	r.GET("/auth", controller.Auth)
	r.POST("/listen", controller.Listen)
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.String(http.StatusNoContent, "")
	})

	if viper.GetBool("app.tls.status") {
		r.RunTLS(
			fmt.Sprintf(":%s", strconv.Itoa(viper.GetInt("app.port"))),
			viper.GetString("app.tls.pemPath"),
			viper.GetString("app.tls.keyPath"),
		)
	} else {
		r.Run(
			fmt.Sprintf(":%s", strconv.Itoa(viper.GetInt("app.port"))),
		)
	}
}
