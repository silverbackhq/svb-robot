// Copyright 2019 Silverbackhq. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package event

import (
	"encoding/json"
)

// RepositoryImport event received any time a successful or unsuccessful repository import finishes for a GitHub organization or a personal repository.
type RepositoryImport struct {
}

// LoadFromJSON update object from json
func (e *RepositoryImport) LoadFromJSON(data []byte) (bool, error) {
	err := json.Unmarshal(data, &e)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ConvertToJSON convert object to json
func (e *RepositoryImport) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&e)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
