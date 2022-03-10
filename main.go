//package clause
package main

//Import statement
import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/jinzhu/gorm"
)

//Package düzeyündeki variable'lar uygulama ayakta olduğu sürece memory üzerinde kalırlar.
//Function scpor variable'lar ise function scope kapandığı anda dispose olurlar.
var testPackageVar = "Package"

//Code
func main() {
	Mertcan()
	printProduct()
	//pointerPassReferance()
	//pointerExample3()
	//pointerExample()
	//definedType()
	//inheritaceExample()
	//structWithTypeGlobal()
	//structWithType()
	//structExample()
	//mapExample()
	//sliceExampleRemove()
	//sliceExampleAppend()
	//switchExample()
	//Constants()
	//testPackageVar = "Package 2"
	//inrementDecrementOperations()
	//typeConversion()
	//variableDeclartions()

	//fmt.Println(testPackageVar)

	//test()
}

func variableDeclartions() {

	//Variable declaration 1
	var name string
	name = "Mertcan"

	var age int
	age = 26

	var isMarried bool
	isMarried = true

	var moment time.Time
	moment = time.Now()

	//******************************//

	//Variable declaration 2
	var lastname string = "Elitaş"
	var identityNumber int = 123
	var isMale bool = true

	//******************************//

	//Variable declaration 4
	city := "İstanbul"
	totalPayment := 500
	isYoung := false

	//******************************//

	//Variable declaration 5
	var (
		name1      = "mertcan"
		age1       = 26
		isMarried1 = true

		weight = 81
		height = 178
	)

	//******************************//

	//Variable declaration 6
	var name2, age2, isMarried2, weight2, height2 = "Mertcan", 26, true, 75, 178

	//******************************//

	//Variable declartion 7
	name3, age3, isMarried3, weight3, height3 := "Mertcan", 26, true, 75, 178

	//Default Values (Zero Values)
	// string => "" int => 0 bool => false
	var name4 string
	var age4 int
	var isMarried4 bool

	var name5 = "Mertcan"
	//name5 := "Mertcan" //Bu bir error oluşturur. Nedeni ise name5 adında 2 adet variable tanımladığımız zaman exception oluşur.
	name5, name6 := "Gamze", "Mine" //Burda hata vermez çünkü := operatörü ile burda nam6 için yeni bir declaration yapar name5 için sadece assignment yapar.

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)
	fmt.Println(moment)

	fmt.Println(lastname)
	fmt.Println(identityNumber)
	fmt.Println(isMale)

	fmt.Println(city)
	fmt.Println(totalPayment)
	fmt.Println(isYoung)

	fmt.Println(name1)
	fmt.Println(age1)
	fmt.Println(isMarried1)
	fmt.Println(weight)
	fmt.Println(height)

	fmt.Println(name2)
	fmt.Println(age2)
	fmt.Println(isMarried2)
	fmt.Println(weight2)
	fmt.Println(height2)

	fmt.Println(name3)
	fmt.Println(age3)
	fmt.Println(isMarried3)
	fmt.Println(weight3)
	fmt.Println(height3)

	fmt.Println(name4)
	fmt.Println(age4)
	fmt.Println(isMarried4)

	fmt.Println(name5)
	fmt.Println(name6)
}

func typeConversion() {
	//Type Conversion 1
	x := 10
	y := 10.0

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	//Type conversion => int(y) => 10.0 => 10
	total := x + int(y)

	fmt.Println(total)

	//Burada y nin tip tekrar float64 olarak ekrana basar.Burdan çıkan sonuç
	//işlem gerçekleşirken y değeri int'e convert edilir ancak daha sonra eski tipi olan
	//float64 den devam eder.
	fmt.Printf("%v %T\n", y, y)

	//******************************//

	//Type Conversion 2
	var a int8 = 10
	var b int16 = 10

	//İkiside int tipinde olmasına rağmen bu işlem'de type'ları aynı sayılmayarak gerçekleşmez.
	fmt.Println(a + int8(b))

	//******************************//

	//Type conversion 3
	var t int8 = 127
	var e int16

	//int16, int8'i her durumda cover etsede yinede manuel olarak type conversion yapılamlı yoksa hata verir.
	e = int16(t)

	fmt.Println(e)

	//******************************//

	//Type conversion 4
	k := 10
	p := "10"

	fmt.Printf("%v %T", k, k)
	fmt.Printf("%v %T", p, p)

	//string tipi bu şekilde type conversion ile int tipine dönüştürülemez.
	//fmt.Println(k + int(p))

	//******************************//

	//Type Conversion 5

	num1 := 106
	str1 := string(106)

	//int değeri 106 olan bir değişkeni stringe type conversion metodu ile dönüştürürsek. 106 nın asci tablosundaki karşılığı olan j değerine convert eder.
	fmt.Printf("%v %T\n", num1, num1)
	fmt.Printf("%v %T\n", str1, str1)
}

