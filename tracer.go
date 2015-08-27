// Copyright 2015 Ezequiel L. Aceto. All rights reserved.
// Use of this source code is governed by a GNU v2.0-style
// license that can be found in the LICENSE file.


package tracer

import (
	"time"
	"io"
	"os"
	"fmt"
)

// Trace to stdout the entering to a function call
func Trace(s string) ( io.Writer,  string,  time.Time )   {
	return TraceToWriter(os.Stdout,s)
}

// Trace to a writer the entering to a function call
func TraceToWriter(writer io.Writer, as string) ( io.Writer, string,  time.Time )   {
	now := time.Now()
	str := fmt.Sprint("[",now,"]\t-> ", as, "\n")
	io.WriteString(writer, str)
	return writer, as, now
}

// Trace to a writer the leaving to a function call
func Un(writer io.Writer, as string, timestamp time.Time ) {
	now := time.Now()
	sub := now.Sub(timestamp).Nanoseconds()/1000
	str := fmt.Sprint("[",now,"]\t<- ", as, "\t\tspent:",  sub ,"ms\n")
	io.WriteString(writer, str)
}
