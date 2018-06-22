package utils

import (
	"fmt"
	"github.com/hoonrepo/go-SimpleServer/params"
	"gopkg.in/urfave/cli.v2"
	"os"
	"path/filepath"
	"strconv"
)

const (
	PORT_NAME = "port"
	PATH_NAME = "path"
)

var (
	PortFlag = &cli.IntFlag{
		Name:        PORT_NAME,
		Aliases:     []string{"p"},
		Value:       params.DEFULT_PORT,
		Usage:       "simplehttp listening port",
		DefaultText: strconv.Itoa(params.DEFULT_PORT),
		//Destination:&port,

	}

	PathFlag = &cli.StringFlag{
		Name:        PATH_NAME,
		Aliases:     []string{"t"},
		Value:       currentDir(),
		Usage:       "simplehttp mapped path",
		DefaultText: ".",
		//Destination:&pathDir,
	}
)

func currentDir() string {
	cDir, err := filepath.Abs(".")

	//fmt.Printf("cDir=%s",cDir)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return ""
	}
	return cDir
}
