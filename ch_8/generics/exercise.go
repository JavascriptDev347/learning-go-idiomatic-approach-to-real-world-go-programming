package main

func main() {

}

// Type parameter T ya'ni tip uchun o'zgaruvchi kvadrat qavs ichida yoziladi.
// T qaysi tiplar bo'lishini belgilaydi: int va float64. Bu yerda ikkita tip belgilangan.
func Max[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}
