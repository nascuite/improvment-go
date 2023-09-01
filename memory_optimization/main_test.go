package main

import (
	"github.com/brianvoe/gofakeit/v6"
	"os"
	"testing"
)

func BenchmarkOne(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		one()
	}
}

func BenchmarkTwo(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		second()
	}
}

func Test(t *testing.T) {

	res := gofakeit.LoremIpsumSentence(sentenceCount)

	_ = os.WriteFile("text.txt", []byte(res), 0644)

}
