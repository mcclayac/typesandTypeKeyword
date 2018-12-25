package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type webPage struct {
	url string
	body []byte
	err error
}



func (w *webPage) get() {
	resp, err := http.Get(w.url)
	if err != nil {
		w.err = err
		return
	}
	defer resp.Body.Close()

	w.body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		w.err = err
	}
}
/*
func (w *webPage) isOK() bool {
	return w.err  == nil

}
*/

func (w webPage) isOK() bool {
	return w.err  == nil

}

func main() {

	w1 := &webPage{url: "http://oreilly.com/"} // explicit
											// create a NEW Structure


	w2 :=  &webPage{}
	w2.url = "https://www.bbc.com/"

	w3 := new(webPage)
	//w3.url = "https://www.bbc.com/"
	w3.url = "https://www.cnn.com/"


	w1.get()
	w2.get()
	w3.get()
	if w1.isOK() && w2.isOK() && w3.isOK() {
		fmt.Printf("URL: %s  Error: %s  Length: %d\n\n", w1.url, w1.err, len(w1.body))
		fmt.Printf("URL: %s  Error: %s  Length: %d\n\n", w2.url, w2.err, len(w2.body))
		fmt.Printf("URL: %s  Error: %s  Length: %d\n\n", w3.url, w3.err, len(w3.body))

	} else {
		fmt.Printf("Something went wrong")
	}
}

/*
func Get(url string) (resp *Response, err error)
Get issues a GET to the specified URL. If the response is one of the
following redirect codes, Get follows the redirect, up to a maximum of
10 redirects:

301 (Moved Permanently)
302 (Found)
303 (See Other)
307 (Temporary Redirect)
308 (Permanent Redirect)

An error is returned if there were too many redirects or if there was an
HTTP protocol error. A non-2xx response doesn't cause an error. Any
returned error will be of type *url.Error. The url.Error value's Timeout
method will report true if request timed out or was canceled.

When err is nil, resp always contains a non-nil resp.Body. Caller should
close resp.Body when done reading from it.

Get is a wrapper around DefaultClient.Get.

To make a request with custom headers, use NewRequest and
DefaultClient.Do.
*/