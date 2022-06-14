package main

import "fmt"

type IBook interface {
	getName() string
	getPrice() int
	getAuthor() string
}

type IComputerBook interface {
	IBook
	getScope() string
}

type NovelBook struct {
	name   string
	price  int
	author string
}

func New(name, author string, price int) *NovelBook {
	return &NovelBook{name: name, author: author, price: price}
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

func main() {
	books := make([]IBook, 0, 4)

	var book IBook
	book = New("设计模式之禅", "秦小波", 100)
	// ibook := (IBook)(book)
	books = append(books, book)

	book = New("设计模式 Head First", "Head First", 100)
	books = append(books, book)

	book = New("Go语言实战", "李兆海", 100)
	books = append(books, book)

	book = NewComputer("HTTP权威指南", "阿里前端", "1.0", 250)
	books = append(books, book)

	fmt.Println(books)
	for idx := range books {
		book := books[idx]
		fmt.Printf("书籍名称：%s, 书籍作者: %s, 书籍价格: %d\n", book.getName(), book.getAuthor(), book.getPrice())
	}

}

type ComputerBook struct {
	name   string
	price  int
	author string
	scope  string
}

func NewComputer(name, author, scope string, price int) *ComputerBook {
	return &ComputerBook{name: name, author: author, price: price, scope: scope}
}

func (n *ComputerBook) getName() string {
	return n.name
}

func (n *ComputerBook) getAuthor() string {
	return n.author
}

func (n *ComputerBook) getPrice() int {
	return n.price
}

func (n *ComputerBook) getScope() string {
	return n.scope
}
