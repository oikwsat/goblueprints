package trace

import (
	"fmt"
	"io"
)

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

// OffはTraceメソッドの呼び出しを無視するTracerを返します

func Off() Tracer {
	return &nilTracer{}
}
