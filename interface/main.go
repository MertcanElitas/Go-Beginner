package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	a, b float64
}

func (r rectangle) area() float64 {
	return r.a * r.b
}

func (r rectangle) circumference() float64 {
	return 2 * (r.a + r.b)
}

type circle struct {
	r float64
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c circle) circumference() float64 {
	return 2 * math.Pi * c.r
}

type shape interface {
	area() float64
	circumference() float64
}

func interfacePrint(sha shape) {
	fmt.Println(sha)
	fmt.Println(sha.area())
	fmt.Println(sha.circumference())
	fmt.Printf("%T", sha)
	fmt.Println()
}

//Kendi oluşturduğumuz type içinde, interface de tanımlı olan bütün metotlar aynı isimle ve geri dönüş tipiyle
//impletemente edildiyse sorunsuz şekilde kullanılabilir.
func main() {
	r1 := rectangle{3, 8}
	fmt.Println("Area", r1.area())
	fmt.Println("Circumference", r1.circumference())

	c1 := circle{4}

	fmt.Println()
	interfacePrint(r1)
	interfacePrint(c1)

	//*********//

	//go anahtar keywordu void olan methodlarda bir gorutine oluştururken void olmayan metotlarda bunu yapamaz.
	//nedeni ise go anahtar keywordun'den sonra kod bloğu akmaya devam eder. Devamında return edilecek olan değere
	//ihtiyacmımız olduğu için ve o an o değer elimzde olmadığı için bu bir hatadır.
	//Bu durumlarda yani geri dönüş yapan metorlarda go keywordünü kullanıp bir gorutine oluşturmak isteniyosa
	//channel lardan yararlanılır.
	c2 := circleTwo{5}
	go c2.display()
	area1 := c2.area()
	fmt.Printf("%T %v", area1, area1)

	// var myChannel chan string
	// myChannel = make(chan string)

	myChannel := make(chan string)

	go merhaba(myChannel)
	fmt.Println(<-myChannel)
}

//************************Channel*********************//

type circleTwo struct {
	r float64
}

func (c circleTwo) display() {
	fmt.Println("A circle")
}

func (c circleTwo) area() float64 {
	return math.Pi * c.r * c.r
}

func merhaba(chan1 chan string) {
	chan1 <- "merhaba"
}
