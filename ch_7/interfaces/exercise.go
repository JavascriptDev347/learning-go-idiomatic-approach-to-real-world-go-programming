package main

import "fmt"

type Stringer interface {
	String() string
}

type Incrementer interface {
	Increment()
}

type Counter struct {
	X int
}

type Doubler interface {
	Double()
}

type DoubleInt int

func (d *DoubleInt) Double() {
	*d = *d * 2
}

type DoubleIntSlice []int

func (d DoubleIntSlice) Double() {
	for i := range d {
		d[i] = d[i] * 2
	}
}

func DoublerCompare(d1, d2 Doubler) {
	fmt.Println(d1 == d2)
}

func main() {
	var myStringer fmt.Stringer
	var myIncrementer Incrementer
	// pointerCounter := &Counter{}
	valueCounter := Counter{}
	// myStringer = pointerCounter    // ok myStringer = valueCounter
	// myIncrementer = pointerCounter // ok myIncrementer = valueCounter

	fmt.Println(myStringer, myIncrementer, valueCounter)

	var di DoubleInt = 10
	var di2 DoubleInt = 10
	var dis = DoubleIntSlice{1, 2, 3}
	// var dis2 = DoubleIntSlice{1, 2, 3}
	// false sababi bu yerda ularning pointerini solishtiryabdi qiymatini emas
	DoublerCompare(&di, &di2)
	// false sababi ularning type match emas
	DoublerCompare(&di, dis)

	// bu compiledan xatosiz o'tadi lekin panic beradi, sababi type bir xil lekin
	// slicelar == orqali comparable qilib bo'lmaydi
	// DoublerCompare(dis, dis2)

}
