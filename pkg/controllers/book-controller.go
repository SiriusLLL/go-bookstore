package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SiriusLLL/go-bookstore/pkg/models"
	"github.com/SiriusLLL/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	// 获取所有图书的信息
	newBooks := models.GetAllBooks()
	// 将newBooks中的图书信息转换为JSON格式
	res, _ := json.Marshal(newBooks)
	// 设置响应头的内容类型
	w.Header().Set("Content-Type", "pkglication/json")
	// 设置响应状态码为200，表示成功
	w.WriteHeader(http.StatusOK)
	// 将JSON数据发送回客户端
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	// 从请求中获取URL路径中的参数，map[string]string
	vars := mux.Vars(r)
	bookId := vars["bookid"]
	// string转int64
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing", err)
	}
	// 从数据库中获取指定ID的图书信息
	bookDetails, _ := models.GetBookById(ID)
	// 图书信息转换为JSON格式
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	// 创建一个空结构体，用于存储从请求中解析的图书数据
	CreateBook := &models.Book{}
	// 将请求体中的JSON数据解析并存储到CreateBook中
	utils.ParseBody(r, CreateBook)
	// 调用结构体方法，返回创建的图书对象
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookid"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// 创建一个空结构体，用于存储从请求中解析的图书数据
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookid"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	// 根据ID获取数据库中的图书信息
	bookDetails, db := models.GetBookById(ID)
	// 更新数据库中的图书信息
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	// 将更新结果保存到数据库中
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
