package hello

import "fmt"

const HelloPrefix = "Hello, "

func Hello(name string) string {
	if name != "" {
		return HelloPrefix + name
	}
	return HelloPrefix + "world"
}

func main() {
	fmt.Print(Hello("Syamm"))
}
