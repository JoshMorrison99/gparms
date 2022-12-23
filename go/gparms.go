package main

import (
	"fmt"
	"bufio"
	"os"
	"net/url"
	"flag"
)

func main(){

	isSplit := flag.Bool("s", false, "Split the query parameters")
	flag.Parse()

	urls := bufio.NewScanner(os.Stdin)

	for urls.Scan(){
		urlString, err := url.Parse(urls.Text())

		if err == nil {

			if urlString.RawQuery != "" {
				if *isSplit {
					query, err := url.ParseQuery(urlString.RawQuery)

					if err == nil {
						for i := range(query){
							fmt.Println(urlString.Scheme + "://" + urlString.Host + urlString.Path + "?" + i + "=")
						}
					}
				}else{
					fmt.Println(urlString)
				}
			}
		}
	}
}
