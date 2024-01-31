package channel

//Aşağıdaki gibi tanımlanır ve initialize edilirken bir size verilir. Channel bu size ulaşana kadar blocklanmaz ulaştığı an blocklanır.
//Aşağıdaki örnekte 6'ya gelince blocklanıcak ve uygulama panic atıcak. Çünkü 6. adımda blocklandığı yerde başka değer set edicek bloğu açacak bir code yok.

func BufferedChannel() {
	//channel initialization
	bufferedChan := make(chan int, 5)

	bufferedChan <- 1
	bufferedChan <- 2
	bufferedChan <- 3
	bufferedChan <- 4
	bufferedChan <- 5
	bufferedChan <- 6

}
