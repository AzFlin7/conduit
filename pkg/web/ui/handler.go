// Copyright © 2022 Meroxa, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ui

import (
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/conduitio/conduit/pkg/foundation/cerrors"
)

// Handler serves Conduit UI.
func Handler() (http.Handler, error) {
	uiAssetFS, err := newUIAssetFS()
	if err != nil {
		return nil, cerrors.Errorf("UI assets error: %w", err)
	}
	return gziphandler.GzipHandler(http.FileServer(uiAssetFS)), nil
}
