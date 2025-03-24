package img2json

import (
	"database/sql"
	"encoding/json"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"iter"
)

type Color struct {
	R uint32 `json:"r"`
	G uint32 `json:"g"`
	B uint32 `json:"b"`
	A uint32 `json:"a"`
}

func (c Color) ToJson() ([]byte, error) {
	return json.Marshal(&c)
}

func ColorNew(i color.Color) sql.Null[Color] {
	var ret sql.Null[Color]
	if nil != i {
		r, g, b, a := i.RGBA()
		ret.V = Color{
			R: r,
			G: g,
			B: b,
			A: a,
		}
		ret.Valid = true
	}
	return ret
}

type Row []Color

type Image struct{ image.Image }

func (i Image) Width() int  { return i.Bounds().Dx() }
func (i Image) Height() int { return i.Bounds().Dy() }

func (i Image) ToRows() iter.Seq[Row] {
	return func(yield func(Row) bool) {
		var height int = i.Height()
		var width int = i.Width()

		var buf []Color

		for y := range height {
			buf = buf[:0]
			for x := range width {
				var col color.Color = i.Image.At(x, y)
				var ncol sql.Null[Color] = ColorNew(col)
				if !ncol.Valid {
					continue
				}
				var rgba Color = ncol.V
				buf = append(buf, rgba)
			}
			if !yield(buf) {
				return
			}
		}
	}
}

func ReaderToImage(rdr io.Reader) (Image, error) {
	img, _, e := image.Decode(rdr)
	return Image{img}, e
}