var db *gorm.DB

func inrementDecrementOperations() {
	x, y := 15, 10

	//Aşağıdaki işlemlerde bölme işlemine kadar sonuçlar tam sayı olduğu için type değşikeni de int olarak ekrana basılır.
	//Ancak bölme işlemin 15/10 1.5 olur böyle olmasına rağmen go dili int/int=int mantığında çalışarak çıkan sonucu int'e dönüştürür ve tipini int olarak basar.
	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)
	fmt.Printf("%v %T\n", (x + y), (x + y))
	fmt.Printf("%v %T\n", (x - y), (x - y))
	fmt.Printf("%v %T\n", (x * y), (x * y))

	fmt.Printf("%v %T\n", (x / y), (x / y))

	//********************************//

	//Aynı mantık ile int/int=int mantığı ile sonuç 2.5 de olsa 2 olarak int'e çevirir.
	z := 5 / 2

	fmt.Printf("%v %T\n", z, z)

	// *********************************//

	//Burda yine aynı mantık int/int=int ile sonucu 2 ye convert eder ancak tipi biz convert ettiğimiz için float64 olarak ekrana basar.
	v := float64(5 / 2)

	fmt.Printf("%v %T\n", v, v)

	//*******************************//

	//Bu örnekte ise float/int=float yarak sonucu 2.5 ve tipide float olarak set eder.
	k := 5.0 / 2

	fmt.Printf("%v %T\n", k, k)

	//********************************//

	d := 10

	fmt.Println(d)

	fmt.Println(d + 1)

	// go programlama dilinde aşağıdaki gibi bir operation hata verir.
	//Nedeni şudur: go programlama dilinde increment ve decrement birer statement'dır. PrintLn işlemide bir statementdır.
	//Go diline göre 1 satırda yalnızca bir statement olmalıdır.
	//Statement bir şeyin yapılmasını ifade eden en küçük birimdir. x=2 bir statement'dır
	//Expression ise sadece bir işlemdir.  5*5 bir assignment'dır
	//fmt.Println(d++)

	//Bu şekilde ayrı ayrı olmalıdır.
	d++

	fmt.Println(d)
}

func Constants() {

	//Go da constant'lar variable gibi isimlendirilir
	//Constant'lar => compile time da oluşturulur
	//Variable'lar => run time da oluşturulur
	const age int = 5
	const name = "Test"

	//****************************************//
	//Typless Constant

	const x int16 = 10
	const y float64 = 3.4

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	//Go dili aşağıdaki toplama işlemini yapmak istersek hata verir
	//Bunun nedeni int16 ve float64 olan iki farklı tipi birbiri ile toplayamaz.
	//Ancak değişkeenleri Örnek 2'deki gibi typeless tanımlarsak o zaman
	//dinamic olarak convert işlemini yapar ve toplama işlemini gerçekleştirir.
	//fmt.Printf("%v %T",x+y,x+y)

	//Örnek 2

	const a = 10
	const b = 3.4

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	//Burada görüldüğü üzere bir hata yoktur
	//İlk compie timeda a değişkenini int olarak belirler
	//Daha sonra float ile gerçekleştirilen bir operasyonda
	//floata otomatik convert edip dinamiclik katar.
	//Her zaman daha duyarlı olana göre convert işlemini yapar.
	fmt.Printf("%v %T\n", a+b, a+b)
}

