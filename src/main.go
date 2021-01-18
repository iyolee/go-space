package main

import (
	logger "custom_log"
	datafile "file"
	"fmt"
	"greeting"
	"io"
	"os"
	"reflect"
	"time"
)

func main() {
	log := logger.NewLog()
	log.Debug("haha")
	arr := [...]int{1, 2}
	bb := &arr
	fmt.Println(reflect.TypeOf(bb))
	fmt.Println("hello")
	fmt.Println(reflect.TypeOf(25))
	fmt.Println((reflect.TypeOf(false)))

	greeting.Hello()

	numbers, err := datafile.GetFloats("src/data.txt")
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println(numbers)
	fmt.Println(sum)

	fmt.Println(time.Now())

	//seconds := time.Now().Unix()
	//rand.Seed(seconds)
	//target := rand.Intn(100) + 1
	//
	//reader := bufio.NewReader(os.Stdin)
	//success := false
	//
	//for guesses := 0; guesses < 10; guesses++ {
	//	input, err := reader.ReadString('\n')
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	input = strings.TrimSpace(input)
	//	guess, err := strconv.Atoi(input)
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	if guess < target {
	//		fmt.Println("small")
	//	} else if guess > target {
	//		fmt.Println("big")
	//	} else {
	//		success = true
	//		fmt.Println("good")
	//		break
	//	}
	//}
	//if !success {
	//	fmt.Println("sorry")
	//}

	fileObj, err := os.Open("src/data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileObj.Close()

	var tmp [128]byte
	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(tmp[:n]))
	}
}
