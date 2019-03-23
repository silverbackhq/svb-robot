// Copyright 2019 Silverbackhq. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/silverbackhq/svb-robot/internal/app/pkg/github"
	"github.com/spf13/viper"
	"net/http"
)

// Auth controller
func Auth(c *gin.Context) {

	githubOauth := &github.OAuthApp{
		ClientID:     viper.GetString("github.client_id"),
		RedirectURI:  viper.GetString("github.redirect_uri"),
		AllowSignup:  viper.GetString("github.allow_signup"),
		Scopes:       viper.GetStringSlice("github.scopes"),
		ClientSecret: viper.GetString("github.client_secret"),
	}

	state, err := c.Cookie("gh_oauth_state")

	if err == nil && state != "" {
		githubOauth.SetState(state)
	}

	ok, err := githubOauth.FetchAccessToken(
		c.DefaultQuery("code", ""),
		c.DefaultQuery("state", ""),
	)

	if ok && err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":      "ok",
			"accessToken": githubOauth.GetAccessToken(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "not ok",
			"error":  err.Error(),
		})
	}
}
