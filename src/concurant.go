package main
import "runtime"

func main() {
	runtime.GOMAXPROCS(8)
	//make a string channel
	ch := make(chan string)
	//make a chanel for passing done status
	doneCh := make(chan bool)
	go abcGen(ch)
	go printer(ch, doneCh)
	println("this comes first")

	//call to receive the message from the done channel (replaces a sleep)
	<-doneCh
	println("this comes second")


}

func abcGen(ch chan string) {
	for i := byte('a'); i <= byte('z'); i++ {
		//adds the letter to the channel
		ch <- string(i)

	}
	close(ch)

}

func printer(ch chan string, doneCh chan bool) {
	//loops over the channel, printing each element
	for i := range ch {
		println(i)
	}
	doneCh <- true
}