package golang_embed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

var logo []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("Logo_new..png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}