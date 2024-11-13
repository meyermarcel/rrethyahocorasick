package main

import (
	"bufio"
	"bytes"
	bobu "github.com/BobuSumisu/aho-corasick"
	anknown "github.com/anknown/ahocorasick"
	cloudflare "github.com/cloudflare/ahocorasick"
	"github.com/meyermarcel/ahocorasick"
	petard "github.com/petar-dambovaliev/aho-corasick"
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

func BenchmarkCloudflare1ModCompile(b *testing.B) {
	benchCloudflareCompile(b, 1)
}

func BenchmarkPetard1ModCompile(b *testing.B) {
	benchPetardCompile(b, 1)
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

func BenchmarkCloudflare10ModCompile(b *testing.B) {
	benchCloudflareCompile(b, 10)
}

func BenchmarkPetard10ModCompile(b *testing.B) {
	benchPetardCompile(b, 10)
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

func BenchmarkCloudflare100ModCompile(b *testing.B) {
	benchCloudflareCompile(b, 100)
}

func BenchmarkPetard100ModCompile(b *testing.B) {
	benchPetardCompile(b, 100)
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

func BenchmarkCloudflare1000ModCompile(b *testing.B) {
	benchCloudflareCompile(b, 1000)
}

func BenchmarkPetard1000ModCompile(b *testing.B) {
	benchPetardCompile(b, 1000)
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

func BenchmarkCloudflare10000ModCompile(b *testing.B) {
	benchCloudflareCompile(b, 10000)
}

func BenchmarkPetard10000ModCompile(b *testing.B) {
	benchPetardCompile(b, 10000)
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

func BenchmarkCloudflare1ModFind(b *testing.B) {
	benchCloudflareFind(b, 1)
}

func BenchmarkPetard1ModFind(b *testing.B) {
	benchPetardFind(b, 1)
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

func BenchmarkCloudflare10ModFind(b *testing.B) {
	benchCloudflareFind(b, 10)
}

func BenchmarkPetard10ModFind(b *testing.B) {
	benchPetardFind(b, 10)
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

func BenchmarkCloudflare100ModFind(b *testing.B) {
	benchCloudflareFind(b, 100)
}

func BenchmarkPetard100ModFind(b *testing.B) {
	benchPetardFind(b, 100)
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

func BenchmarkCloudflare1000ModFind(b *testing.B) {
	benchCloudflareFind(b, 1000)
}

func BenchmarkPetard1000ModFind(b *testing.B) {
	benchPetardFind(b, 1000)
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

// BobuSumisu/aho-corasick
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

// anknown/ahocorasick
func benchAnknownCompile(b *testing.B, every int) {
	patterns, err := readRunePatterns("./words.txt", every)
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
	patterns, err := readRunePatterns("./words.txt", 100)
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

// cloudflare/ahocorasick
func benchCloudflareCompile(b *testing.B, every int) {
	patterns, err := readBytePatterns("./words.txt", every)
	if err != nil {
		b.Error(err)
		return
	}

	for i := 0; i < b.N; i++ {
		_ = cloudflare.NewMatcher(patterns)
	}
}

func benchCloudflareFind(b *testing.B, every int) {
	patterns, err := readBytePatterns("./words.txt", 100)
	if err != nil {
		b.Error(err)
		return
	}
	m := cloudflare.NewMatcher(patterns)

	text, err := readBytes("./war-and-peace.txt", every)
	if err != nil {
		b.Error(err)
		return
	}

	for i := 0; i < b.N; i++ {
		m.Match(text)
	}
}

// petar-dambovaliev/aho-corasick
func benchPetardCompile(b *testing.B, every int) {
	patterns, err := readBytePatterns("./words.txt", every)
	if err != nil {
		b.Error(err)
		return
	}

	for i := 0; i < b.N; i++ {
		builder := petard.NewAhoCorasickBuilder(petard.Opts{
			AsciiCaseInsensitive: true,
			MatchOnlyWholeWords:  true,
			MatchKind:            petard.LeftMostLongestMatch,
			DFA:                  true,
		})
		_ = builder.BuildByte(patterns)
	}
}

func benchPetardFind(b *testing.B, every int) {
	patterns, err := readBytePatterns("./words.txt", 100)
	if err != nil {
		b.Error(err)
		return
	}
	builder := petard.NewAhoCorasickBuilder(petard.Opts{
		AsciiCaseInsensitive: true,
		MatchOnlyWholeWords:  true,
		MatchKind:            petard.LeftMostLongestMatch,
		DFA:                  true,
	})
	ac := builder.BuildByte(patterns)

	text, err := readBytes("./war-and-peace.txt", every)
	if err != nil {
		b.Error(err)
		return
	}

	textStr := string(text)

	for i := 0; i < b.N; i++ {
		ac.FindAll(textStr)
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

func readBytePatterns(filename string, every int) ([][]byte, error) {
	var dict [][]byte

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
			dict = append(dict, l)
		}
		i++
	}

	return dict, nil
}

func readRunePatterns(filename string, every int) ([][]rune, error) {
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
