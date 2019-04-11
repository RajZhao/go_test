package main
import (
	"fmt"
	"strings"
)
func processUrl(url string)string{
	result := strings.HasPrefix(url,"http://")
	if !result{
		url = "http://" + url
	}
	return url

}

func processPath(path string)string{
	result := strings.HasSuffix(path,"/")
	if !result{
		path += "/"
	}
	return path
}

func main(){
    var (
    	url string
    	path string
	)
    fmt.Scanf("%s%s",&url,&path)
    fmt.Println(processUrl(url))
    fmt.Println(processPath(path))
}
