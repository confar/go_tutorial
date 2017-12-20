package main

import ("fmt"
        "net/http"
        "html/template"
	"io/ioutil"
	"encoding/xml"
	"sync"
)

var wg sync.WaitGroup

type NewsAggPage struct {
	Title string
	News map[string]NewsMap
}

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword string
	Location string
}


func index_handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w,`<h1>Whoa, mate this stuff is neat</h1>
					<p>Whoa, go is something</p>
					<p>For sure my mate</p>
`) //for multiple line printing we use string in ``


	//fmt.Fprintf(w, "<h1>Whoa, mate thi	s stuff is neat</h1>")
	//fmt.Fprintf(w, "<p>Whoa, go is something</p>")
	//fmt.Fprintf(w, "<p>For sure my mate</p>")
	fmt.Fprintf(w, "<p>You %s even add %s </p>", "can", "<strong>variables</strong>")
}

func newsRoutine(c chan News, Location string) {
	defer wg.Done()
	var n News
	resp, _ := http.Get(Location)
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &n)
	resp.Body.Close()
	c <- n

}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	var s SitemapIndex

	news_map := make(map[string]NewsMap)
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	//string_body := string(bytes)
	//fmt.Println(string_body)
	xml.Unmarshal(bytes, &s)
	resp.Body.Close()
	queue := make(chan News, 30)
	//fmt.Println(s.Locations)
	for _, Location := range s.Locations {
		wg.Add(1)
		go newsRoutine(queue, Location)
	}
		wg.Wait()
		close(queue)
		for elem := range queue {
			for idx, _ := range elem.Keywords{
				news_map[elem.Titles[idx]] = NewsMap{elem.Keywords[idx], elem.Locations[idx]}
			}
		}



	p := NewsAggPage{Title: "This great news aggregator", News: news_map}
	t, _ := template.ParseFiles("basictemplating.html")
	fmt.Println(t.Execute(w, p))
	//fmt.Fprintf(w, "Whoa, mate this stuff is neat and this is about")
}

func main () {
    http.HandleFunc("/", index_handler)
    http.HandleFunc("/agg", newsAggHandler)
    http.ListenAndServe(":8000", nil)

}
