package utils

import (
	"bufio"
	"fmt"
	"github.com/hoonrepo/go-SimpleServer/params"
	"gopkg.in/urfave/cli.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

var (
	HttpCmd = &cli.Command{
		Name:    "http",
		Aliases: []string{"ht"},
		Usage:   "create a simplehttp with http",
		Action:  launchServer,
		Flags: []cli.Flag{
			PortFlag,
			PathFlag,
		},
	}
)

func launchServer(ctx *cli.Context) error {

	port := ctx.String(PORT_NAME)

	path, _ := filepath.Abs(ctx.String(PATH_NAME))

	fmt.Printf("Serving HTTP on localhost path %s port %s ...\n", path, port)

	strPort := fmt.Sprintf(":%s", port)

	header := http.Header{}
	header.Add("Access-Control-Allow-Origin", "*")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Printf("%s ", r.Host)
		log.Printf("\"%s %s %s\" %d", r.Method, r.URL.Path, r.Proto, http.StatusOK)

		/*
					self.send_response(HTTPStatus.OK)
			        self.send_header("Content-type", "text/html; charset=%s" % enc)
			        self.send_header("Content-Length", str(len(encoded)))
			        self.end_headers()
		*/

		header.Add("Content-type", "text/html; charset=utf-8")

		netPath := r.URL.Path[0:]

		finalPath := path + netPath

		if strings.Index(netPath, "favicon") > 0 {
			return
		}
		fileInfo, err := os.Stat(finalPath)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			log.Printf("%s 'Stat' is error", finalPath)
			fmt.Fprint(w, "request's path occured a error")
		} else {
			if fileInfo.IsDir() {

				displayContent := params.SimpleListDirPage(netPath, finalPath)
				fmt.Fprint(w, displayContent)
			} else {
				buffer, errReadF := ioutil.ReadFile(finalPath)
				if errReadF == nil {
					writer := bufio.NewWriter(w)
					writer.Write(buffer)
					writer.Flush()

				}else{
					fmt.Fprint(os.Stderr, errReadF)
					log.Printf("%s ", errReadF.Error())
				}
				//file,errFile := os.OpenFile(finalPath,os.O_RDONLY,os.ModePerm)
				//
				//if(errFile != nil){
				//	buf := []byte{}
				//	reader := bufio.NewReader(file)
				//
				//	len,errRead := reader.Read(buf)
			}

		}

		w.WriteHeader(http.StatusOK)

		//fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	})
	http.ListenAndServe(strPort, nil)

	//pathtest()
	return nil
}

func pathtest() {

	pot, _ := filepath.Abs(".")
	log.Println("pot:", pot)

	file, _ := os.Getwd()
	log.Println("current path:", file)

	file, _ = exec.LookPath(os.Args[0])
	log.Println("exec path:", file)

	dir, _ := path.Split(file)
	log.Println("exec folder relative path:", dir)

	os.Chdir(dir)
	wd, _ := os.Getwd()
	log.Println("exec folder absolute path:", wd)
}
