package sysutil

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"testing"
)

func Test(t *testing.T) {
	if md5_OpenFalcon("1234567890123") != "ee76971984d25138a199ac90553401dd" {
		t.FailNow()
	}
	if md5_EncodeToString_Write_string("1234567890123") != "ee76971984d25138a199ac90553401dd" {
		t.FailNow()
	}
	if md5_EncodeToString_Write_byte("1234567890123") != "ee76971984d25138a199ac90553401dd" {
		t.FailNow()
	}
	if Md5("1234567890123") != "ee76971984d25138a199ac90553401dd" {
		t.FailNow()
	}
}

func Benchmark_md5_OpenFalcon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		md5_OpenFalcon("1234567890123")
	}
}

func Benchmark_md5_EncodeToString_Write_string(b *testing.B) {
	for i := 0; i < b.N; i++ {
		md5_EncodeToString_Write_string("1234567890123")
	}
}

func Benchmark_md5_EncodeToString_Write_byte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		md5_EncodeToString_Write_byte("1234567890123")
	}
}

func Benchmark_Md5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Md5("1234567890123")
	}
}

func md5_OpenFalcon(raw string) string {
	h := md5.New()
	io.WriteString(h, raw)

	return fmt.Sprintf("%x", h.Sum(nil))
}

func md5_EncodeToString_Write_string(raw string) string {
	h := md5.New()
	io.WriteString(h, raw)

	return hex.EncodeToString(h.Sum(nil))
}

func md5_EncodeToString_Write_byte(raw string) string {
	h := md5.New()
	h.Write([]byte(raw))

	return hex.EncodeToString(h.Sum(nil))
}
