package main

import (
	"context"
	"io"
	"iter"
	"log"
	"os"

	ij "github.com/takanoriyanagitani/go-img2json"
	ws "github.com/takanoriyanagitani/go-img2json/json/writer/std"
	. "github.com/takanoriyanagitani/go-img2json/util"
)

var input io.Reader = os.Stdin

var img IO[ij.Image] = Bind(
	Of(input),
	Lift(ij.ReaderToImage),
)

var rows IO[iter.Seq[ij.Row]] = Bind(
	img,
	Lift(func(i ij.Image) (iter.Seq[ij.Row], error) { return i.ToRows(), nil }),
)

var rows2stdout ws.WriteRows = ws.RowsToStdout

var stdin2rows2stdout IO[Void] = Bind(
	rows,
	rows2stdout,
)

func main() {
	_, e := stdin2rows2stdout(context.Background())
	if nil != e {
		log.Printf("%v\n", e)
	}
}
