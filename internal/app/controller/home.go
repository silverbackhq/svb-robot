// Copyright 2019 Silverbackhq. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Index controller
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
