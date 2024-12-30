package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	fontSizeXXL   = 256
	fontSizeXL    = 128
	fontSizeL     = 64
	fontSizeM     = 32
	fontSizeS     = 16
	SPLASH_FIRST  = 5
	SPLASH_SECOND = 10
	SPLASH_THIRD  = 15
	SPLASH_FOURTH = 20
	SPLASH_END    = 30
)

type txt struct {
	text string
	font rl.Font
	pos  rl.Vector2
	size float32
}

func (t txt) Size() rl.Vector2 {
	return rl.MeasureTextEx(t.font, t.text, t.size, 0)
}

// globals
var splashTitle txt
var splashSubTitle1 txt
var splashSubTitle2 txt
var cancelTxt txt

var music rl.Music
var font rl.Font
var font_mecha rl.Font

var tileSize int = 16
var tileSizeI32 int32 = int32(tileSize)
var tileSizeF32 float32 = float32(tileSize)

// grid of rectangles making up the background
var grid = [][]rl.Rectangle{}

var field rl.Texture2D

// array of grass rectangles that can be applied to a texture
var fieldRects = []rl.Rectangle{}

func main() {
	// init window
	rl.InitWindow(tileSizeI32*100, tileSizeI32*50, "Vendel")
	defer rl.CloseWindow()
	// init audio
	rl.InitAudioDevice()
	music = rl.LoadMusicStream("resources/sound/dark.mp3")
	rl.PlayMusicStream(music)

	width_f32 := float32(rl.GetScreenWidth())
	height_f32 := float32(rl.GetScreenHeight())

	// load resources
	font = rl.LoadFontEx("resources/fonts/gothical.ttf", fontSizeXXL, nil, 250)
	font_mecha = rl.LoadFontEx("resources/fonts/mecha.ttf", fontSizeXL, nil, 250)
	field = rl.LoadTexture("resources/textures/tilesets/field_green.png")

	// build grid of tiles
	for y := 0; y < rl.GetScreenHeight(); y += tileSize {
		var row []rl.Rectangle
		for x := 0; x < rl.GetScreenWidth(); x += tileSize {
			row = append(row, rl.NewRectangle(float32(x), float32(y), tileSizeF32, tileSizeF32))
		}
		grid = append(grid, row)
	}

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			fieldRects = append(fieldRects, rl.NewRectangle(tileSizeF32*float32(x), tileSizeF32*float32(y), tileSizeF32, tileSizeF32))
		}
	}

	// splash screen texts
	splashTitle = txt{
		"Vendel Period",
		font,
		rl.NewVector2(0, 0),
		fontSizeXXL,
	}
	splashTitle.pos = rl.NewVector2(width_f32/2-splashTitle.Size().X/2, height_f32/2-(2*splashTitle.Size().Y/2))

	splashSubTitle1 = txt{
		"Scandinavia, a few dacades after ",
		font,
		rl.NewVector2(0, 0),
		fontSizeL,
	}
	splashSubTitle1.pos = rl.NewVector2(width_f32/2-splashSubTitle1.Size().X/2, height_f32/2+(2*splashSubTitle1.Size().Y/2))

	splashSubTitle2 = txt{
		"the collaps of the roman empire",
		font,
		rl.NewVector2(0, 0),
		fontSizeL,
	}
	splashSubTitle2.pos = rl.NewVector2(width_f32/2-splashSubTitle2.Size().X/2, height_f32/2+(4*splashSubTitle2.Size().Y/2))

	cancelTxt = txt{
		"PRESS ESC TO EXIT",
		font_mecha,
		rl.NewVector2(0, 0),
		fontSizeS,
	}
	cancelTxt.pos = rl.NewVector2(width_f32-cancelTxt.Size().X, 0)

	// Define the camera to look into our 3d world
	camera := rl.Camera{}
	camera.Position = rl.NewVector3(5.0, 5.0, 5.0) // Camera position
	camera.Target = rl.NewVector3(0.0, 2.0, 0.0)   // Camera looking at point
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)       // Camera up vector (rotation towards target)
	camera.Fovy = 45.0                             // Camera field-of-view Y
	camera.Projection = rl.CameraPerspective       // Camera mode type

	rl.DisableCursor() // Limit cursor to relative movement inside the window

	rl.SetTargetFPS(60)

	splashSpeed := 1.0
	splashDone := false

	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream(music)
		rl.BeginDrawing()
		if !splashDone {
			splashDone = drawSplash(&splashSpeed)
		} else {
			drawBackground(&field)
		}

		rl.ClearBackground(rl.Black)
		rl.EndDrawing()
	}
}

// TODO: apply WFC
func drawBackground(field *rl.Texture2D) {
	for y := 0; y < len(grid); y++ {
		var i int = 0
		for x := 0; x < len(grid[y]); x++ {
			rl.DrawTextureRec(*field, fieldRects[i], rl.NewVector2(grid[y][x].X, grid[y][x].Y), rl.White)
			if i < len(fieldRects)-1 {
				i++
			} else {
				i = 0
			}
		}
	}

	rl.DrawTextEx(cancelTxt.font, cancelTxt.text, cancelTxt.pos, cancelTxt.size, 0, rl.NewColor(190, 190, 190, 255))
}

func drawSplash(splashSpeed *float64) bool {

	time := rl.GetTime() * *splashSpeed

	if time > SPLASH_END {
		rl.StopMusicStream(music)
		return true
	}

	if rl.IsKeyDown(rl.KeySpace) {
		*splashSpeed++
	}

	rl.DrawTextEx(splashTitle.font, splashTitle.text, splashTitle.pos, splashTitle.size, 0, rl.NewColor(190, 33, 55, 255))
	rl.DrawTextEx(cancelTxt.font, cancelTxt.text, cancelTxt.pos, cancelTxt.size, 0, rl.NewColor(190, 190, 190, 255))

	if time > SPLASH_FIRST && time <= SPLASH_SECOND {
		rl.DrawTextEx(splashSubTitle1.font, splashSubTitle1.text, splashSubTitle1.pos, splashSubTitle1.size, 0, rl.NewColor(190, 33, 55, uint8(255*((time-SPLASH_FIRST)/SPLASH_FIRST))))
		rl.DrawTextEx(splashSubTitle2.font, splashSubTitle2.text, splashSubTitle2.pos, splashSubTitle2.size, 0, rl.NewColor(190, 33, 55, uint8(255*((time-SPLASH_FIRST)/SPLASH_FIRST))))
	} else if time >= SPLASH_SECOND && time <= SPLASH_THIRD {
		rl.DrawTextEx(splashSubTitle1.font, splashSubTitle1.text, splashSubTitle1.pos, splashSubTitle1.size, 0, rl.NewColor(190, 33, 55, 255))
		rl.DrawTextEx(splashSubTitle2.font, splashSubTitle2.text, splashSubTitle2.pos, splashSubTitle2.size, 0, rl.NewColor(190, 33, 55, 255))
	} else if time > SPLASH_THIRD && time < SPLASH_FOURTH {
		transp := uint8(255 - (time-SPLASH_THIRD)*51)
		rl.DrawTextEx(splashSubTitle1.font, splashSubTitle1.text, splashSubTitle1.pos, splashSubTitle1.size, 0, rl.NewColor(190, 33, 55, transp))
		rl.DrawTextEx(splashSubTitle2.font, splashSubTitle2.text, splashSubTitle2.pos, splashSubTitle2.size, 0, rl.NewColor(190, 33, 55, transp))
	}
	return false
}
