package main

import (
	"encoding/base64"
	//"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	//"strconv"
)

func handleCatalogPiece(w http.ResponseWriter, r *http.Request, fname string, base64Encode bool) {
	log.Println("-->> Process file ", fname)
	fin, err := os.Open(fname)
	defer fin.Close()
	if err == nil {
		buf := make([]byte, 10240)
		for {
			n, err := fin.Read(buf)
			if err == nil && n != 0 {
				var contents string
				if base64Encode {
					contents = base64.StdEncoding.EncodeToString(buf[:n])
				} else {
					contents = string(buf[:n])
				}
				log.Println(contents)
				io.WriteString(w, contents)
			} else {
				break
			}
		}
	} else {
		log.Fatal(err)
	}
}

func handleCatalog(w http.ResponseWriter, r *http.Request) {
	log.Println("****************************")
	log.Println("-->> Handle get catalog requst...")
	handleCatalogPiece(w, r, "catalogPrefix.xml", false)
	handleCatalogPiece(w, r, "catalogContents.xml", true)
	handleCatalogPiece(w, r, "catalogSuffix.xml", false)
	log.Println("-->> End of catalog requst...")
	log.Println("****************************")
}

/*
func handlePackage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	filename := params["filename"]
	fin, err := os.Open(filename)
	defer fin.Close()
	if err == nil {
		w.Header().Add("Content-Type", "application/octet-stream")
		fileInfo, _ := fin.Stat()
		w.Header().Add("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))
		w.Header().Add("Content-Disposition", "attachment; filename="+filename)
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
*/

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/soap/rpc", handleCatalog).Methods("POST")
	//rtr.HandleFunc("/downloads/{filename:.*}", handlePackage).Methods("GET")
	http.Handle("/", rtr)
	log.Println("Server starts to listen on port 443...")
	err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}
