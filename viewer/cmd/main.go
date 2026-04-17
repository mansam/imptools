package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	TileWidth  = 56
	TileHeight = 27
)

var DebugMode string

func init() {
	flag.StringVar(&DebugMode, "debug", "", "enable debug mode")
	flag.Parse()
}

type Viewer struct {
	Surface Surface
	Tiles   [][]Tile
	Tileset map[uint8]*ebiten.Image
	Alpha   bool
}

func (r *Viewer) Update() (err error) {
	return
}

func (r *Viewer) Draw(screen *ebiten.Image) {
	if !r.Alpha {
		r.calculateAlpha()
	}

	for h := range len(r.Tiles) {
		for w := range len(r.Tiles[h]) {
			//if h > 5 {
			//	break
			//}
			//tileX := -1
			//tileY := -1
			tileX := (h - w - 1)
			tileY := (-w - 1)
			fmt.Printf("(%d,%d), ", tileX, tileY)
			//if w > 4 {
			//	return
			//}
			//
			//	}
			//}

			//h and w specified in tile coordinates
			//for h, _ := range r.Tiles {
			//	for w, _ := range r.Tiles[h] {
			//if h > 0 {
			//	continue
			//}
			//fw := float64(w)
			//fh := float64(h)
			//fw := w
			//fh := h
			//idx := h*r.Surface.Width + w
			//idx := w*r.Surface.Height + h
			//tile := r.Tiles[len(r.Tiles)-(h+1)][len(r.Tiles[0])-(w+1)]
			tile := r.Tiles[h][w]

			////
			//TileWidth := uint16(bounds.Dx())
			//TileHeight := uint16(bounds.Dy())
			//tile = r.Tiles[len(r.Tiles)-int(idx)-1]
			//
			//screenW := screen.Bounds().Dx()
			//screenH := screen.Bounds().Dy()

			img := r.Tileset[tile.Number]
			bounds := img.Bounds()
			//if bounds.Dy() > 27 {
			//	println(bounds.Dy())
			//}

			screenX := (TileWidth / 2) * (tileX + tileY)
			screenY := (TileHeight/2)*(tileY-tileX) - (2 * w)
			fmt.Printf("screen (%d,%d), ", screenX, screenY)
			//
			//screenX := (tileX - tileY) * (TileWidth / 2)
			//screenY := (tileX + tileY) * (TileHeight / 2)
			//screenY := (TileHeight * tileX / 2) //+ (TileHeight * w / 2)
			//screenX := (TileWidth * tileY)
			if h > 0 && bounds.Dx() > 28 {
				//screenX -= TileWidth + bounds.Dx()
			}
			screenX -= h * TileWidth / 2
			if h%2 == 1 {
				screenX += TileWidth / 2
			}
			//(screenW / 2) + ((TileWidth * int(w) / 2) - (TileWidth * int(h) / 2))
			//if w%2 == 1 {
			//	screenX += TileWidth //(screenW / 2) + ((TileWidth * int(w) / 2) - (TileWidth * int(h) / 2))
			//}
			//if tileY%2 == 1 {
			//	screenX -= TileWidth // / 2
			//}
			//if tileX%2 == 1 {
			//	screenY += TileHeight / 2
			//}
			//if bounds.Dy() > 30 {
			//	screenY -= (bounds.Dy()/2 - TileHeight/2)
			//}

			//screenx := (TileWidth * fw / 2) + (TileWidth * r.Surface.Height / 2) - (TileWidth * fh / 2) //- (fw)
			//screeny := (((r.Surface.Height - fh - 1) * TileHeight) / 2) + (r.Surface.Width * TileHeight / 2) - (TileHeight * fw / 2)
			//if h+w != 0 {
			//	//this worked when rendering bottom up to get everything aligned
			//	screeny += (fw * 2)
			//	screenx += (fh * 2)
			//}
			//if r.Debug {
			//	tile = r.Tiles[len(r.Tiles)-int(idx)-1]
			//	screenx = 2048 - (TileWidth * fw / 2) + (TileWidth * r.Surface.Height / 2) - (TileWidth * fh / 2) //- (fw)
			//	screeny = 2048 - (((r.Surface.Height - fh - 1) * TileHeight) / 2) + (r.Surface.Width * TileHeight / 2) - (TileHeight * fw / 2)
			//	if h+w != 0 {
			//		screeny -= (fw + fh)
			//		screenx += (fh + fw)
			//	}
			//}

			//img := r.Tileset[tile.Number]
			//bounds := img.Bounds()
			//if bounds.Dx() > TileWidth {
			//	scree
			//}

			//screeny = math.Ceil(screeny)
			//if int(screenx)%2 == 1 {
			//	screenx -= 1
			//}
			//if int(screeny)%2 == 1 {
			//	screeny += 1
			//}
			//fmt.Println(screenx, screeny)

			if tile.Empty() {
				continue
			}
			if tile.Flags != 0 {
				continue
			}

			options := &ebiten.DrawImageOptions{}
			options.GeoM.Translate(float64(screenX*-1), float64(screenY*-1))
			//debug := ebiten.NewImageFromImage(img)
			//ebitenutil.DebugPrint(debug, fmt.Sprintf("(%d,%d)", tileX, tileY))
			screen.DrawImage(img, options)
		}
	}
}

