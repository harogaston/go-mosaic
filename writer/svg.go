package writer

import (
	"fmt"
	"image/color"
	"math"
	"os"

	svg "github.com/twpayne/go-svg"
	"github.com/twpayne/go-svg/svgpath"
)

const (
	output_file_path string = "qr.svg"
	logoRelativeSize        = 2. / 7.
	cell_gap                = 0.125
	logoBorderWidth         = 0.4
)

type SVGRequest struct {
	Scale             int
	Cells             [][]color.Color
	AlignmentPatterns [][]int
	Shape             Shape
	Logo              string
	Color             color.Color
	Debug             bool
}

// GenerateRoundedSquare returns an SVG path string for a 1x1 square
// with a corner radius 'r'.
// r should be between 0 (sharp square) and 0.5 (circle).
func GenerateRoundedSquare(r float64) *svgpath.Path {
	// Clamp r to valid range [0, 0.5]
	if r > 0.5 {
		r = 0.5
	} else if r < 0 {
		r = 0
	}

	// k (kappa) is the magic constant to approximate a circular arc
	// with a cubic Bézier curve: (4/3)*(sqrt(2)-1)
	const k = 0.552284749831

	off := r * (1 - k)

	return svgpath.New().MoveToAbs([]float64{0, 0.5}).
		VLineToAbs(r).CurveToAbs([]float64{0, off}, []float64{off, 0}, []float64{r, 0}).
		HLineToAbs(1-r).CurveToAbs([]float64{1 - off, 0}, []float64{1, off}, []float64{1, r}).
		VLineToAbs(1-r).CurveToAbs([]float64{1, 1 - off}, []float64{1 - off, 1}, []float64{1 - r, 1}).
		HLineToAbs(r).CurveToAbs([]float64{off, 1}, []float64{0, 1 - off}, []float64{0, 1 - r}).ClosePath()
}

// GenerateSquircle returns an SVG path string for a 1x1 squircle
// with a curviness parameter.
// curviness should be between 0 (sharp square) and 0.5 (circle).
func GenerateSquircle(curviness float64) *svgpath.Path {
	// Clamp curviness to valid range [0, 0.5]
	if curviness > 0.5 {
		curviness = 0.5
	} else if curviness < 0 {
		curviness = 0
	}

	pull := curviness
	invPull := 1.0 - pull

	return svgpath.New().
		MoveToAbs([]float64{0, 0.5}).
		CurveToAbs([]float64{0, pull}, []float64{pull, 0}, []float64{0.5, 0}).
		SCurveToAbs(
			[]float64{1, pull}, []float64{1.0, 0.5},
			[]float64{invPull, 1}, []float64{0.5, 1.0},
			[]float64{0, invPull}, []float64{0, 0.5},
		).
		ClosePath()
}

