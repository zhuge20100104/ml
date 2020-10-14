package main

import "testing"

// TestReadDirtyData 测试脏数据读取
func TestReadDirtyData(t *testing.T) {
	ReadDirtyData()
}

func TestReadStructData(t *testing.T) {
	ReadStructData()
}

func TestReadSimpleDF(t *testing.T) {
	ReadSimpleDF()
}

func TestReadAndFilterDF(t *testing.T) {
	ReadAndFilterDF()
}
