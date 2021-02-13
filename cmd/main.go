package main

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, middleware!")
}

func middleware1(nextFunc http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[START] middleware1")
		nextFunc.ServeHTTP(w, r)
		fmt.Println("[END] middleware1")
	}
}

func middleware2(nextFunc http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[START] middleware2")
		nextFunc.ServeHTTP(w, r)
		fmt.Println("[END] middleware2")
	}
}

func middleware3(nextFunc http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[START] middleware3")
		nextFunc.ServeHTTP(w, r)
		fmt.Println("[END] middleware3")
	}
}

// ミドルウェア型（1つ1つのミドルウェア）
type middleware func(http.HandlerFunc) func(w http.ResponseWriter, r *http.Request)

// ミドルウェアスタック（ミドルウェアをまとめたもの）
type mwStack struct {
	middlewares []middleware
}

// ミドルウェアスタックの初期化
func newMws(mws ...middleware) mwStack {
	// 空のスライスに追加して構造体に格納して返す
	return mwStack{append([]middleware(nil), mws...)}
}

// ミドルウェアを実装
func (m mwStack) then(h http.HandlerFunc) http.HandlerFunc {
	for i := range m.middlewares {
		h = m.middlewares[len(m.middlewares)-1-i](h)
		// ループ1回目「middleware3(helloHandler)」
		// ループ2回目「middleware2(middleware3(helloHandler))」
		// ループ3回目「middleware1(middleware2(middleware3((helloHandler)))」
	}
	// hは「middleware1(middleware2(middleware3(helloHandler)))」
	return h
}

var log = logrus.New()

func main() {
	// ミドルウェアをまとめる（初期化処理）
	middlewares := newMws(middleware1, middleware2, middleware3)

	mux := http.NewServeMux()

	// 第2引数は「func(ResponseWriter, *Request)」
	// ミドルウェアを実装
	mux.HandleFunc("/hello", middlewares.then(helloHandler))

	log.Formatter = new(logrus.JSONFormatter) // JSONで出力
	log.SetLevel(logrus.WarnLevel)            // ログレベルの設定

	/* logrusのログレベルは7つ
	/*
	   panic
	   fatal
	   error
	   warn
	   info
	   debug
	   trace
	*/
	/* セットしたレベル以上のものが出力 */

	log.WithFields(logrus.Fields{ // 出力されない
		"number": 1,
		"size":   10,
	}).Info("this is info level")

	log.WithFields(logrus.Fields{ // 出力される
		"number": 1,
		"size":   10,
	}).Error("this is error level")

	http.ListenAndServe(":8080", mux)
}