func (r *Viewer) Layout(w int, h int) (sw int, sh int) {
	return w, h
}

type Surface struct {
	Height uint16
	Width  uint16
}

func (s *Surface) W() float64 {
	return float64(s.Width)
}

func (s *Surface) H() float64 {
	return float64(s.Height)
}

func (s *Surface) String() string {
	return fmt.Sprintf("(%2d, %2d)", s.Height, s.Width)
}

type tile struct {
	Number uint8
	Flags  uint8
}

type Tile struct {
	tile
	X int
	Y int
}

func (r *Tile) Empty() bool {
	return r.Number+41 >= 255
}

func (t *Tile) String() string {
	return fmt.Sprintf("%03d.png", t.Number)
}

func ReadTileEntry(f *os.File) (t tile) {
	_ = binary.Read(f, binary.LittleEndian, &t)
	return
}

func ReadSurfaceEntry(f *os.File) (s Surface) {
	_ = binary.Read(f, binary.LittleEndian, &s)
	return
}

func (r *Viewer) LoadMap(file string) {
	mapFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer mapFile.Close()
	counter := make(map[uint8]int)
	r.Tiles = [][]Tile{}
	r.Surface = ReadSurfaceEntry(mapFile)
	fmt.Println(r.Surface.String())
	var i, j int
	for i = 0; i < 65; i++ {
		r.Tiles = append(r.Tiles, []Tile{})
		for j = 0; j < 65; j++ {
			rt := ReadTileEntry(mapFile)
			t := Tile{tile: rt}
			t.X = i - j - 1
			t.Y = -j - 1
			// why is the offset 41?!
			t.Number = t.Number - 41
			//if t.Flags != 0 {
			//	fmt.Printf("%3d,", t.Flags)
			//} else {
			//}
			counter[t.Number]++
			r.Tiles[i] = append(r.Tiles[i], t)
		}
		fmt.Println()
	}
	//fmt.Println()
	//fmt.Println()
	//fmt.Println()
	//fmt.Println()
	//fmt.Println()
	//
	//slices.Reverse(r.Tiles)
	//fmt.Printf("%v", r.Tiles)
	//for h, _ := range r.Tiles {
	//	for w, _ := range r.Tiles[h] {
	//		t := r.Tiles[h][w]
	//		fmt.Printf("%3d,", t.Number)
	//	}
	//	fmt.Println()
	//
	//}
	//for _, s := range r.Tiles {
	//	slices.Reverse(s)
	//}
	//fmt.Println(counter)
}

func (r *Viewer) LoadTileset(dir string) {
	r.Tileset = make(map[uint8]*ebiten.Image)
	entries, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		if strings.HasSuffix(strings.ToLower(entry.Name()), ".png") {
			//fmt.Printf("Load tile %s\n", entry.Name())
			r.loadTileImage(path.Join(dir, entry.Name()))
		}
	}
	fmt.Printf("Loaded %d tiles.", len(r.Tileset))
}

func (r *Viewer) loadTileImage(png string) {
	f, err := os.Open(png)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	basename := path.Base(png)
	name := strings.Split(basename, ".")[0]
	id, err := strconv.Atoi(name)
	if err != nil {
		panic(err)
	}
	ebitenImg := ebiten.NewImageFromImage(img)

	r.Tileset[uint8(id)] = ebitenImg
}

func (r *Viewer) calculateAlpha() {
	for _, v := range r.Tileset {
		r.calculateTileAlpha(v)
	}
	r.Alpha = true
}

func (r *Viewer) calculateTileAlpha(img *ebiten.Image) {
	bounds := img.Bounds()
	alpha := img.At(0, 0)
	ar, ag, ab, _ := alpha.RGBA()
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			pixel := img.At(x, y)
			r, g, b, _ := pixel.RGBA()
			if r == ar && g == ag && b == ab {
				img.Set(x, y, color.RGBA{R: uint8(0), G: uint8(0), B: uint8(0), A: 0x0})
			} else {
				img.Set(x, y, color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 0xFF})
			}
		}
	}
	switch DebugMode {
	case "b":
		ebitenutil.DebugPrint(img, fmt.Sprintf("(%d,%d)", img.Bounds().Dx(), img.Bounds().Dy()))
	}
}

func main() {
	tilesetPath := os.Args[2]
	mapPath := os.Args[1]

	viewer := &Viewer{}
	viewer.LoadMap(mapPath)
	viewer.LoadTileset(tilesetPath)
	ebiten.SetWindowSize(2048, 2048)
	ebiten.SetWindowTitle("Surface Viewer")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	err := ebiten.RunGame(viewer)
	if err != nil {
		panic(err)
	}
}
