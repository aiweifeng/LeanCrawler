package fetch

import (
	"net/http"
	"fmt"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"log"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"io/ioutil"
	"time"
)

var rateLimiter = time.Tick(100*time.Millisecond)

func Fetch(url string) ([]byte,error)  {
     <-rateLimiter
	/*resp,err := http.Get(url)
	if err!=nil{
		panic(err)
	}

	defer resp.Body.Close()*/

	client := &http.Client{}
	req,err := http.NewRequest("GET",url,nil)
	if err != nil{
		panic(err)
	}

	req.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.62 Safari/537.36")
    resp,err := client.Do(req)

	if resp.StatusCode !=http.StatusOK{
		return nil,fmt.Errorf("wrong status code:%d",resp.StatusCode)
	}

	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())
	return  ioutil.ReadAll(utf8Reader)

}

/**
自动获取编码
 */

 func determineEncoding(r io.Reader) encoding.Encoding{
 	bytes,err:= bufio.NewReader(r).Peek(1024)
 	if err!=nil{
 		log.Printf("fetcher err %v",err)
 		return unicode.UTF8
	}

	e,_,_ := charset.DetermineEncoding(bytes,"")
	return  e

 }
