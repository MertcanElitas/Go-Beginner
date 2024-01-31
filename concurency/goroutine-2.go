package concurency

import (
	"fmt"
	"time"
)

//Aşağıdaki kod bloğu çalıştığında oluşacak durum şudur. for ile dönerek 10 adet goroutine schedula edilir.
//Daha sonrasında bu goroutine'ler çalışmaya başladığı anda i değeri ne ise (Genel olarak %90 oranında 10 olur) o print edilir.
//Yani run ettiğimizde 10 adet 10 görme olasılığımız yüksek. Ama döngü dönme sayısı değiştikce bu sayıda farklı olacaktır.
//Eğer i içerdeki functiona parametre geçilirse 0 dan 10 a kadar bütün rakamları basar ama dağınık basar çünkü hangi gorotuine önce
//çalışacak bilemeyiz.

func PrintGoRoutineWithSleep() {

	for i := 0; i <= 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(time.Second * 3)
}
