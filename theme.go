// MIT License

// Copyright (c) 2021 Tree Xie

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package charts

import (
	"strings"

	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

var hiddenColor = drawing.Color{R: 110, G: 112, B: 121, A: 0}

var AxisColorLight = drawing.Color{
	R: 110,
	G: 112,
	B: 121,
	A: 255,
}
var AxisColorDark = drawing.Color{
	R: 185,
	G: 184,
	B: 206,
	A: 255,
}

var BackgroundColorDark = drawing.Color{
	R: 16,
	G: 12,
	B: 42,
	A: 255,
}

var TextColorDark = drawing.Color{
	R: 204,
	G: 204,
	B: 204,
	A: 255,
}

func getAxisColor(theme string) drawing.Color {
	if theme == ThemeDark {
		return AxisColorDark
	}
	return AxisColorLight
}

func getGridColor(theme string) drawing.Color {
	if theme == ThemeDark {
		return drawing.Color{
			R: 72,
			G: 71,
			B: 83,
			A: 255,
		}
	}
	return drawing.Color{
		R: 224,
		G: 230,
		B: 241,
		A: 255,
	}
}

var SeriesColorsLight = []drawing.Color{
	{
		R: 84,
		G: 112,
		B: 198,
		A: 255,
	},
	{
		R: 145,
		G: 204,
		B: 117,
		A: 255,
	},
	{
		R: 250,
		G: 200,
		B: 88,
		A: 255,
	},
	{
		R: 238,
		G: 102,
		B: 102,
		A: 255,
	},
	{
		R: 115,
		G: 192,
		B: 222,
		A: 255,
	},
}

func getBackgroundColor(theme string) drawing.Color {
	if theme == ThemeDark {
		return BackgroundColorDark
	}
	return chart.DefaultBackgroundColor
}

func getTextColor(theme string) drawing.Color {
	if theme == ThemeDark {
		return TextColorDark
	}
	return chart.DefaultTextColor
}

type ThemeColorPalette struct {
	Theme string
}

func (tp ThemeColorPalette) BackgroundColor() drawing.Color {
	return getBackgroundColor(tp.Theme)
}

func (tp ThemeColorPalette) BackgroundStrokeColor() drawing.Color {
	return chart.DefaultBackgroundStrokeColor
}

func (tp ThemeColorPalette) CanvasColor() drawing.Color {
	if tp.Theme == ThemeDark {
		return BackgroundColorDark
	}
	return chart.DefaultCanvasColor
}

func (tp ThemeColorPalette) CanvasStrokeColor() drawing.Color {
	return chart.DefaultCanvasStrokeColor
}

func (tp ThemeColorPalette) AxisStrokeColor() drawing.Color {
	if tp.Theme == ThemeDark {
		return BackgroundColorDark
	}
	return chart.DefaultAxisColor
}

func (tp ThemeColorPalette) TextColor() drawing.Color {
	return getTextColor(tp.Theme)
}

func (tp ThemeColorPalette) GetSeriesColor(index int) drawing.Color {
	return getSeriesColor(tp.Theme, index)
}

func getSeriesColor(theme string, index int) drawing.Color {
	return SeriesColorsLight[index%len(SeriesColorsLight)]
}

func parseColor(color string) drawing.Color {
	if color == "" {
		return drawing.Color{}
	}
	if strings.HasPrefix(color, "#") {
		return drawing.ColorFromHex(color[1:])
	}
	// TODO rgba
	return drawing.Color{}
}
