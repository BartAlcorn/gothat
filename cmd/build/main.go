package main

import (
	"github.com/bartalcorn/gothat/pkg/build"
	"github.com/bartalcorn/gothat/pkg/embedded"
)

func main() {
	scripts := embedded.Scripts

	files := map[string][]string{}
	for _, script := range scripts {
		files[script.Folder] = append(files[script.Folder], "./pkg/embedded/assets/js/"+script.Folder+"/"+script.Name)
	}

	for folder := range files {
		build.Concat("./pkg/embedded/assets/js/", folder+".js", files[folder])
	}

}
