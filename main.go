package main
import (
  "fmt"
  "net/http"
  "os"
)
func handler(w http.ResponseWriter, r *http.Request){
  port := os.Getenv("PORT")
  if port=="" { port = "8009" }
  fmt.Fprintf(w, `{"msg":"Hello from SAIF","port":%s}`, port)
}
func main(){
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8009", nil)
}
