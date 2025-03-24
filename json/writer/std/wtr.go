package writer

import (
	"context"
	"encoding/json"
	"io"
	"iter"
	"os"

	ij "github.com/takanoriyanagitani/go-img2json"
	. "github.com/takanoriyanagitani/go-img2json/util"
)

type WriteRows func(iter.Seq[ij.Row]) IO[Void]

func WriterToWriteRows(wtr io.Writer) WriteRows {
	return func(rows iter.Seq[ij.Row]) IO[Void] {
		return func(ctx context.Context) (Void, error) {
			var enc *json.Encoder = json.NewEncoder(wtr)
			for row := range rows {
				select {
				case <-ctx.Done():
					return Empty, ctx.Err()
				default:
				}

				var cols []ij.Color = row
				e := enc.Encode(cols)
				if nil != e {
					return Empty, e
				}
			}
			return Empty, nil
		}
	}
}

var RowsToStdout WriteRows = WriterToWriteRows(os.Stdout)
