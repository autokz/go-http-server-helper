package handler

import (
	"io/ioutil"
	"log"
	"net/http"
)

type TestHandler struct {
}

func Test(resp http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(body))

	resp.WriteHeader(200)
	resp.Write([]byte("\"message\":\"ok\""))
}
