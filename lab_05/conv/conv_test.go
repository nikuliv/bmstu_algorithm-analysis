package conv

import (
	"testing"
)

func Benchmark10(b *testing.B) {
	N := 10
	ch := make(chan int)

	for i := 0; i < b.N; i++ {
		Conveyor(N, ch)
		<-ch
	}
}

func Benchmark20(b *testing.B) {
	N := 20

	for i := 0; i < b.N; i++ {
		ch := make(chan int)
		Conveyor(N, ch)
		<-ch
	}
}

func Benchmark30(b *testing.B) {
	N := 30

	for i := 0; i < b.N; i++ {
		ch := make(chan int)
		Conveyor(N, ch)
		<-ch
	}
}

func Benchmark40(b *testing.B) {
	N := 40

	for i := 0; i < b.N; i++ {
		ch := make(chan int)
		Conveyor(N, ch)
		<-ch
	}
}

func Benchmark50(b *testing.B) {
	N := 50

	for i := 0; i < b.N; i++ {
		ch := make(chan int)
		Conveyor(N, ch)
		<-ch
	}
}

//func BenchmarkStandart10(b *testing.B) {
//	N := 10
//	var mes *Message
//	mes = new(Message)
//	for i := 0; i < N; i++ {
//		massage(mes)
//		encrypt(mes)
//		decrypt(mes)
//	}
//}
//
//func BenchmarkStandart20(b *testing.B) {
//	N := 20
//	var mes *Message
//	mes = new(Message)
//	for i := 0; i < N; i++ {
//		massage(mes)
//		encrypt(mes)
//		decrypt(mes)
//	}
//}
//
//func BenchmarkStandart30(b *testing.B) {
//	N := 30
//	var mes *Message
//	mes = new(Message)
//	for i := 0; i < N; i++ {
//		massage(mes)
//		encrypt(mes)
//		decrypt(mes)
//	}
//}
//
//func BenchmarkStandart40(b *testing.B) {
//	N := 40
//	var mes *Message
//	mes = new(Message)
//	for i := 0; i < N; i++ {
//		massage(mes)
//		encrypt(mes)
//		decrypt(mes)
//	}
//}
//
//func BenchmarkStandart50(b *testing.B) {
//	N := 50
//	var mes *Message
//	mes = new(Message)
//	for i := 0; i < N; i++ {
//		massage(mes)
//		encrypt(mes)
//		decrypt(mes)
//	}
//}
