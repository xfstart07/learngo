// Author: xufei
// Date: 2019/1/18

package http_api

import (
	"log"
	"net/http"
	"time"
	"io"
)

type Decorator func(APIHandler) APIHandler

type APIHandler func(http.ResponseWriter, *http.Request) (interface{}, error)

func Decorate(f APIHandler, ds ...Decorator) http.HandlerFunc {
	decorated := f

	// 将装饰器一一装饰在 f 上，就像穿衣服，一件一件套上去
	log.Println("Decorator number", len(ds))
	for _, decorate := range ds {
		decorated = decorate(decorated)
	}

	return func(writer http.ResponseWriter, req *http.Request) {
		decorated(writer, req)
	}
}

func Logf(l *log.Logger) Decorator {
	return func(f APIHandler) APIHandler {
		return func(w http.ResponseWriter, req *http.Request) (interface{}, error) {
			start := time.Now()
			resp, err := f(w, req)
			elapsed := time.Since(start)

			l.Printf("时间= %v, resp = %v，error = %v", elapsed, resp, err)

			io.WriteString(w, resp.(string))
			return resp, err
		}
	}
}

func V1(f APIHandler) APIHandler {
	return func(w http.ResponseWriter, req *http.Request) (interface{}, error) {
		resp, err := f(w, req)

		log.Println("write api v1")
		return resp, err
	}
}
