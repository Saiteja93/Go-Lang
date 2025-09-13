package main





import "fmt"

func swap(a, b *int) {
    temp := *a
    *a = *b
    *b = temp
}

func main() {
    x, y := 7, 9
    fmt.Println("Before swap:", x, y)
    swap(&x, &y)
    fmt.Println("After swap:", x, y)
}
