// This file has automatically been generated on Wed Feb 26 02:10:15 +05 2020.
// DO NOT EDIT.
package testing

import (
	"testing"
	_ "unsafe"
)

//go:linkname bfailed testing.sub_bfailed
func bfailed(c *testing.B) bool {
	return c.Failed()
}

//go:linkname BFailed testing.sub_bfailed
//go:noescape
func BFailed(c *testing.B) bool

//go:linkname brun testing.sub_brun
func brun(b *testing.B, name string, f func(b *testing.B)) bool {
	return b.Run(name, f)
}

//go:linkname BRun testing.sub_brun
//go:noescape
func BRun(b *testing.B, name string, f func(b *testing.B)) bool

//go:linkname bskipped testing.sub_bskipped
func bskipped(c *testing.B) bool {
	return c.Skipped()
}

//go:linkname BSkipped testing.sub_bskipped
//go:noescape
func BSkipped(c *testing.B) bool

//go:linkname Benchmark testing.Benchmark
//go:noescape
func Benchmark(f func(b *testing.B)) testing.BenchmarkResult

//go:linkname AllocsPerRun testing.AllocsPerRun
//go:noescape
func AllocsPerRun(runs int, f func()) float64

//go:linkname CoverMode testing.CoverMode
//go:noescape
func CoverMode() string

//go:linkname Coverage testing.Coverage
//go:noescape
func Coverage() float64

//go:linkname Main testing.Main
//go:noescape
func Main(matchString func(pat, str string) (bool, error), tests []testing.InternalTest, benchmarks []testing.InternalBenchmark, examples []testing.InternalExample)

//go:linkname trun testing.sub_trun
func trun(t *testing.T, name string, f func(t *testing.T)) bool {
	return t.Run(name, f)
}

//go:linkname TRun testing.sub_trun
//go:noescape
func TRun(t *testing.T, name string, f func(t *testing.T)) bool

//go:linkname benchmarkresultallocsperop testing.sub_benchmarkresultallocsperop
func benchmarkresultallocsperop(r testing.BenchmarkResult) int64 {
	return r.AllocsPerOp()
}

//go:linkname BenchmarkResultAllocsPerOp testing.sub_benchmarkresultallocsperop
//go:noescape
func BenchmarkResultAllocsPerOp(r testing.BenchmarkResult) int64

//go:linkname mrun testing.sub_mrun
func mrun(m *testing.M) int {
	return m.Run()
}

//go:linkname MRun testing.sub_mrun
//go:noescape
func MRun(m *testing.M) int

//go:linkname tfailed testing.sub_tfailed
func tfailed(c *testing.T) bool {
	return c.Failed()
}

//go:linkname TFailed testing.sub_tfailed
//go:noescape
func TFailed(c *testing.T) bool

//go:linkname tname testing.sub_tname
func tname(c *testing.T) string {
	return c.Name()
}

//go:linkname TName testing.sub_tname
//go:noescape
func TName(c *testing.T) string

//go:linkname RunExamples testing.RunExamples
//go:noescape
func RunExamples(matchString func(pat, str string) (bool, error), examples []testing.InternalExample) bool

//go:linkname benchmarkresultallocedbytesperop testing.sub_benchmarkresultallocedbytesperop
func benchmarkresultallocedbytesperop(r testing.BenchmarkResult) int64 {
	return r.AllocedBytesPerOp()
}

//go:linkname BenchmarkResultAllocedBytesPerOp testing.sub_benchmarkresultallocedbytesperop
//go:noescape
func BenchmarkResultAllocedBytesPerOp(r testing.BenchmarkResult) int64

//go:linkname benchmarkresultmemstring testing.sub_benchmarkresultmemstring
func benchmarkresultmemstring(r testing.BenchmarkResult) string {
	return r.MemString()
}

//go:linkname BenchmarkResultMemString testing.sub_benchmarkresultmemstring
//go:noescape
func BenchmarkResultMemString(r testing.BenchmarkResult) string

//go:linkname pbnext testing.sub_pbnext
func pbnext(pb *testing.PB) bool {
	return pb.Next()
}

//go:linkname PBNext testing.sub_pbnext
//go:noescape
func PBNext(pb *testing.PB) bool

//go:linkname benchmarkresultstring testing.sub_benchmarkresultstring
func benchmarkresultstring(r testing.BenchmarkResult) string {
	return r.String()
}

//go:linkname BenchmarkResultString testing.sub_benchmarkresultstring
//go:noescape
func BenchmarkResultString(r testing.BenchmarkResult) string

//go:linkname RunBenchmarks testing.RunBenchmarks
//go:noescape
func RunBenchmarks(matchString func(pat, str string) (bool, error), benchmarks []testing.InternalBenchmark)

//go:linkname RunTests testing.RunTests
//go:noescape
func RunTests(matchString func(pat, str string) (bool, error), tests []testing.InternalTest) bool

//go:linkname Short testing.Short
//go:noescape
func Short() bool

//go:linkname bname testing.sub_bname
func bname(c *testing.B) string {
	return c.Name()
}

//go:linkname BName testing.sub_bname
//go:noescape
func BName(c *testing.B) string

//go:linkname Verbose testing.Verbose
//go:noescape
func Verbose() bool

//go:linkname benchmarkresultnsperop testing.sub_benchmarkresultnsperop
func benchmarkresultnsperop(r testing.BenchmarkResult) int64 {
	return r.NsPerOp()
}

//go:linkname BenchmarkResultNsPerOp testing.sub_benchmarkresultnsperop
//go:noescape
func BenchmarkResultNsPerOp(r testing.BenchmarkResult) int64

//go:linkname MainStart testing.MainStart
//go:noescape
func MainStart(deps testDeps, tests []testing.InternalTest, benchmarks []testing.InternalBenchmark, examples []testing.InternalExample) *testing.M

//go:linkname tskipped testing.sub_tskipped
func tskipped(c *testing.T) bool {
	return c.Skipped()
}

//go:linkname TSkipped testing.sub_tskipped
//go:noescape
func TSkipped(c *testing.T) bool
