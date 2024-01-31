package concurency

import "fmt"

//Bu kod bloğunu çalıştırdığımız zaman şöyle bir senaryo ile karşılaşırız.
//Kod çalışmaya başlar ve ilk olarak go keywordü ile oluşturulmuş ve tetiklenmiş
//bir function ile karşılaşır. Daha sonra go keywordü ile oluşturulduğu için ayrı bir gorotuine ile
//bu işlemi manage etmeye çalışr ancak herhangi bir şekilde bu işlem bekletilmediği için blocklanmadığı için
//main goroutine bi alt satırdaki "hello from main" çıktısını yazıp main procesi tamamladığı için biz "hello from goroutine"
//çıktısını göremeyiz. Ama bu da sürekli yaşanacak diye bir kural yoktur. 10 kez çalıştırsak belki 1 sefer "hello from goroutine"
// yazdığını görebiliriz. Ayrıca açtığımız goroutine'nin onu işleme hızı ile alakalıdır.

func PrintHelloGoroutine() {
	go func() {
		fmt.Println("hello from goroutine")
	}()

	fmt.Println("hello from main")
}
