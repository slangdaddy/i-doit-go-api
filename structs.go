/*
	Go library for simple i-doit api usage

	Copyright (C) 2017 Carsten Seeger

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <http://www.gnu.org/licenses/>.

	@author Carsten Seeger
	@copyright Copyright (C) 2017 Carsten Seeger
	@license http://www.gnu.org/licenses/gpl-3.0 GNU General Public License 3
	@link https://github.com/cseeger-epages/i-doit-go-api
*/

package goidoit

// globals
var (
	id       int
	debug    bool = false
	insecure bool = false
)

// Api struct used for implementing the apiMethods interface
type Api struct {
	Url, Apikey, Username, Password, SessionId string
}

// Request implements i-doit api request structure
type Request struct {
	Version string      `json:"version"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      int         `json:"id"`
}

// Response implements i-doit api response structure
type Response struct {
	Jsonrpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Error   IdoitError  `json:"error"`
}

// GenericResponse implements a generic i-doit api response structure
// the map is used to handle type assertions
type GenericResponse struct {
	Jsonrpc string
	Result  []map[string]interface{}
	Error   IdoitError
}

// IdoitError implements i-doit api error structure
type IdoitError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Apikey used for requests
type Apikey struct {
	Apikey string `json:"apikey"`
}

// F1 implements an object filter type int or []int
type F1 struct {
	Data []int `json:"ids"`
}

// F2 implements an Object filter type string
type F2 struct {
	Data string `json:"title"`
}

// OF1 implements a more complex object type filter
// for ids and type
type OF1 struct {
	Title []int  `json:"ids"`
	Type  string `json:"type"`
}

// OF2 implements a more complex object type filter
// for title and type
type OF2 struct {
	Title string `json:"title"`
	Type  string `json:"type"`
}

// OSF1 implements object type only filter
type OSF1 struct {
	Type string `json:"type"`
}
