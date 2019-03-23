// Copyright 2019 Silverbackhq. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nbio/st"
	"github.com/spf13/viper"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"
)

// init setup stuff
func init() {

	dir, err := os.Getwd()

	if err != nil {
		panic(fmt.Sprintf(
			"Error while loading working directory: %s",
			err.Error(),
		))
	}

	basePath := strings.Split(dir, "/internal")

	configFile := fmt.Sprintf("%s/%s", basePath[0], "config.test.yml")

	viper.SetConfigFile(configFile)

	err = viper.ReadInConfig()

	if err != nil {
		panic(fmt.Sprintf(
			"Error while loading config file [%s]: %s",
			configFile,
			err.Error(),
		))
	}

	os.Setenv("PORT", strconv.Itoa(viper.GetInt("app.port")))
}

// TestHealthCheckController test case
func TestHealthCheckController(t *testing.T) {

	router := gin.Default()
	router.GET("/_healthcheck", HealthCheck)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/_healthcheck", nil)
	router.ServeHTTP(w, req)

	st.Expect(t, 200, w.Code)
	st.Expect(t, `{"status":"ok"}`, w.Body.String())
}
