// Copyright 2019 Silverbackhq. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/silverbackhq/svb-robot/internal/app/pkg/utils"
	"strings"
)

// Correlation middleware
func Correlation() gin.HandlerFunc {
	return func(c *gin.Context) {
		corralationID := c.Request.Header.Get("X-Correlation-ID")
		if strings.TrimSpace(corralationID) == "" {
			c.Request.Header.Add("X-Correlation-ID", utils.GenerateUUID4())
		}
		c.Next()
	}
}