func ifExample() {

	x := 27

	if x%2 == 0 {
		fmt.Println(x, "bir çift sayıdır")
	} else {
		fmt.Println(x, "bir tek sayıdır")
	}

	//2 statement ifade tek satırda yazılmak istenirse aralarına noktalı virgül konulur.
	//Ve bu örnekte a if scope'u süresince ramde kalır if bloğu bitince a değişkeni dispose olur
	if a := 10; a < 20 {
		fmt.Println(a, "a küçüktür 20 den")
	} else if a%2 == 0 {
		fmt.Println(a, "bir çift sayıdır")
	} else if a%2 != 0 {
		fmt.Println(a, "bir tek sayıdır")
	}
}

func switchExample() {

	//c# dakinden farklı olarak break anahtar sözcüğü kullanılmaz default olarak zaten var gibi davranır.
	switch grade := 3; grade {
	case 1:
		fmt.Println("Başarısız")
	case 2:
		fmt.Println("Eh işte")
	case 3:
		fmt.Println("Orta")
	case 4:
		fmt.Println("Başarılı")
		y := 4
		fmt.Println(y)
	case 5:
		fmt.Println("Çok Başarılı")
	}
}

func forLoopExample() {
	//for c# dakinin aynısı
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	//break ve continue aynı c# da kullanıldığı gibidir.

	//while döngüsüde aşağıdaki gibi kullanılabilir.
	a := 5

	for a < 10 {
		fmt.Print(a)
		a++
	}
}

func blankIdentifier() {

	//Blank identifier: Go dilinde her tanımlanan variable mutlaka bir yerde kullanılmalıdır.
	//Aksi halde compile timeda hata oluşur. Ancak bazı durumlarda kullanmıyacağımız değişkenler olabilir
	//işte o değişkenleri _ ifadesi ile tanımlarız.

	fmt.Print("Lütfen Sınav Sonucunu Giriniz")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n') //_ => blank identifier
	fmt.Println(value)
}

func coupleReturn(bölünen, bölen int) (int, int) {
	bölüm := bölünen / bölen
	kalan := bölünen % bölen

	return bölüm, kalan
}

func exceptionHandling(number int) (string, error) {
	if number < 0 {
		return "", errors.New("O dan küçük sayı girdiniz.")
	}

	return "Girdiğiniz sayı doğru", nil
}

func exceptionHandling2(number int) (float64, error) {
	if number%2 != 0 {
		return math.Sqrt(float64(number)), errors.New("Girdiğiniz sayı çift değil")
	}

	return 0.0, nil
}

func exceptionHandling3() {
	file, err := os.Open("test.txt")
	if err == nil {
		fmt.Println(err)
	} else {
		fmt.Println(file)
	}
}

func arrayExample() {
	cities := [2](string){"sivas", "istanbul"}

	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	for index, city := range cities {
		fmt.Println(index, city)
	}
}

func sliceExample() {
	//Array'ler go programlama dilinde value type'dır.
	//Slice'lar ise referance type'dır.

	//Array tanımı yapılırken boyut static olarak belirtilir
	//Slice'da ise böyle bi kural yoktur

	//Arraylar tanımlandıkları anca generate edilirler
	var myArr [4]int
	fmt.Println(myArr)
	myArr[0] = 5
	fmt.Println(myArr)

	//Slicelar tanımlandıkları anda generate edilmezler
	//Make metodu yardımıyla generate ettik
	var mySlc []int
	mySlc = make([]int, 4) //Make yöntemi ile oluşturma
	fmt.Println(mySlc)
	mySlc[0] = 10

}

func sliceExampleAppend() {
	//Bu şekilde yapılan bir tanımlamada slice null default value ile oluşur.
	var mySlice []int
	fmt.Printf("%#v", mySlice)

	mySlice = append(mySlice, 1, 2)

	fmt.Printf("%#v", mySlice)

	fmt.Println()

	//Bu şekilde yapılan bir tanımlamada ise obejct initilaze olur null durumda olmaz.
	slice := []int{}
	fmt.Printf("%#v", slice)

	slice = append(slice, 1, 2)

	fmt.Printf("%#v", slice)
}

