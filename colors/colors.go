// Package colors provides default color values and a convert method for colors
// Spice up your embed life!
package colors

// Color is a basic struct for RGB Color values
type Color struct {
	r int
	g int
	b int
}

// Red is 0xFF0000
func Red() *Color {
	return &Color{r: 255}
}

// Green is 0x00FF00
func Green() *Color {
	return &Color{g: 255}
}

// Blue is 0x0000FF
func Blue() *Color {
	return &Color{b: 255}
}

// Black is 0x000000
func Black() *Color {
	return &Color{0, 0, 0}
}

// White is 0xFFFFFF
func White() *Color {
	return &Color{255, 255, 255}
}

// Pink is 0xFF1493
func Pink() *Color {
	return &Color{255, 20, 147}
}

// Orange is 0xFFA500
func Orange() *Color {
	return &Color{255, 165, 0}
}

// Brown is 0xA52A2A
func Brown() *Color {
	return &Color{165, 42, 42}
}

// Yellow is 0xFFFF00
func Yellow() *Color {
	return &Color{255, 255, 0}
}

// Cyan is 0x00FFFF
func Cyan() *Color {
	return &Color{0, 255, 255}
}

// Magenta is 0xFF00FF
func Magenta() *Color {
	return &Color{255, 0, 255}
}

// Convert converts the RGB struct to a single RGB integer
func (c Color) Convert() int {
	return (c.r * 256 * 256) + (c.g * 256) + c.b
}
