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
	books := make([]*NovelBook, 0, 3)

	books = append(books, New("设计模式之禅", "秦小波", 100))
	books = append(books, New("设计模式 Head First", "Head First", 100))
	books = append(books, New("Go语言实战", "李兆海", 100))

	fmt.Println(books)
	for idx := range books {
		fmt.Printf("书籍名称：%s, 书籍作者: %s, 书籍价格: %d\n", books[idx].getName(), books[idx].getAuthor(), books[idx].getPrice())
	}

}
