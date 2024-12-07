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

var splashTitle txt
var splashSubTitle1 txt
var splashSubTitle2 txt
var cancelTxt txt

var music rl.Music
var font rl.Font
var font_mecha rl.Font

func main() {
	rl.InitWindow(1280, 720, "Vendel")
	defer rl.CloseWindow()

	width_f32 := float32(rl.GetScreenWidth())
	height_f32 := float32(rl.GetScreenHeight())

	rl.InitAudioDevice()
	font = rl.LoadFontEx("resources/fonts/gothical.ttf", fontSizeXXL, nil, 250)
	font_mecha = rl.LoadFontEx("resources/fonts/mecha.ttf", fontSizeXL, nil, 250)

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
		"PRESS ESC TO CANCEL, SPACE TO FAST FORWARD",
		font_mecha,
		rl.NewVector2(0, 0),
		fontSizeS,
	}
	cancelTxt.pos = rl.NewVector2(width_f32/2-cancelTxt.Size().X/2, height_f32-cancelTxt.Size().Y)

	// Define the camera to look into our 3d world
	camera := rl.Camera{}
	camera.Position = rl.NewVector3(5.0, 5.0, 5.0) // Camera position
	camera.Target = rl.NewVector3(0.0, 2.0, 0.0)   // Camera looking at point
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)       // Camera up vector (rotation towards target)
	camera.Fovy = 45.0                             // Camera field-of-view Y
	camera.Projection = rl.CameraPerspective       // Camera mode type

	rl.DisableCursor() // Limit cursor to relative movement inside the window

	rl.SetTargetFPS(60)

	music = rl.LoadMusicStream("resources/sound/dark.mp3")
	rl.PlayMusicStream(music)

	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream(music)
		rl.BeginDrawing()

		drawSplash()

		rl.ClearBackground(rl.Black)
		rl.EndDrawing()
	}
}

var splashSpeed = 1.0

func drawSplash() {

	time := rl.GetTime() * splashSpeed

	if time > SPLASH_END {
		rl.StopMusicStream(music)
		return
	}

	if rl.IsKeyDown(rl.KeySpace) {
		splashSpeed++
	}

	rl.ClearBackground(rl.Black)
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
}
