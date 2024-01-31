package concurency

import (
	"fmt"
	"math"
	"sync"
	"time"
)

//Normal bir programlam dilinde geliştirme yapılırken. Multi thread bir uygualama geliştirmek istersek
//Bu cpu üzerindeki schedular ile threadlarin oluşturulması yönetilmesi ve öldürülmesi şeklinde gerçekleşir.
//Bu işlem diğer programlama dillerinde cpu seviyesindeyken go'da programlama dili seviyesinde gorutine'ler
//ile gerçekleşir. Böylece çok daha hızlı ve az alan maliyeti ile tamamlanır.

//****************************

//Gorutine oluşturabilmek için go keywordun den yararlanılır.
// Bu şekilde işlemler concurent olarak çalışmaya başlar.
//Multithread değil. Single thread asenkron gibi düşünebilirz.
//Paralel aynı anda iki işlem yapmaz bir işlem uzun sürünce diğerinden devam eder.
//Ancak bunları bir sıraya koyabilmek için waitgroup package'ından yararlanılır.
//Eğer bi şekilde sıraya konmaz ise main gorutine diğer gorutineler çalıştımı çalışmadımı bakmadn uygulamayı sonlandırır.
var wg sync.WaitGroup

func main() {

	//Bu şekilde şunu demiş oluruz.Main gorutine senin bekliyeceğin 1 tane gorutine var.
	//wg.Add(1)

	//go PrintX()
	//fmt.Println()
	//wg.Wait() //wg Add metodunda parametre olarak geçen sayı kadar Gorutine tamamlanana kadar bekle
	//PrintY()

	//***

	//ExampleForLoop()

	//***

	//ExampleForLoop2()

	//***

	//raceCondition()

	//***

	//raceConditionFixed()

	//***

	raceConditionWithPointer()
}

func PrintX() {
	for i := 0; i < 20; i++ {
		fmt.Print("X")
	}

	//Tamamlandı bilgisini gönderir
	wg.Done()
}

func PrintY() {
	for i := 0; i < 20; i++ {
		fmt.Print("Y")
	}
}

func ExampleForLoop() {
	//Burdaki örnekte ekrana 10 adet 10 basar.Şu şekilde çalışır.
	//for döngüsü dönmeye başlar gorutineler schedule edilir.
	//Schedule edilmiş gorutineler çalışmaya başladığı anda o anki i değeri ne ise onu basar
	// 10'a kadar dönen bir döngüde genelde en son 10 çıktısı daha çok gözükür.Hepsi 10 da olablir
	//arada farklı rakamlarda olabilir.Gorutine çalıştığında o andaki i değeri ne ise onu basar.

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(time.Second * 3)
}

func ExampleForLoop2() {
	//Burdada yukarıdakine benzer bir mantık vardır. Ama burası asenkron olmasına rağmen
	//sub functiona parametre olarak değer geçtiğimiz için bütün paramere geçilen rakamları basar
	//Ancak sırası farklı olur 0 7 3 5 9. şeklinde.

	for i := 0; i < 10; i++ {
		go func(value int) {
			fmt.Println(value)
		}(i)
	}

	time.Sleep(time.Second * 3)
}

//*** Race Condition ***

//Race condition: Farklı threadlarin memory üzerindeki aynı objeye erişip manipule etmeye çalışmasıdır.

func raceCondition() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	//shared value
	val := 0

	go func() {
		for i := 0; i < 1000000; i++ {
			val++
		}

		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000000; i++ {
			val++
		}

		wg.Done()
	}()

	wg.Wait()

	fmt.Println(val)
}

func raceConditionFixed() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	//shared value
	val := 0

	//Bildiğmiz c# daki lock mantığı.Lock'ladığımız objeye birden fazla thread aynı anda manipulasyon işlemi yapamaz.
	lock := sync.Mutex{}

	go func() {
		for i := 0; i < 1000000; i++ {
			lock.Lock()
			val++
			lock.Unlock()
		}

		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000000; i++ {
			lock.Lock()
			val++
			lock.Unlock()
		}

		wg.Done()
	}()

	wg.Wait()

	fmt.Println(val)
}

// *** RaceConditionWithPointer

type RaceTest struct {
	Val int
}

func raceConditionWithPointer() {
	raceTest := &RaceTest{}

	wg := &sync.WaitGroup{}
	wg.Add(1000)

	lock := &sync.Mutex{}

	for i := 0; i < 1000; i++ {
		go increment(raceTest, wg, lock)
	}

	wg.Wait()

	fmt.Println(raceTest.Val)
}

func increment(rt *RaceTest, wg *sync.WaitGroup, lock *sync.Mutex) {
	lock.Lock()
	rt.Val++
	wg.Done()

	lock.Unlock()
}

// *** /RaceConditionWithPointer
// *** /Race Condition ***

//*** Channel ***

type circle struct {
	r float64
}

func (c circle) display() {
	fmt.Println("A Circle")
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c circle) areatwo(chan1 chan string) {
	chan1 <- "mertcan"
}

func exec() {
	c1 := circle{5}

	//Böyle bi tanımlamada hata verir nedeni ise bi alt satırda
	//burdan dönen response göre bir işlem var bundan dolayı
	//bu gibi durumlarda channel kullanılması daha doğrudur.
	area1 :=go c1.area()
	fmt.Println("%2f\n", area1)


	var testChannel chan string
	testChannel = make(chan string)

	go c1.areatwo(testChannel)


	fmt.Println(testChannel)

	c1.display()
}

//*** /Channel ***