func WriteSVG(req SVGRequest) {
	if req.Color == nil {
		req.Color = color.Black
	}

	file, err := os.Create(output_file_path)
	if err != nil {
		fmt.Println("Error creating SVG file:", err)
	}
	defer file.Close()

	dim := len(req.Cells)
	quietZone := 4
	canvas := svg.New().
		WidthHeight(float64(dim), float64(dim), svg.Number).
		Transform(svg.String(fmt.Sprintf("scale(%d) translate(%d, %d)", req.Scale, quietZone, quietZone)))
	canvas.Attrs["transform-origin"] = svg.String("0 0")

	// Definitions
	circle := svg.Circle().R(svg.Number(0.5)).ID(svg.String(ShapeCircle))
	circle.Attrs["transform"] = svg.String("translate(0.5,0.5)")
	square := svg.Rect().XYWidthHeight(0, 0, 1, 1, svg.Number).ID(svg.String(ShapeSquare))
	rounded := svg.Path().D(GenerateRoundedSquare(0.35)).ID(svg.String(ShapeRounded))
	squircle := svg.Path().D(GenerateSquircle(0.125)).ID(svg.String(ShapeSquircle))

	alignmentBackground := svg.Use().Href(svg.String("#square")).Style("fill:white")
	alignmentBackground.Attrs["transform"] = svg.String(GetTransform(ShapeSquare, 5.0, 0.0, 0.))
	alignmentOuterRing := svg.Use().Href(svg.String("#" + string(req.Shape))).Style(svg.String(GetStyle(req.Shape, req.Color, color.Black, 5.0)))
	alignmentOuterRing.Attrs["transform"] = svg.String(GetTransform(req.Shape, 5.0, 0.0, 0.2))
	alignmentMiddleRing := svg.Use().Href(svg.String("#" + string(req.Shape))).Style(svg.String(NoStrokeStyle(color.White, color.White)))
	alignmentMiddleRing.Attrs["transform"] = svg.String(GetTransform(req.Shape, 3.0, 1.0, 0.))
	alignmentCenterRing := svg.Use().Href(svg.String("#" + string(req.Shape))).Style(svg.String(NoStrokeStyle(req.Color, color.Black)))
	alignmentCenterRing.Attrs["transform"] = svg.String(GetTransform(req.Shape, 1.0, 2.0, 0.))
	alignmentPatternGroup := svg.G(alignmentBackground, alignmentOuterRing, alignmentMiddleRing, alignmentCenterRing).ID("alignmentpattern")

	finderBackground := svg.Use().Href(svg.String("#square")).Style("fill:white")
	finderBackground.Attrs["transform"] = svg.String(GetTransform(ShapeSquare, 7.0, 0.0, 0.))
	finderOuterRing := svg.Use().Href(svg.String("#" + string(req.Shape))).Style(svg.String(GetStyle(req.Shape, req.Color, color.Black, 7.)))
	finderOuterRing.Attrs["transform"] = svg.String(GetTransform(req.Shape, 7.0, 0.0, 0.2))
	finderMiddleRing := svg.Use().Href(svg.String("#" + string(req.Shape))).Style(svg.String(NoStrokeStyle(color.White, color.White)))
	finderMiddleRing.Attrs["transform"] = svg.String(GetTransform(req.Shape, 5.0, 1.0, 0.))
	finderCenterRing := svg.Use().Href(svg.String("#" + string(req.Shape))).Style(svg.String(NoStrokeStyle(req.Color, color.Black)))
	finderCenterRing.Attrs["transform"] = svg.String(GetTransform(req.Shape, 3.0, 2.0, 0.))
	finderPatternGroup := svg.G(finderBackground, finderOuterRing, finderMiddleRing, finderCenterRing).ID("finderpattern")

	canvas.AppendChildren(
		svg.Defs(
			circle,
			square,
			rounded,
			squircle,
			finderPatternGroup,
			alignmentPatternGroup,
		),
	)

	// Draw Modules
	for y, row := range req.Cells {
		for x, c := range row {
			if c == color.Black {
				canvas.AppendChildren(
					svg.Use().XY(float64(x), float64(y), svg.Number).Href(svg.String(fmt.Sprintf("#%s", req.Shape))).Style(
						svg.String(GetStyle(req.Shape, req.Color, c, 1.0)),
					),
				)
			}
		}
	}

	// Superimpose alignment patterns (for versions >= 2)
	if dim >= 21+4 {
		for _, ap := range req.AlignmentPatterns {
			canvas.AppendChildren(
				svg.Use().Href(svg.String("#alignmentpattern")).XY(float64(ap[0]-2), float64(ap[1]-2), svg.Number),
			)
		}
	}

	// Superimpose finder patterns
	canvas.AppendChildren(
		svg.Use().Href(svg.String("#finderpattern")).XY(0, 0, svg.Number),
		svg.Use().Href(svg.String("#finderpattern")).XY(float64(dim-7), 0, svg.Number),
		svg.Use().Href(svg.String("#finderpattern")).XY(0, float64(dim-7), svg.Number),
	)

	// Ensure logo size is always odd
	logoSize := int(math.Floor(float64(dim) * logoRelativeSize))
	logoSize += (logoSize + 1) % 2

	logoPos := dim/2. - logoSize/2.

	logoCenter := float64(logoPos) + float64(logoSize)/2.

	// Draw logo ensuring a minimum size of 5 modules
	if req.Logo != "" && logoSize >= 5 {

		// Create safe zone around logo
		var padding float64
		if req.Shape != ShapeSquare {
			padding = 2.0
		} else {
			padding = 1.0
		}
		startCell := logoPos - 1
		endCell := logoPos + logoSize + 1

		radius := float64(logoSize)/2. + padding
		for y := startCell; y < endCell; y++ {
			for x := startCell; x < endCell; x++ {
				dx := float64(x) + .5 - logoCenter
				dy := float64(y) + .5 - logoCenter
				distance := dx*dx + dy*dy
				if distance < radius*radius {
					canvas.AppendChildren(
						svg.Use().XY(float64(x), float64(y), svg.Number).Href("#square").Style("fill:white"),
					)
				}
			}
		}

		// Place logo with clipping path
		logoClipPath := svg.Use().Href(svg.String("#" + string(req.Shape)))
		logoClipPath.Attrs["transform"] = svg.String(GetTransform(req.Shape, float64(logoSize), float64(logoPos), 0.))

		logoBorderScale := float64(logoSize) + 1.0
		logoBorderPos := float64(logoPos) - 0.5
		logoBorder := svg.Use().Href(svg.String("#" + string(req.Shape))).Style(svg.String(fmt.Sprintf("fill:none;stroke:%s;stroke-width:%f", ColorToFill(req.Color), logoBorderWidth/logoBorderScale)))
		logoBorder.Attrs["transform"] = svg.String(GetTransform(req.Shape, logoBorderScale, logoBorderPos, 0.))
		if req.Shape == ShapeSquare {
			logoBorder.Visibility(svg.String("hidden"))
		}

		canvas.AppendChildren(
			svg.ClipPath(logoClipPath).ID("logoClip"),
			logoBorder,
			svg.Image().Href(svg.String(req.Logo)).XYWidthHeight(
				float64(logoPos), float64(logoPos), float64(logoSize), float64(logoSize), svg.Number,
			).ClipPath("url(#logoClip)"),
		)
	}

	// Write to file
	if _, err := canvas.WriteToIndent(file, "", "  "); err != nil {
		panic(err)
	}
}

func GetTransform(shape Shape, scale float64, pos float64, padding float64) string {
	switch shape {
	case ShapeSquare:
		padding = 0
	}
	return fmt.Sprintf("scale(%f) translate(%f, %f)", scale-padding, (pos+padding/2.)/(scale-padding), (pos+padding/2.)/(scale-padding))
}

func GetStyle(shape Shape, target, source color.Color, scale float64) string {
	var fill color.Color
	fill = color.White
	if source == color.Black {
		fill = target
	}

	var width float64
	if scale != 0 {
		width = cell_gap / scale
	}

	switch shape {
	case ShapeSquare:
		return fmt.Sprintf("fill:%s;stroke:none", ColorToFill(fill))
	default:
		return fmt.Sprintf("fill:%s;stroke:white;stroke-width:%f", ColorToFill(fill), width)
	}
}

func NoStrokeStyle(target, source color.Color) string {
	if source == color.Black {
		return fmt.Sprintf("fill:%s;stroke:none", ColorToFill(target))
	}
	return "fill:white;stroke:none"
}

func ColorToFill(c color.Color) string {
	if c == nil {
		return "none"
	}
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("rgb(%d %d %d)", r>>8, g>>8, b>>8)
}
