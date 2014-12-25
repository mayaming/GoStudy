package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func handleCatalog(w http.ResponseWriter, r *http.Request) {
	catalogOutFile := "catalogOut.xml"
	fin, err := os.Open(catalogOutFile)
	defer fin.Close()
	if err == nil {
		buf := make([]byte, 1024)
		for {
			n, _ := fin.Read(buf)
			if 0 != n {
				io.WriteString(w, string(buf[:n]))
			} else {
				break
			}
		}
	} else {
		io.WriteString(w, fmt.Sprint(err))
	}

}

func handlePackage(w http.ResponseWriter, r *http.Request) {
	pkg := "pkg.bin"
	fin, err := os.Open(pkg)
	defer fin.Close()
	if err == nil {
		w.Header().Add("Content-Type", "application/octet-stream")
		fileInfo, _ := fin.Stat()
		w.Header().Add("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))
		w.Header().Add("Content-Disposition", "attachment; filename=pkg.bin")
		buf := make([]byte, 1024)
		for {
			n, _ := fin.Read(buf)
			if 0 != n {
				w.Write(buf)
			} else {
				break
			}
		}
	} else {
		io.WriteString(w, fmt.Sprint(err))
	}

}

func main() {
	http.HandleFunc("/soap/rpc", handleCatalog)
	http.HandleFunc("/pkg", handlePackage)
	http.ListenAndServe(":8000", nil)
}
