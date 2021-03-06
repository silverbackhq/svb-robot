// Copyright 2019 Silverbackhq. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/silverbackhq/svb-robot/internal/app/driver"
	"github.com/silverbackhq/svb-robot/internal/app/pkg/logger"
)

// HealthStatus check the current app health. Make it compatible with process managers like systemd & docker
func HealthStatus() (bool, error) {
	redisConnection := driver.NewRedisDriver()

	result, err := redisConnection.Connect()

	if !result {
		logger.Fatalf(
			"I am not ok: Error while connecting to redis: %s",
			err.Error(),
		)
		return false, err
	}

	logger.Info("I am ok")
	return true, nil
}
