ic := make(chan int)

sc := make(chan string, 10)

ic <- 3
n := <-sc

chan Sushi
chan<- string
<-chan int
