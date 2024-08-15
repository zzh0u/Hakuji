package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"blockchain/block"
	"github.com/gorilla/mux"
)

var BlockChain *block.Blockchain

// Handler将区块链作为json字符串写回浏览器
func getBlockchain(w http.ResponseWriter, r *http.Request) {
	jbytes, err := json.MarshalIndent(BlockChain.Blocks, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	io.WriteString(w, string(jbytes))
}

// Handler根据发送的信息添加一个新区块
func writeBlock(w http.ResponseWriter, r *http.Request) {
	var checkoutItem block.BookCheckout
	if err := json.NewDecoder(r.Body).Decode(&checkoutItem); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not write Block: %v", err)
		return
	}
	BlockChain.AddBlock(checkoutItem)
	resp, err := json.MarshalIndent(checkoutItem, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not marshal payload: %v", err)
		w.Write([]byte("could not write block"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// 创建新的Book数据
func newBook(w http.ResponseWriter, r *http.Request) {
	var book block.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not create new book")
		return
	}

	// 使用生成ID作为区块添加
	// 创建ID 连接(concatenating) ISDB和发售日
	h := md5.New()
	io.WriteString(h, book.ISBN+book.PublishDate)
	book.ID = fmt.Sprintf("%x", h.Sum(nil))

	// send back payload
	resp, err := json.MarshalIndent(book, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not marshal payload:%v", err)
		w.Write([]byte("could not save book date"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func main() {
	BlockChain = block.NewBlockchain()

	r := mux.NewRouter()
	r.HandleFunc("/", getBlockchain).Methods("GET")
	r.HandleFunc("/", writeBlock).Methods("POST")
	r.HandleFunc("/new", newBook).Methods("POST")

	go func() {
		for _, block := range BlockChain.Blocks {
			if block == nil {
				fmt.Println("Error: Encountered a nil block in the chain.")
				continue
			}
			fmt.Printf("Prev.hash:%x\n", block.PrevHash)
			bytes, _ := json.MarshalIndent(block.Data, "", " ")
			fmt.Printf("Data:%v\n", string(bytes))
			fmt.Printf("Hash:%v\n", block.Hash)
			fmt.Println()
		}
	}()

	log.Println("Listening on port 3000")

	log.Fatal(http.ListenAndServe(":3000", r))

}
