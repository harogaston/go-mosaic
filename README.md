# Go Mosaic

**Go Mosaic** is a powerful Go tool designed for generating customizable and aesthetic QR codes. Beyond standard QR code generation, it offers advanced styling options, including various module shapes and logo integration, exported directly to SVG for high-quality scaling.

## Features

- **Complete QR Code Support**: standardized QR Code Model 2 (Versions 1-40).
- **Error Correction**: specific levels L (Low), M (Medium), Q (Quartile), and H (High).
- **Custom Shapes**: Style your QR codes with unique module shapes:
  - Square (Successor)
  - Circle
  - Rounded
  - Slanted
  - Squircle
- **Logo Integration**: Embed logos directly into the QR code center with automatic padding.
- **SVG Output**: High-quality vector output suitable for web and print.
- **Micro QR Support**: (Experimental) support for localized Micro QR codes.

## Installation

To install the CLI tool, clone the repository and run:

```bash
git clone https://github.com/harogaston/go-mosaic.git
cd go-mosaic
go run . --help
```

## CLI Usage

Generate QR codes directly from your terminal.

### Basic Usage

```bash
go run main.go -data "Hello, World!"
```

This will generate a `qr.svg` file in the current directory.

### Options

| Flag       | Description                                              | Default                    |
| :--------- | :------------------------------------------------------- | :------------------------- |
| `-data`    | The string data to encode.                               | `"01234567"`               |
| `-shape`   | Module shape: `square`, `circle`, `rounded`, `slanted`, `squircle` | `square`                   |
| `-level`   | Error correction level: `L`, `M`, `Q`, `H`.              | `L`                        |
| `-logo`    | Include a logo (bool). Uses default resources if true.   | `false`                    |
| `-version` | Fixed QR version (1-40). 0 for auto-detection.           | `0`                        |
| `-micro`   | Generate Micro QR code (experimental).                   | `false`                    |
| `-debug`   | Enable debug output and patterns.                        | `false`                    |

### Examples

**Circular Modules with Medium Error Correction:**

```bash
go run main.go -data "https://example.com" -shape circle -level M
```

**Rounded Modules with Logo:**

```bash
go run main.go -data "Go Mosaic" -shape rounded -logo "path/to/logo.png"
```

*Note: Enabling the logo option automatically sets the error correction level to 'H' to ensure decodability.*

## Documentation

For a deep dive into the technical details of QR code structure, encoding procedures, and standards implementation, please refer to the [QR Specification Summary](docs/QR_SPECIFICATION.md).
