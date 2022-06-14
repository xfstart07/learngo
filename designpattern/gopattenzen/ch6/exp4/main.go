package main

import "fmt"

type IBook interface {
	getName() string
	getPrice() int
	getAuthor() string
}

type NovelBook struct {
	name   string
	price  int
	author string
}

type OffNevelBook struct {
	NovelBook
}

func New(name, author string, price int) *OffNevelBook {
	return &OffNevelBook{NovelBook{name: name, author: author, price: price}}
}

func (n *NovelBook) getName() string {
	return n.name
}

func (n *NovelBook) getAuthor() string {
	return n.author
}

func (n *NovelBook) getPrice() int {
	return n.price
}

func (o *OffNevelBook) getPrice() int {
	selfPrice := o.NovelBook.getPrice()

	if selfPrice > 100 {
		return (int)(float64(selfPrice) * 0.9)
	} else {
		return (int)(float64(selfPrice) * 0.8)
	}
}

func main() {
	books := make([]IBook, 0, 4)

	var book IBook
	book = New("设计模式之禅", "秦小波", 100)
	// ibook := (IBook)(book)
	books = append(books, book)

	book = New("设计模式 Head First", "Head First", 111)
	books = append(books, book)

	book = New("Go语言实战", "李兆海", 100)
	books = append(books, book)
	fmt.Println(books)
	for idx := range books {
		book := books[idx]
		fmt.Printf("书籍名称：%s, 书籍作者: %s, 书籍价格: %d\n", book.getName(), book.getAuthor(), book.getPrice())
	}

}
