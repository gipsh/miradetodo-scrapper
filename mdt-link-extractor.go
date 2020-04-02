package main

import (
    "fmt"
    "log"
    "github.com/PuerkitoBio/goquery"
    "strings"
    "net/url"
    "encoding/base64"
    "os"
)


func fixString(s string) string {

	out := []byte(s)
	out[0] = (out[0] - 1)

	return string(out)
}

func extractLinks(url string) {

    doc, err := goquery.NewDocument(url)

    if err != nil {
        log.Fatal(err)
    }

    doc.Find("a[href]").Each(func(index int, item *goquery.Selection) {
        href, _ := item.Attr("href")
	 decodedId := extractId(href)
	 fmt.Printf("\t %d. %s\n", index, decodedId)
    })
}

func extractId(movieUrl string) string {

    u, err := url.Parse(movieUrl)
    if err != nil {
        panic(err)
    }

    m, _ := url.ParseQuery(u.RawQuery)
    return decodeId(m["id"][0])

}


func decodeId(id string) string {

	ss := strings.Split(id, "!Ar")

	if len(ss) == 1 {
	  dcd, _ := base64.StdEncoding.DecodeString(id)
          return  string(dcd)

	}

	for i := 1; i < len(ss); i++ {
		ss[i] = fixString(ss[i])
	}

	finalout := strings.Join(ss[:], "")

	dd, _ := base64.StdEncoding.DecodeString(string(finalout))

	return string(dd)

}

func main() {


  //  doc, err := goquery.NewDocument("https://miradetodo.co/fangio-el-hombre-que-domaba-las-maquinas-2020-720p-hd/")
  doc, err := goquery.NewDocument(os.Args[1])



    if err != nil {
        log.Fatal(err)
    }

    doc.Find("iframe[data-lazy-src]").Each(func(index int, item *goquery.Selection) {
        href, _ := item.Attr("data-lazy-src")
	if  strings.Contains(href, "generator") {
		fmt.Printf("[*] link: [%s] \n", href)
		extractLinks(href)
	}

    })

}
