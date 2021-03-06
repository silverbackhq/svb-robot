// Copyright 2019 Silverbackhq. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package event

import (
	"encoding/json"
)

// GithubAppAuthorization event received any time someone revokes their authorization of a GitHub App. GitHub Apps receive this webhook by default and cannot unsubscribe from this event.
type GithubAppAuthorization struct {
}

// LoadFromJSON update object from json
func (e *GithubAppAuthorization) LoadFromJSON(data []byte) (bool, error) {
	err := json.Unmarshal(data, &e)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ConvertToJSON convert object to json
func (e *GithubAppAuthorization) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&e)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
