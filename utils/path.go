package utils

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
)

var RunWithoutGoenv bool = func() bool {
	run_without_goenv, _ := strconv.ParseBool(os.Getenv("RUN_WITHOUT_GOENV"))
	return run_without_goenv
}()

var Buildpath string = func() string {
	// Assume the executable exists directly under BUILDPATH/bin
	buildpath, err := filepath.Abs(filepath.Join(filepath.Dir(os.Args[0]), ".."))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("BuildPath buildpath='%v'\n", buildpath)
	return buildpath
}()

