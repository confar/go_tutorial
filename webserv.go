package main

import ("fmt"
        "net/http")

func index_handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w,`<h1>Whoa, mate this stuff is neat</h1>
					<p>Whoa, go is something</p>
					<p>For sure my mate</p>
`) //for multiple line printing we use string in ``


	//fmt.Fprintf(w, "<h1>Whoa, mate this stuff is neat</h1>")
	//fmt.Fprintf(w, "<p>Whoa, go is something</p>")
	//fmt.Fprintf(w, "<p>For sure my mate</p>")
	fmt.Fprintf(w, "<p>You %s even add %s </p>", "can", "<strong>variables</strong>")
}
 //ой братишка
func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, mate this stuff is neat and this is about")
}

func main () {
    http.HandleFunc("/", index_handler)
    http.HandleFunc("/about", about_handler)
    http.ListenAndServe(":8000", nil)

}
