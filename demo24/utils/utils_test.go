package utils

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, but got: %v", want, got)
	}
}

func TestMoreSplit(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

// 测试组
func TestSplitGroup(t *testing.T) {
	type test struct {
		input string
		sep   string
		want  []string
	}

	tests := []test{
		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
		{input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}

	for _, ts := range tests {
		got := Split(ts.input, ts.sep)
		if !reflect.DeepEqual(got, ts.want) {
			t.Errorf("expected:%#v, got:%#v", ts.want, got)
		}
	}
}

// 子测试
func TestSplitGroup2(t *testing.T) {
	type test struct {
		input string
		sep   string
		want  []string
	}

	tests := map[string]test{ // 测试用例使用map存储
		"simple":    {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep": {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":  {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{
			"", "河有", "又有河"}},
	}

	for name, ts := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(ts.input, ts.sep)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("expected:%#v, got:%#v", ts.want, got)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}
