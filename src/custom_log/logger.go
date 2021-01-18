package logger

import "fmt"

// 构造函数
type Logger struct {
}

// class log
func NewLog() Logger {
	return Logger{}
}

func (l Logger) Debug(msg string) {
	fmt.Println(msg)
}

// func main() {
// 	fileObj, err := os.OpenFile("src/custom_log/test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		fmt.Println("open file failed, err: %v\n", err)
// 		return
// 	}
// 	log.SetOutput(fileObj)

// 	for {
// 		log.Print("this is log")
// 		time.Sleep(time.Second * 3)
// 	}
// }
