package filelisting

import (
	"net/http"
	"os"
	"io/ioutil"
	"strings"
	"fmt"
)

const prefix  = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HandlerFileList(writer http.ResponseWriter, request *http.Request) error{
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError(
			fmt.Sprintf("path %s must start with %s", request.URL.Path, prefix))
	}

	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		//panic(err)
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}
