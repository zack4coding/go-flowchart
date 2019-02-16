package analysis

import (
	"os"
	"log"
	"strings"
	"go-flowchart/dir"
)

var gopath = os.Getenv("GOPATH")

// Subdir tree of given path.
var dirTree *dir.DirTreeNode
// Cache tree data. Maybe stores in some DB in the future.
var trees map[string]*dir.DirTreeNode

func ParsePath(path string) *dir.DirTreeNode {
	// dirTree.Init(path)
	if trees == nil {
		trees = make(map[string]*dir.DirTreeNode)
	}

	if trees[path] == nil {
		dirTree = &dir.DirTreeNode{Path:path, Name:"root", Contained:false}
		trees[path] = dirTree
		log.Println(path)
		dir.ScanSubdir(absolutePath(path), handleFile, filter)
	}
	return trees[path]
}

func WCacheFile(path string, data []byte) string {
	cachePath := absolutePath("go-flowchart")+"/assets/.cache/"+path+".json"
	packagePath := ".cache/"+path+".json";
	dir.MakeDirAll(cachePath)
	file, err := os.Create(cachePath)
	if err !=nil {
		log.Println(err)
	}
	defer file.Close()
	file.Write(data)
	return packagePath
}

func handleFile(file os.FileInfo) {

	currentPath := strings.Join([]string{ dirTree.Path, file.Name() } ,"/")
	absolutePath := absolutePath(currentPath)
	if file.IsDir() {

		currentTreeNode := &dir.DirTreeNode{Path:currentPath, Name:file.Name(), Parent:dirTree}
		dirTree = currentTreeNode
		
		dir.ScanSubdir(absolutePath, handleFile, filter)
		dirTree = currentTreeNode.Parent

		if (currentTreeNode.Contained) {
			log.Println(currentPath)
			dirTree.PutChild(currentTreeNode)
		}


	} else {
		// This must be .go file due to the filter func 
		// So parse it.
		dirTree.PutFile(absolutePath)
		dirTree.MarkContained()
		log.Println(dirTree.Path, file.Name())
		
	}
}

func filter(file os.FileInfo) bool {

	if ( !file.IsDir() && !strings.HasSuffix(file.Name(), ".go") ) {
		return false
	}

	if ( strings.HasSuffix(file.Name(), ".git") ) {
		return false
	}

	if ( strings.HasSuffix(file.Name(), "_test.go") ) {
		return false
	}

	if ( strings.HasSuffix(file.Name(), "vendor") ) {
		return false
	}

	return true
}

func absolutePath(path string) string {
	absolutePath := gopath+"/src/"+path;
	return absolutePath
}