// Copyright 2023 The Perses Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package datasource

import (
	"net/url"

	"github.com/perses/perses/pkg/model/api/v1/datasource/http"
)

// Prometheus is only used for testing purpose.
// It doesn't reflect the nature of the actual prometheus datasource
type Prometheus struct {
	DirectURL *url.URL   `json:"direct_url,omitempty" yaml:"direct_url,omitempty"`
	Proxy     http.Proxy `json:"proxy" yaml:"proxy"`
}
