package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Person struct {
	Name string
	Age int
	Job string
}

func(p Person) Teach1() {
	fmt.Println("Teach1:", p.Name, p.Age, p.Job)	
}

func(p Person) Teach2(c, s string) {
	fmt.Println("Teach2:", p.Name, p.Age, p.Job, c, s)
}

func(p Person) Teach3(c, s string) (info string) {
	return p.Name + " - " + strconv.Itoa(p.Age) + " - " + p.Job + " - " + c + " - " + s
}

func main() {
	p := Person{Name: "王二妮", Age: 19, Job: "Teacher"}

	value := reflect.ValueOf(p)
	fmt.Printf("%s %s\n", value.Kind(), value.Type())

	f1 := value.MethodByName("Teach1")
	fmt.Println(f1.Kind(), f1.Type())
	f1.Call(nil)

	f2 := value.MethodByName("Teach2")
	fmt.Println(f2.Kind(), f2.Type())
	args2 := []reflect.Value{reflect.ValueOf("数学"), reflect.ValueOf("上海市第八中学")}
	f2.Call(args2)

	f3 := value.MethodByName("Teach3")
	fmt.Println(f3.Kind(), f3.Type())
	args3 := []reflect.Value{reflect.ValueOf("英语"), reflect.ValueOf("上海市第一中学")}
	v3 := f3.Call(args3)
	fmt.Println(v3, v3[0])
	fmt.Printf("%T %T\n", v3, v3[0])
}
