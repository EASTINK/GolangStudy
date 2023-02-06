package main

import "fmt"

type ReaderIf interface {
	ReadBook()
}

type WriterIf interface {
	WriteBook()
}

type Book struct {
	id string
}

func (this *Book) ReadBook() {
	fmt.Println(this.id, ":Read a Book")
}

func (this *Book) WriteBook() {
	fmt.Println(this.id, ":Write a Book")
}

func main() {
	b := &Book{"b"}
	/*
		当b => Book{}
		那么b 就拥有了两个func ReadBook 、 WriteBook 和 属性id
		比如：
	*/
	b.ReadBook()
	b.WriteBook()

	/*
		尝试定义一个interface ReaderIf 的 r
		ReaderIf{
			ReadBook()
		}
		当 r => b 的时候 需要b提供func ReadBook给 r  此时他们之间的关系是这样的
		r{
			b[面向r]{
				private id string
				public ReadBook()
				private WriteBook()
			}
		}
		比如：
	*/
	var r ReaderIf
	r = b
	r.ReadBook() //b :Read a Book

	/*
		当我们尝试再定义一个 interface WriteIf 给 w
		WriteIf{
			WriteBook()
		}
		此时他们之间的关系是这样的
		w{
			b[面向w]{
				private  id string
				private ReadBook()
				public WriteBook()
			}
		}
		//比如：
		var w WriterIf = b
		w.WriteBook() //b :Write a Book
	*/

	/*
		假如不为w提供b,而是用一个断言将 ReadIf的 r 强制转换为  WriteIf 也可行吗，为什么？
	*/

	var w WriterIf
	/*
		因为 定义 w为 interface WriterIf 是一个动态类型
		w 此时的内部结构为<pair type: 空 value: 空/>
	*/
	w = r.(WriterIf)
	/*
		r 此时的内部结构为<pair type：Book value: Book{"b"}>
		r.(WriterIf) 以后 WriterIf实现的type也依旧是Book
		因为
			r = &b,
			b = Book{"b"}  内部实现了Book{"b"}.WriteBook()
			当r（即ReaderIf的实现)强制转换到WriterIf的时候
			w(WriterIf的实现）继承的内部结构需要和r的pair是一致的
			<r.(WriterIf) pair type:Book value: Book{"b"}>

	*/
	w.WriteBook() //b :Write a Book
}
