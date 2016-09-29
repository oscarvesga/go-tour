package main

import( 
	"net/http"
	"io/ioutil"
	"strings"
	"fmt"
	"os"
	"log"
)

/*
*   Manejo de errores
*/
func check(e error) {
    if e != nil {
        log.Fatal(e)
        panic(e)
    }
}

/*
*   Funcion main del programa
*/
func main(){
	resp, err := http.Get("http://google.com/")
	check(err)
	// cierra el body al final
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	
	fName := strings.Split(resp.Request.URL.Host,".")[1]
	fmt.Println(fName)
	
    fo, err := os.Create(fName + ".txt")
    check(err)
    // close output file 
    defer func() {
        err := fo.Close()
        check(err)
    }()
    
    n2, err := fo.Write(body)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)
}