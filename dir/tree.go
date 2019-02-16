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

// deal with subdir tree of path.
package dir

type DirTreeNode struct {
	Path string
	Name string
	Contained bool // This flag mark this node whether or not contained .go file.
	Parent *DirTreeNode
	Children []*DirTreeNode
	Files []string
}

func (dt *DirTreeNode)PutChild(child *DirTreeNode)  {
	dt.Children = append(dt.Children, child)
}


func (dt *DirTreeNode)PutFile(filename string)  {
	dt.Files = append(dt.Files, filename)
}

// backtracking Parent's Contained mark
func (dt *DirTreeNode)MarkContained() {
	dt.Contained = true
	var parent = dt.Parent
	if parent != nil {
		for {
			parent.Contained = true
			parent = parent.Parent
			if parent == nil || parent.Contained == true {
				break
			}
		}
	}
}
