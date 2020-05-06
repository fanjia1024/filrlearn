package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func HttpGet(url string) (result string,err error){
	resp,err:=http.Get(url)

	defer resp.Body.Close()

	buf :=make([]byte,4096)

	for  {
		n,err1:=resp.Body.Read(buf)
		if n == 0{
			break
		}
		if err1 !=nil && err1!=io.EOF {
			break
		}

		result+=string(buf[:n])

	}
	return result,err

}

func working(start,end int)  {


	fmt.Printf("爬取第%d到%d\n",start,end)

	for i:=start;i<=end;i++ {
		url :="https://tieba.baidu.com/f?kw=lol&ie=utf-8&pn="+strconv.Itoa((i-1)*50)
		resp,err:=HttpGet(url)
		if err != nil {
			fmt.Println("调用出错")
			continue
		}

		//fmt.Println(resp)

		f,err :=os.Create("tieba"+strconv.Itoa(i)+".html")
		f.WriteString(resp)
		f.Close()
	}
}



func main()  {

	start :=1
	end :=2
	working(start,end)

}



