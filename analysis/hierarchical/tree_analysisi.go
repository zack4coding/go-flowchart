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

// deal with subdir of path.
// referred to go/parses ParseDir func.
package hierarchical

import (
	"log"
	"go-flowchart/dir"
	"go/ast"
	"go/parser"
	"go/token"
)

// In order to find import links.
var  pathIndexMap map[string]int

func Analyze(tree *dir.DirTreeNode) LayoutData {
	pathIndexMap = make(map[string]int)
	var layout = LayoutData{}
	// layout.Links = append(layout.Links, LayoutLink{Source:1,Target:2})
	scanTreeNode(tree, &layout, nil)
	makeLinks(tree, &layout)
	return layout
}

func scanTreeNode(node *dir.DirTreeNode, layout *LayoutData, group *LayoutGroup) {

	width := len(node.Name)
	var layoutNode = LayoutNode{Name:node.Name,Width:width * 15 + 50,Height:70}
	layout.Nodes = append(layout.Nodes, layoutNode)
	layoutNode.Index = len(layout.Nodes) -1
	if group != nil {
		group.Leaves = append(group.Leaves, layoutNode.Index)
	}

	log.Println(node.Path)
	pathIndexMap["\""+node.Path+"\""] = layoutNode.Index

	if node.Children != nil && len(node.Children) > 0 {
		var nextId = 0
		if group != nil {
			nextId = group.Index
		}
		var newGroup = LayoutGroup{Index:nextId}
		for _, child := range node.Children {
			// recursion
			scanTreeNode(child, layout, &newGroup)
		}
		// if group != nil {
		// 	group.Groups = append(group.Groups, nextId)
		// }
		if len(newGroup.Leaves) > 1 {
			layout.Groups = append(layout.Groups, newGroup)
		}
	}
}

func makeLinks(node *dir.DirTreeNode, layout *LayoutData) {

	if node.Files != nil || len(node.Files) > 0 {
		
		if _, ok := pathIndexMap["\""+node.Path+"\""]; ok {
			source := pathIndexMap["\""+node.Path+"\""]
			
			var (
				fset        = token.NewFileSet()
			)
			for _, filename := range node.Files {
				f, _ := parser.ParseFile(fset, filename, nil, 0)
				ast.Inspect(f, func(n ast.Node) bool {
					switch nodeType := n.(type) {
					// skip comments
					case *ast.CommentGroup, *ast.Comment:
						return false
					case *ast.ImportSpec:
						
						if _, ok := pathIndexMap[nodeType.Path.Value]; ok {
							
							target := pathIndexMap[nodeType.Path.Value]
							log.Println(nodeType.Path.Value, source, target)
							layout.Links = append(layout.Links, LayoutLink{Source:source, Target:target})
						}
						return true
					}

					return true
				})
			}
		}
	}

	if node.Children != nil && len(node.Children) > 0 {
		for _, child := range node.Children {
			makeLinks(child, layout)
		}
	}

}
