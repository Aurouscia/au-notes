package main

import "fmt"

// ChangeValue 尝试将传入的整数改为 100
func ChangeValue(n *int) {
	*n = 100
}

// Person 结构体
type Person struct {
	Name string
	Age  int
}

// SetPerson 设置 Person 的字段
func SetPerson(p **Person, name string, age int) {
	*p = new(Person)
	(**p).Name = name
	(**p).Age = age
}

func main() {
	// 问题1：ChangeValue 无法修改 num
	num := 10
	ChangeValue(&num)
	fmt.Println("修改后 num =", num)

	// 问题2：空指针问题
	var person *Person
	SetPerson(&person, "Bob", 25)
	fmt.Println("Person:", *person)
}
