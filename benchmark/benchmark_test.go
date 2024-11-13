package main

import (
	"bufio"
	"bytes"
	bobu "github.com/BobuSumisu/aho-corasick"
	anknown "github.com/anknown/ahocorasick"
	"github.com/meyermarcel/ahocorasick"
	"io"
	"os"
	"testing"
)

func BenchmarkMeyermarcelAhocorasick1ModCompileByteSlices(b *testing.B) {
	benchMeyermarcelAhocorasickCompileByteSlices(b, 1)
}

func BenchmarkBobu1ModCompile(b *testing.B) {
	benchBobuCompile(b, 1)
}

func BenchmarkAnknown1ModCompile(b *testing.B) {
	benchAnknownCompile(b, 1)
}

func BenchmarkMeyermarcelAhocorasick10ModCompileByteSlices(b *testing.B) {
	benchMeyermarcelAhocorasickCompileByteSlices(b, 10)
}

func BenchmarkBobu10ModCompile(b *testing.B) {
	benchBobuCompile(b, 10)
}

func BenchmarkAnknown10ModCompile(b *testing.B) {
	benchAnknownCompile(b, 10)
}

func BenchmarkMeyermarcelAhocorasick100ModCompileByteSlices(b *testing.B) {
	benchMeyermarcelAhocorasickCompileByteSlices(b, 100)
}

func BenchmarkBobu100ModCompile(b *testing.B) {
	benchBobuCompile(b, 100)
}

func BenchmarkAnknown100ModCompile(b *testing.B) {
	benchAnknownCompile(b, 100)
}

func BenchmarkMeyermarcelAhocorasick1000ModCompileByteSlices(b *testing.B) {
	benchMeyermarcelAhocorasickCompileByteSlices(b, 1000)
}

func BenchmarkBobu1000ModCompile(b *testing.B) {
	benchBobuCompile(b, 1000)
}

func BenchmarkAnknown1000ModCompile(b *testing.B) {
	benchAnknownCompile(b, 1000)
}

func BenchmarkMeyermarcelAhocorasick10000ModCompileByteSlices(b *testing.B) {
	benchMeyermarcelAhocorasickCompileByteSlices(b, 10000)
}

func BenchmarkBobu10000ModCompile(b *testing.B) {
	benchBobuCompile(b, 10000)
}

func BenchmarkAnknown10000ModCompile(b *testing.B) {
	benchAnknownCompile(b, 10000)
}

func BenchmarkMeyermarcelAhocorasick1ModFindAllByteSlice(b *testing.B) {
	benchMeyermarcelAhocorasickFindAllByteSlice(b, 1)
}

func BenchmarkBobu1ModFind(b *testing.B) {
	benchBobuFind(b, 1)
}

func BenchmarkAnknown1ModFind(b *testing.B) {
	benchAnknownFind(b, 1)
}

func BenchmarkMeyermarcelAhocorasick10ModFindAllByteSlice(b *testing.B) {
	benchMeyermarcelAhocorasickFindAllByteSlice(b, 10)
}

func BenchmarkBobu10ModFind(b *testing.B) {
	benchBobuFind(b, 10)
}

func BenchmarkAnknown10ModFind(b *testing.B) {
	benchAnknownFind(b, 10)
}

func BenchmarkMeyermarcelAhocorasick100ModFindAllByteSlice(b *testing.B) {
	benchMeyermarcelAhocorasickFindAllByteSlice(b, 100)
}

func BenchmarkBobu100ModFind(b *testing.B) {
	benchBobuFind(b, 100)
}

func BenchmarkAnknown100ModFind(b *testing.B) {
	benchAnknownFind(b, 100)
}

func BenchmarkMeyermarcelAhocorasick1000ModFindAllByteSlice(b *testing.B) {
	benchMeyermarcelAhocorasickFindAllByteSlice(b, 1000)
}

func BenchmarkBobu1000ModFind(b *testing.B) {
	benchBobuFind(b, 1000)
}

func BenchmarkAnknown1000ModFind(b *testing.B) {
	benchAnknownFind(b, 1000)
}

func benchMeyermarcelAhocorasickCompileByteSlices(b *testing.B, every int) {
	patterns, err := readLines("./words.txt", every)
	if err != nil {
		b.Error(err)
		return
	}

	for i := 0; i < b.N; i++ {
		ahocorasick.CompileByteSlices(patterns)
	}
}

func benchMeyermarcelAhocorasickFindAllByteSlice(b *testing.B, every int) {
	patterns, err := readLines("./words.txt", 100)
	if err != nil {
		b.Error(err)
		return
	}

	text, err := readBytes("./war-and-peace.txt", every)
	if err != nil {
		b.Error(err)
		return
	}

	m := ahocorasick.CompileByteSlices(patterns)
	for i := 0; i < b.N; i++ {
		m.FindAllByteSlice(text)
	}
}

// BOBU
func benchBobuCompile(b *testing.B, every int) {
	patterns, err := readLines("./words.txt", every)
	if err != nil {
		b.Error(err)
		return
	}

	for i := 0; i < b.N; i++ {
		bobu.NewTrieBuilder().AddPatterns(patterns).Build()
	}
}

func benchBobuFind(b *testing.B, every int) {
	patterns, err := readLines("./words.txt", 100)
	if err != nil {
		b.Error(err)
		return
	}

	text, err := readBytes("./war-and-peace.txt", every)
	if err != nil {
		b.Error(err)
		return
	}

	t := bobu.NewTrieBuilder().AddPatterns(patterns).Build()
	for i := 0; i < b.N; i++ {
		t.Match(text)
	}
}

// ANKNOWN
func benchAnknownCompile(b *testing.B, every int) {
	patterns, err := readRunes("./words.txt", every)
	if err != nil {
		b.Error(err)
		return
	}

	for i := 0; i < b.N; i++ {
		m := new(anknown.Machine)
		m.Build(patterns)
	}
}

func benchAnknownFind(b *testing.B, every int) {
	patterns, err := readRunes("./words.txt", 100)
	if err != nil {
		b.Error(err)
		return
	}
	m := new(anknown.Machine)
	m.Build(patterns)

	text, err := readBytes("./war-and-peace.txt", every)
	if err != nil {
		b.Error(err)
		return
	}
	textRunes := bytes.Runes(text)

	for i := 0; i < b.N; i++ {
		m.MultiPatternSearch(textRunes, false)
	}
}

func readLines(fname string, every int) ([][]byte, error) {
	file, err := os.Open(fname)
	defer file.Close()
	var pattens [][]byte
	if err != nil {
		return pattens, err
	}

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		if i%every == 0 {
			text := scanner.Text()
			if len(text) > 0 {
				pattens = append(pattens, []byte(text))
			}
		}
		i++
	}
	return pattens, nil
}

func readBytes(fname string, every int) ([]byte, error) {
	file, err := os.Open(fname)
	defer file.Close()
	var bytes []byte
	if err != nil {
		return bytes, err
	}

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		if i%every == 0 {
			bytes = append(bytes, []byte(scanner.Text())...)
		}
		i++
	}
	return bytes, nil
}

func readRunes(filename string, every int) ([][]rune, error) {
	dict := [][]rune{}

	f, err := os.OpenFile(filename, os.O_RDONLY, 0660)
	if err != nil {
		return nil, err
	}

	r := bufio.NewReader(f)
	i := 0
	for {
		l, err := r.ReadBytes('\n')
		if err != nil || err == io.EOF {
			break
		}
		l = bytes.TrimSpace(l)

		if i%every == 0 {
			dict = append(dict, bytes.Runes(l))
		}
		i++
	}

	return dict, nil
}
