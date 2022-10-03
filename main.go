package main

import(
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter,r *http.Request)  {
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w,"ParseForm err: %v",err)
		return
	}
	fmt.Fprintf(w,"POST req successful\n")
	name :=r.FormValue("name")
	address := r.FormValue("address")
	age := r.FormValue("age")
	fmt.Fprintf(w,"Name:%s\n",name)
	fmt.Fprintf(w,"Address:%s\n",address)
	fmt.Fprintf(w,"Age:%s\n",age)

}


func helloHandler(w http.ResponseWriter,r *http.Request)  {
if r.URL.Path != "/hello" {
	http.Error(w,"404 not Found",http.StatusNotFound)
	return
}
if r.Method !="GET"{
	http.Error(w,"Method not supported",http.StatusNotFound)
	return
}
fmt.Fprintf(w,"Hello")
}
	


func main(){
	// telling the web server to look into the "./static" folder
	file_server:=http.FileServer(http.Dir("./static"))
	// route to the root page
	http.Handle("/",file_server)
	// rout to the form page
	http.HandleFunc("/form",formHandler)
	// rout to the hello page
	http.HandleFunc("/hello",helloHandler)

	//http.HandleFunc("/age",ageHandler)

	fmt.Println("Starting server at port :7000")
	if err := http.ListenAndServe(":7000",nil); err != nil{
		log.Fatal(err)

	}

	


	
}