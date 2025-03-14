// Copyright 2021 The Perses Authors
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

package api

// this file is just there to run the command generate
//go:generate go run generate.go -package=globaldatasource -plural=globaldatasources -kind=GlobalDatasource
//go:generate go run generate.go -package=datasource -plural=datasources -kind=Datasource -isProjectResource=true
//go:generate go run generate.go -package=project -plural=projects -kind=Project
//go:generate go run generate.go -package=dashboard -plural=dashboards -kind=Dashboard -isProjectResource=true
//go:generate go run generate.go -package=folder -plural=folders -kind=Folder -isProjectResource=true
