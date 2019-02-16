// Copyright 2019 The go-flowchart Authors
// This file is part of the go-flowchart library.
//
// The go-flowchart library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-flowchart library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-flowchart library. If not, see <http://www.gnu.org/licenses/>.

// This is the layout data what layout.html needed.
// referred to cola.js see <https://ialab.it.monash.edu/webcola/>.
package hierarchical

import (
	"encoding/json"
	"log"
)

type LayoutData struct {
	Nodes []LayoutNode `json:"nodes,omitempty"`
	Links []LayoutLink `json:"links,omitempty"`
	Groups []LayoutGroup `json:"groups,omitempty"`
}

type LayoutNode struct {
	Index int `json:"-"`
	Name string `json:"name,omitempty"`
	Width int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

type LayoutLink struct {
	Source int `json:"source"`
	Target int `json:"target"`
}

type LayoutGroup struct {
	Index int `json:"-"`
	Leaves []int `json:"leaves,omitempty"`
	Groups []int `json:"groups,omitempty"`
}

func (d *LayoutData)ToJson() []byte {
	jsonBytes, err := json.Marshal(d)
	if err != nil {
		log.Panicln("layout data serialization error.")
	}
	return jsonBytes
}

func (d *LayoutData)ToJsonString() string {
	return string(d.ToJson())
}