func sliceExampleRemove() {
	mySlice := []int{}
	mySlice2 := []int{}

	mySlice = append(mySlice, 1, 2, 3, 4)
	mySlice2 = append(mySlice2, 5, 6, 7, 8)

	mySlice = append(mySlice, mySlice2...)

	fmt.Println(mySlice)

	//Bu kod bloğu şu anlama gelir. 2. indexden sonuna kadar al ve my slice' ata.
	//Böylece ilk iki eleman silinmiş olur
	mySlice = mySlice[2:]

	fmt.Println(mySlice)

	//Sondan iki elemanı atar ve bana yeni bir slice verir
	mySlice = mySlice[:2]

	fmt.Println(mySlice)

	//Silme işlemindeki kritik konu şudur.
	//Eğerki silme işlemi şu şekilde yapılırsa. [x:] bu pattern şunu ifade eder.
	//Dizinin 0. indexinden saymaya başlanır x. sayı hangi elemena geliyorsa oraya kadar olan elemanları atar
	//ve yeni bir slice oluşturur.

	//Eğerki pattern şu şekilde ise. [:x] bu pattern ise şunu ifader.
	//Dizinin 0. indexinden saymaya başla x. sayı hangi elemena kadar geliyor ise oraya kadar olan elemanları al ve yeni bir slice yap.
	//Geri kalanları at. Aşağıdaki örnekler ile anlatılmıştır

	mySlice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	//Burdaki işlemde 1 ve 2 itemlarını listeden atar.
	mySlice3 = mySlice3[2:]

	fmt.Println(mySlice3)

	//Burdaki işlemde 1 ve 2 yukarda atılmıştı.3 den saymaya başlar
	//6 oluncaya kadar sayar geldiği item hangisiyse oraya kadar olanları alır gerisi alınmaz.
	//Yani çıktı 3,4,5,6,7,8 şekilden olur
	mySlice3 = mySlice3[:6]

	fmt.Println(mySlice3)
}

func mapExample() {
	//Map c# daki dictionary gibi çalışır. Referance Type'dır.
	studentGrades := map[string]int{
		"Mertcan": 100,
		"Gamze":   20,
		"Ali":     30,
		"Veli":    40,
	}

	value, ok := studentGrades["Mertcan"]

	//Added Item
	studentGrades["Osman"] = 12

	//Remove Item
	delete(studentGrades, "Veli")

	for k, v := range studentGrades {
		fmt.Println(k, v)
	}

	fmt.Println(studentGrades)
	fmt.Println(value)
	fmt.Println(ok)
	fmt.Println(len(studentGrades))
}

func structExample() {
	//struct temel olarak bu şekilde kullanılır.
	//Burada bir başka employee oluşturabilmek için bir struct daha oluşturmak lazım
	//Bunu intance bazlı yönetebilmek için variable olark değil bir type olarak tanımlamamız gerekli

	var employee struct {
		name      string
		age       int
		isMarried bool
	}

	fmt.Println(employee)
	fmt.Printf("%#v\n", employee)

	employee.name = "Mertcan"
	employee.age = 10
	employee.isMarried = true

	fmt.Println(employee)

}

func structWithType() {
	//C# daki class mantığıda bu şekilde işletilir.

	type employee struct {
		name      string
		age       int
		isMarried bool
	}

	e1 := employee{
		name:      "Mertcan",
		age:       10,
		isMarried: true,
	}

	var e2 employee
	e2.name = "Gamze"
	e2.age = 12
	e2.isMarried = true

	fmt.Printf("%#v\n", e1)
	fmt.Printf("%#v\n", e2)
}

type Student struct {
	name     string
	age      int
	number   int
	isSucces bool
}

func structWithTypeGlobal() {
	s1 := Student{
		name:     "Mertcan",
		age:      10,
		isSucces: true,
		number:   1212,
	}

	var s2 Student
	s2.age = 10
	s2.name = "Gamze"
	s2.number = 1111
	s2.isSucces = false

	fmt.Printf("%#v\n", s1)
	fmt.Printf("%#v\n", s2)
}

type Employee struct {
	name      string
	age       int
	isMarried bool
}

type Manager struct { // => underlying type, employee ==> Defined Type, Named Type. Burda Manager struct tipinde değildir. Manager struct üzerine kurulmuş Maanger tipidir.
	Employee
	isHasDegree bool
}

