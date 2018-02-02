package gamedata

import (
	"fmt"
	"testing"
)

func TestGetInt(t *testing.T) {
	fmt.Println(GetInt("test", 1, "INT_x"))
	fmt.Println(GetInt("testA", 1, "INT_y"))
}

func TestGetString(t *testing.T) {
	fmt.Println(GetString("test", 1, "STR_x"))
	fmt.Println(GetString("testA", 1, "STR_y"))
}

func TestGetFloat(t *testing.T) {
	fmt.Println(GetFloat("test", 1, "FLOAT_x"))
	fmt.Println(GetFloat("testA", 1, "FLOAT_y"))
}

func BenchmarkGetInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetInt("testA", 1, "INT_x")
	}
}

func BenchmarkGetString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetString("testA", 1, "STR_x")
	}
}
