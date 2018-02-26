package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"strings"
)

var (
	suffixes = []string{"", "Unsafe"}
	cases    = []string{"Lower", "Title", "Upper"}
	funcs    = []string{
		"IsDigit", "IsGraphic", "IsLetter", "IsLower", "IsMark", "IsNumber",
		"IsPrint", "IsPunct", "IsSpace", "IsSymbol", "IsTitle", "IsUpper",
		"SimpleFold",
	}
)

func main() {
	errs := make(chan error, 3)

	go func() {
		errs <- generateTables()
	}()

	go func() {
		errs <- generateTests()
	}()

	go func() {
		errs <- generateBenchmarks()
	}()

	for i := 0; i < 3; i++ {
		if err := <-errs; err != nil {
			panic(err)
		}
	}
}

func generateTables() error {
	buf := bytes.NewBufferString(`
package runes

import (
	"unicode"
	"unicode/utf8"
)

const (`)

	for i, v := range funcs {
		if v == "SimpleFold" {
			continue
		}

		buf.WriteByte('\n')
		buf.WriteString(strings.ToLower(v[2:]))

		if i == 0 {
			buf.WriteString(" = 1 << iota")
		}
	}

	buf.WriteString(`
)

type char struct {
	lower  rune
	title  rune
	upper  rune
	folded rune
	flags  uint
}

var properties [unicode.MaxRune + 1]char

func init() {
	for r := rune(0); r <= unicode.MaxRune; r++ {
		if !utf8.ValidRune(r) {
			continue
		}

		var v char
`)

	for _, v := range funcs {
		if v == "SimpleFold" {
			continue
		}

		fmt.Fprintf(buf, `

      if unicode.%s(r) {
        v.flags |= %s
      }`, v, strings.ToLower(v[2:]))
	}

	buf.WriteString(`

		v.lower = unicode.ToLower(r)
		v.title = unicode.ToTitle(r)
		v.upper = unicode.ToUpper(r)
		v.folded = unicode.SimpleFold(r)
		properties[r] = v
	}
}`)

	b, err := format.Source(buf.Bytes())

	if err != nil {
		return err
	}

	return ioutil.WriteFile("tables.go", b, 0644)
}

func generateTests() error {
	buf := bytes.NewBufferString(`
package runes

import (
	"io/ioutil"
	"strings"
	"testing"
	"unicode"
)

var (
	testString    string
	testStringLen int64
)

func init() {
	b, err := ioutil.ReadFile("./testdata/corpus.txt")

	if err != nil {
		panic(err)
	}

	testString = string(b)
	testStringLen = int64(len(testString))
}

func TestFuncs(t *testing.T) {
	for r := rune(0); r <= unicode.MaxRune; r++ {`)

	for i, v := range funcs {
		if i > 0 {
			buf.WriteString("\n")
		}

		for j, suffix := range suffixes {
			if v == "SimpleFold" && suffix == "Unsafe" {
				continue
			}

			if j > 0 {
				buf.WriteString("\n")
			}

			fmt.Fprintf(buf, `
				if %[2]s%[1]s(r) != unicode.%[2]s(r) {
					t.Fatalf("%[2]s%[1]s failed for: %%U, expected %%v, got %%v",
						r, unicode.%[2]s(r), %[2]s%[1]s(r))
				}`, suffix, v)
		}
	}

	for _, v := range cases {
		fmt.Fprintf(buf, `

				if To%[1]s(r) != unicode.To%[1]s(r) {
					t.Fatalf("%[1]s failed for: %%U, expected %%U, got %%U",
						r, unicode.To%[1]s(r), To%[1]s(r))
				}

				if To(unicode.%[1]sCase, r) != unicode.To(unicode.%[1]sCase, r) {
					t.Fatalf("To failed for: %%U, expected %%U, got %%U",
						r, unicode.To(unicode.%[1]sCase, r), To(unicode.%[1]sCase, r))
				}`, v)
	}

	buf.WriteString(`
		}

		const (
			_case = 99
			r			= 'A'
		)

		if To(_case, r) != unicode.To(_case, r) {
			t.Fatalf("To failed for: %U, expected %U, got %U",
				r, unicode.To(_case, r), To(_case, r))
		}`)

	for i, v := range cases {
		if i > 0 {
			buf.WriteByte('\n')
		}

		fmt.Fprintf(buf, `

			if strings.Map(To%[1]s, testString) != strings.Map(unicode.To%[1]s, testString) {
				t.Fatal("To%[1]s failed ")
			}`, v)
	}

	buf.WriteString("\n}")

	b, err := format.Source(buf.Bytes())

	if err != nil {
		return err
	}

	return ioutil.WriteFile("runes_test.go", b, 0644)
}

func generateBenchmarks() error {
	for _, prefix := range []string{"", "unicode."} {
		buf := bytes.NewBufferString("// +build ")

		if prefix == "" {
			buf.WriteByte('!')
		}

		buf.WriteString(`std

        package runes

        import (
					"strings"
					"testing"
					"unicode"
					"unicode/utf8"
				)
			`)

		for _, v := range funcs {
			if v == "SimpleFold" {
				continue
			}

			for _, suffix := range suffixes {
				title := suffix

				if prefix != "" {
					suffix = ""
				}

				fmt.Fprintf(buf, `

					func Benchmark%[1]s%[5]s(b *testing.B) {
						pos, neg := getTestCase(%[2]s)

						b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
						b.ResetTimer()

						for i := 0; i <= b.N; i++ {
							%[3]s%[1]s%[4]s(pos)
							%[3]s%[1]s%[4]s(neg)
						}
					}`, v, strings.ToLower(v[2:]), prefix, suffix, title)
			}
		}

		for _, suffix := range suffixes {
			title := suffix

			if prefix != "" {
				suffix = ""
			}

			fmt.Fprintf(buf, `

				func BenchmarkSimpleFold%[3]s(b *testing.B) {
					pos, neg := '\u212A', 'a'

					b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
					b.ResetTimer()

					for i := 0; i <= b.N; i++ {
						%[1]sSimpleFold%[2]s(pos)
						%[1]sSimpleFold%[2]s(neg)
					}
				}`, prefix, suffix, title)
		}

		for _, v := range cases {
			for _, suffix := range suffixes {
				title := suffix

				if prefix != "" {
					suffix = ""
				}

				fmt.Fprintf(buf, `

					func BenchmarkTo%[1]s%[4]s(b *testing.B) {
						b.SetBytes(testStringLen)
						b.ResetTimer()

						for i := 0; i <= b.N; i++ {
							strings.Map(%[2]sTo%[1]s%[3]s, testString)
						}
					}`, v, prefix, suffix, title)
			}
		}

		for _, v := range cases {
			for _, suffix := range suffixes {
				title := suffix

				if prefix != "" {
					suffix = ""
				}

				fmt.Fprintf(buf, `

					func BenchmarkStringTo%[1]sCase%[4]s(b *testing.B) {
						f := func (r rune) rune {
							return %[2]sTo%[3]s(unicode.%[1]sCase, r)
						}

						b.SetBytes(testStringLen)
	    			b.ResetTimer()

						for i := 0; i <= b.N; i++ {
							strings.Map(f, testString)
						}
					}`, v, prefix, suffix, title)
			}
		}

		b, err := format.Source(buf.Bytes())

		if err != nil {
			return err
		}

		file := "bench_"

		if prefix != "" {
			file += "std_"
		}

		if err = ioutil.WriteFile(file+"test.go", b, 0644); err != nil {
			return err
		}
	}

	return nil
}