func inheritaceExample() {
	//Normal oop yaklaşımında inheritance mantığında is a ilişkisi vardır.
	// yani yukarıdaki örneğe göre managerda bir çalışandır.
	//Ama go daki yaklaşımda has a ilişkisi vardır. Yani manager çalışan özelliklerini kapsar.

	var manager Manager
	manager.name = "Mertcan"
	manager.age = 10
	manager.isHasDegree = true
	manager.isHasDegree = true

	manager2 := Manager{}
	manager2.name = "Gamze"
	manager2.age = 10
	manager2.isHasDegree = true
	manager2.isHasDegree = true

	fmt.Println(manager)
	fmt.Println(manager2)
}

func anonymusStruct() {

	theAnimal := struct {
		name  string
		isMan bool
	}{name: "Mertcan", isMan: true}

	fmt.Println(theAnimal)
}

type mile float64      //Burda mile float64 tipinde değil. Float64 üzerinde kurulmuş kendi defined veri tipimiz.
type kilometer float64 // Kilometer float64 tipinde değildir. Float64 üzerinde kurulmuş kedni define veri tipimizdir.

func definedType() {
	var m1 mile
	m1 = 3.2
	fmt.Println(m1)
	fmt.Printf("%T %v", m1, m1)

	f1 := 4.4

	fmt.Println(f1)
	fmt.Println(m1 + mile(f1))

	//**************************

	fmt.Println()
	fmt.Printf("%T %v", (m1 + mile(f1)), (m1 + mile(f1)))

	//***************************

	m2 := mile(10)

	k1 := toKilometer(m2)

	fmt.Println(k1)
}

func toKilometer(m mile) kilometer {
	return kilometer(m * 1.6)
}

func pointerExample() {
	name := "Mertcan"

	fmt.Println(name)
	fmt.Println(&name) // ------> & adress operator. Değişkenin ram üzerinde tutulan adresini gösterir.

	number := 20

	fmt.Println(number)
	fmt.Println(&number)

	fmt.Println()

	fmt.Printf("%T %v\n", number, number)   //Burda number değişkenin değeri 20 tipi ise int dir.
	fmt.Printf("%T %v\n", &number, &number) //Burada ise number değişkenin adresi'nin değeri 0xc00000000b2020 iken tipi *int

	number2 := &number

	fmt.Printf("%T %v\n", number2, number2)
}

func pointerExample2() {
	x := 10

	fmt.Println(x)     //Değişkenin değeri
	fmt.Println(&x)    //Ram üzerindeki adresi
	fmt.Println(*(&x)) //Ram üzerindeki adresine * asterix operatörünü eklersek değerini döner.
}

func pointerExample3() {
	//x1 := 10
	//x2 := x1

	//fmt.Println(x1)
	//fmt.Println(x2)

	//x2 = 5 //sadece x2 nin değeri değişir çünkü value type dır.

	//fmt.Println(x1)
	//fmt.Println(x2)

	a1 := 10  //Bu atamada ramde bir alan ayarılır ve onun içine 10 değeri konur.Örnek: 0x123 [10].
	a2 := &a1 //Bu atamada ise xin adresi a2 değişkenine atanır. Ornek: 0x456 [0x123]

	fmt.Println(a1)
	fmt.Println(a2)

	fmt.Println("*************")

	*a2 = 3 //Burdaki atama şu anlama gelir.Git a2 nin değerinde var olan adrese. a2 yukarda =0x123 değerine eşittir.
	//Bu adrese git ve değerini 3 yap der. Bu adres a1'in adresidir. a1 3 değerini alır. 0x123 [3]

	a3 := &a2 //Burdaki atamada a3 değeri a2 nin adresi olur. 0x789 [0x456]
	a4 := *a3 //Burdaki atadama şu anlama gelir git a3 değerindeki adrese onun değerini a4' ata.

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
}

func pointerPassReferance() {
	//Go dili genel olara pass by value olarak çalışır. Yani metotlara value geçer.
	//Referance geçmek istersek eğer double methodunda olduğu gibi parametreyi almalıyız ve & operatörü ile göndermeliyiz
	x := 5
	fmt.Println(x)
	double(&x)
	fmt.Println(x)
}

func double(num *int) {
	fmt.Println(num)

	*num *= 2

	fmt.Println(num)
}
