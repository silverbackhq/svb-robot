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

// Login controller
func Login(c *gin.Context) {

	githubOauth := &github.OAuthApp{
		ClientID:     viper.GetString("github.client_id"),
		RedirectURI:  viper.GetString("github.redirect_uri"),
		AllowSignup:  viper.GetString("github.allow_signup"),
		Scopes:       viper.GetStringSlice("github.scopes"),
		ClientSecret: viper.GetString("github.client_secret"),
	}

	state, err := c.Cookie("gh_oauth_state")

	if err != nil || state == "" {
		githubOauth.GenerateState()
		c.SetCookie("gh_oauth_state", githubOauth.GetState(), 3600, "/", viper.GetString("app.domain"), true, true)
	} else {
		githubOauth.SetState(state)
	}

	c.JSON(http.StatusOK, gin.H{
		"title":    "Silverback Robot",
		"oauthURL": githubOauth.BuildAuthorizeURL(),
	})
}
