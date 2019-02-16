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
package dir

import (
	"os"
	"log"
	"strings"
)

func ScanSubdir(path string, handle func(os.FileInfo), filter func(os.FileInfo) bool)  {
	fd, err := os.Open(path)
	if err != nil {
		return
	}
	defer fd.Close()

	list, err := fd.Readdir(-1)
	if err != nil {
		return
	}

	for _, d := range list {
		if (filter == nil || filter(d)) {
			handle(d)
		}
	}

	return
}

func MakeDirAll(path string) {

	log.Println(path[0:strings.LastIndex(path, "/")])
	os.MkdirAll(path[0:strings.LastIndex(path, "/")], os.ModeDir)
}