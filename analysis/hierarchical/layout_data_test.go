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

package hierarchical

import (
	"fmt"
	"strconv"
	"testing"
)

func TestToJsonString(t *testing.T)  {
	var (
		d LayoutData
		nodes []LayoutNode
		links []LayoutLink
		groups []LayoutGroup
	)
	nodes = make([]LayoutNode, 10)
	for index := 0; index < 10; index++ {
		nodes[index] = LayoutNode{
			index,
			"name"+strconv.Itoa(index),
			60,
			60,
		}
	}

	d = LayoutData{
		nodes,
		links,
		groups,
	}

	fmt.Println(d.ToJsonString())
}