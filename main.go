package main

import(
	"fmt"
	//"http/net"
	"gitihub.com/gorilla/mux"
)


func main()  {

	fmt.Println("hi")
	fmt.Println("go!")

	router := mux.NewRouter()

    router.HandleFunc("/test", test)

    http.ListenAndServe(":5000", router)
}

func test(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("this is a test"))
    //json.NewEncoder(w).Encode(struct)
}
