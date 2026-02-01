package writer

import (
	"fmt"
	"image/color"
	"math"
	"math/rand/v2"
	"os"

	svg "github.com/twpayne/go-svg"
	"github.com/twpayne/go-svg/svgpath"
)

const (
	output_file_path string = "qr.svg"
	logoRelativeSize        = 1. / 6.
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
	// TODO: Support Minimum Quiet Zone
	quietZone := 4
	canvas := svg.New().
		WidthHeight(float64(dim), float64(dim), svg.Number).
		Transform(svg.String(fmt.Sprintf("scale(%d) translate(%d, %d)", req.Scale, quietZone, quietZone)))
	canvas.Attrs["transform-origin"] = svg.String("0 0")

	// Add center cross hair for debugging
	if req.Debug {
		canvas.AppendChildren(
			svg.Path().D(
				svgpath.New().MoveToAbs([]float64{float64(dim) / 2, 0}).LineToAbs([]float64{float64(dim) / 2, float64(dim)}).MoveToAbs([]float64{0, float64(dim) / 2}).LineToAbs([]float64{float64(dim), float64(dim) / 2}),
			).Style("stroke:red;stroke-width:0.1"),
		)
	}

	// Definitions
	circle := svg.Circle().R(svg.Number(0.5)).ID(svg.String(ShapeCircle))
	circle.Attrs["transform"] = svg.String("translate(0.5,0.5)")

	square := svg.Rect().XYWidthHeight(0, 0, 1, 1, svg.Number).ID(svg.String(ShapeSquare))

	squircle := svg.Path().D(
		svgpath.New().MoveToAbs([]float64{0, 0.5}).CurveToAbs([]float64{0, 0.125}, []float64{0.125, 0}, []float64{0.5, 0}).SCurveToAbs([]float64{1, 0.125}, []float64{1., 0.5}, []float64{0.875, 1.}, []float64{0.5, 1.}, []float64{0, 0.875}, []float64{0, 0.5}).ClosePath(),
	).ID(svg.String(ShapeSquircle))

	alignmentBackground := svg.Use().Href(svg.String("#square")).Style("fill:white")
	alignmentBackground.Attrs["transform"] = svg.String(GetTransform(req.Shape, 5.0, 0.0, 0.))
	alignmentOuterRing := svg.Use().Href(svg.String("#" + string(req.Shape))).Style(svg.String(GetStyle(req.Shape, req.Color, color.Black, 5.0)))
	alignmentOuterRing.Attrs["transform"] = svg.String(GetTransform(req.Shape, 5.0, 0.0, 0.2))
	alignmentMiddleRing := svg.Use().Href(svg.String("#" + string(req.Shape))).Style(svg.String(NoStrokeStyle(color.White, color.White)))
	alignmentMiddleRing.Attrs["transform"] = svg.String(GetTransform(req.Shape, 3.0, 1.0, 0.))
	alignmentCenterRing := svg.Use().Href(svg.String("#" + string(req.Shape))).Style(svg.String(NoStrokeStyle(req.Color, color.Black)))
	alignmentCenterRing.Attrs["transform"] = svg.String(GetTransform(req.Shape, 1.0, 2.0, 0.))
	alignmentPatternGroup := svg.G(alignmentBackground, alignmentOuterRing, alignmentMiddleRing, alignmentCenterRing).ID("alignmentpattern")

	finderBackground := svg.Use().Href(svg.String("#square")).Style("fill:white")
	finderBackground.Attrs["transform"] = svg.String(GetTransform(req.Shape, 7.0, 0.0, 0.))
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

	// Connect neighboring modules
	connect(req, canvas, dim)

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

// connect encapsulates the logic for drawing connected shapes (rectangles) based on module color.
func connect(req SVGRequest, canvas *svg.SVGElement, dim int) {
	// featEnabled is false in the original code, thus this whole function is effectively disabled.
	// We preserve the original behavior.
	featEnabled := false
	if !featEnabled {
		return
	}

	width := len(req.Cells) // Using len(req.Cells) for correct indexing into the visited map
	visited := make(map[int]struct{}, width*width)

	// Helper function to draw a connecting rectangle
	drawConnectingRect := func(x, y int, dir uint64, c color.Color) {
		if dir == 0 { // Horizontal
			canvas.AppendChildren(
				svg.Rect().XYWidthHeight(float64(x-4)+0.5, float64(y-4), 1, 1, svg.Number).Style(
					svg.String(NoStrokeStyle(req.Color, c)),
				),
			)
		} else { // Vertical
			canvas.AppendChildren(
				svg.Rect().XYWidthHeight(float64(x-4), float64(y-4)+0.5, 1, 1, svg.Number).Style(
					svg.String(NoStrokeStyle(req.Color, c)),
				),
			)
		}
	}

	// Helper function to attempt connection in a given direction
	tryConnectInDirection := func(startX, startY int, c color.Color, currentDir uint64) bool {
		foundNeighbor := false
		if currentDir == 0 { // Horizontal
			for nx := startX + 1; nx < dim && nx < width && req.Cells[startY][nx] == c; nx++ { // Original code used `nx < dim`. Preserving it.
				if _, ok := visited[startY*width+nx]; ok {
					break
				}
				foundNeighbor = true
				visited[startY*width+nx] = struct{}{}
				drawConnectingRect(nx-1, startY, currentDir, c) // nx-1 is the x-coordinate of the left module for the connection
			}
		} else { // Vertical
			for ny := startY + 1; ny < width && req.Cells[ny][startX] == c; ny++ {
				if _, ok := visited[ny*width+startX]; ok {
					break
				}
				foundNeighbor = true
				visited[ny*width+startX] = struct{}{}
				drawConnectingRect(startX, ny-1, currentDir, c) // ny-1 is the y-coordinate of the top module for the connection
			}
		}
		return foundNeighbor
	}

	for y, row := range req.Cells {
		for x, c := range row {
			if c == color.Black { // Only consider black modules
				if _, ok := visited[y*width+x]; !ok {
					visited[y*width+x] = struct{}{} // Mark current cell as visited

					dir := rand.Uint64() & 1 // Pick random initial direction (0 for horizontal, 1 for vertical)

					// Try connecting in the first direction
					if !tryConnectInDirection(x, y, c, dir) {
						// If no connection was found in the first direction, try the other direction
						tryConnectInDirection(x, y, c, 1-dir)
					}
				}
			}
		}
	}
}

func NoPaddingTransform(scale float64, pos float64) string {
	return fmt.Sprintf("scale(%f) translate(%f, %f)", scale, pos/scale, pos/scale)
}

func PaddedTransform(scale float64, pos float64, padding float64) string {
	return fmt.Sprintf("scale(%f) translate(%f, %f)", scale-padding, (pos+padding/2.)/(scale-padding), (pos+padding/2.)/(scale-padding))
}

func GetTransform(shape Shape, scale float64, pos float64, padding float64) string {
	switch shape {
	case ShapeSquare:
		return NoPaddingTransform(scale, pos)
	default:
		return PaddedTransform(scale, pos, padding)
	}
}

func GetStyle(shape Shape, target, source color.Color, scale float64) string {
	switch shape {
	case ShapeSquare:
		return NoStrokeStyle(target, source)
	default:
		return StrokeStyle(target, source, cell_gap/scale)
	}
}

func NoStrokeStyle(target, source color.Color) string {
	if source == color.Black {
		return fmt.Sprintf("fill:%s;stroke:none", ColorToFill(target))
	}
	return "fill:white;stroke:none"
}

func StrokeStyle(target, source color.Color, width float64) string {
	if source == color.Black {
		return fmt.Sprintf("fill:%s;stroke:white;stroke-width:%f", ColorToFill(target), width)
	}
	return fmt.Sprintf("fill:white;stroke:white;stroke-width:%f", width)
}

func ColorToFill(c color.Color) string {
	if c == nil {
		return "none"
	}
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("rgb(%d %d %d)", r>>8, g>>8, b>>8)
}
