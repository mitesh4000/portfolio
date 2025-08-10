package main

import (
  "fmt"
  "net/http"
)

func server(){
  
}

func main (){


  http.handle("/home",loadHome)

http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})

log.Fatal(http.ListenAndServe(":8080", nil))

  fmt.Println("hi man");
}
