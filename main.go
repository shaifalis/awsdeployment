package main
import "net/http"
func main(){
	r := http.NewServeMux()
	r.HandleFunc("/",greetUser)
	http.ListenAndServe(":8080",r)
}

func greetUser(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Hello World"))
}