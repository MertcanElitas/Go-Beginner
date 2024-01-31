package channel

// 1 numaralı case'de değer olarak 1 de basabilir 2 de basabilir. Net bir durum yoktur nedeni ise şudur select scope içerisinde kalan case'lere random olarak düşer ve çalıştırır. Bundan dolayı
// 1 de basabilir 2 de basabilir.

func ChannelWithSelect() {
	chan1 := make(chan int, 1)
	chan1 <- 1

	chan2 := make(chan int, 2)
	chan2 <- 2

	// Bu case'de değer olarak 1 de basabilir 2 de basabilir. Net bir durum yoktur nedeni ise şudur select scope içerisinde kalan case'lere random olarak düşer ve çalıştırır.
	// Bundan dolayı 1 de basabilir 2 de basabilir.
	// select {
	// case c1Val := <-chan1:
	// 	fmt.Println(c1Val)
	// case c2Val := <-chan2:
	// 	fmt.Println(c2Val)
	// }

	// Burda ise çalıştırdığımızda bir hata ile karşılaşırız. Sebebi ise şudur select ifadesi çalışır random olarak chan1 veya chan2 yi dinler ve değeri basar
	//daha sonra for döngüsünün limiti olmadığı için bir daha döner bu sefer ilk seferde hangisi çalıştırmadı ise onu print eder. For tekrar çalışır ve bu sefer
	// uygulama hat verir nedeni şu select chan1 veya chan2 ye birdaha subscrine olur ve onlara data yazacak bir code kalmadığı için exception fırlatır
	////this will not end
	// for {
	// 	select {
	// 	case c1Val := <-chan1:
	// 		fmt.Println(c1Val)
	// 	case c2Val := <-chan2:
	// 		fmt.Println(c2Val)
	// 	}
	// }
	//

	// Burda ise çalıştırdığımızda bir hata ile karşılaşılmaz ama uygulama sonsuz döngüye girer kapanmaz.
	// Sebebi ise şudur select ifadesi çalışır random olarak chan1 veya chan2 yi dinler ve değeri basar
	// daha sonra for döngüsünün limiti olmadığı için bir daha döner bu sefer ilk seferde hangisi çalıştırmadı ise onu print eder.
	// For tekrar çalışır ve bu sefer default bloğu çalışacağı için for başa döner ve bu durum sürekli default bloğuna düşeceği için
	// devam eder ve sonsuz bir loop oluşur.

	////this will not end
	// for {
	// 	select {
	// 	case c1Val := <-chan1:
	// 		fmt.Println(c1Val)
	// 	case c2Val := <-chan2:
	// 		fmt.Println(c2Val)
	// 	default:
	// 		break
	// 	}
	//}

	// Burda ise çalıştırdığımızda bir herhangibir problem ile karşılaşılmaz.
	// Default blogunda berak etmeyip flag'i set ettiğimiz ve for döngüsünüde bu 
	//flag değerine göre kurguladığımız için for dönmeye devam etmez.
	// Aynı şekilde sonuç olarak bazen ilk chan1 değerini bazen chan2 değerini okur stabilite yoktur.
	// var done bool
	// for !done {
	// 	//selects randomly from the channels
	// 	//output order is non-deterministic
	// 	select {
	// 	case c1Val := <-chan1:
	// 		fmt.Println(c1Val)
	// 	case c2Val := <-chan2:
	// 		fmt.Println(c2Val)
	// 	default:
	// 		done = true
	// 	}
	// }
}
