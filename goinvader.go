package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"image"
	"image/color"
	_ "image/png"
	"os"
)

var gosavior float64 = 0
var gosaviorold float64 = 0
var gosaviorold2 float64 = 0
var gosaviorold3 float64 = 0
var pydestroyer float64 = 5
var varia bool = false
var variadoi bool = false
var variatre bool = false
var altura float64 = 185
var alturadois float64 = 185
var alturatres float64 = 185
var py1 bool = true
var py2 bool = true
var py3 bool = true
var py4 bool = true

func jogo(screen *ebiten.Image) error {

	reader, err := os.Open("assets/standing.png")
	if err != nil {
		//log.Fatal(err)
	}

	defer reader.Close()

	screen.Fill(color.Black)
	m, _, err := image.Decode(reader)
	jogador, _ := ebiten.NewImageFromImage(m, ebiten.FilterNearest)
	readerpython, err := os.Open("assets/python.png")
	if err != nil {
		//log.Fatal(err)
	}

	defer readerpython.Close()

	screen.Fill(color.Black)
	p, _, err := image.Decode(readerpython)
	pyinvader, _ := ebiten.NewImageFromImage(p, ebiten.FilterNearest)
	//ebitenutil.DebugPrint(screen, "ta funfando :3")
	quadrado, _ := ebiten.NewImage(5, 5, ebiten.FilterNearest)

	quadrado.Fill(color.White)

	pyops := &ebiten.DrawImageOptions{}
	jogadorops := &ebiten.DrawImageOptions{}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {

		gosavior++
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {

		gosavior--
	}
	jogadorops.GeoM.Translate(gosavior, 185)
	screen.DrawImage(jogador, jogadorops)
	if altura <= 53 && gosaviorold <= 53 || alturadois <= 53 && gosaviorold <= 53 || alturatres <= 53 && gosaviorold <= 53 {
		py1 = false
	} else if altura <= 53 && gosaviorold <= 53 || alturadois <= 53 && gosaviorold >= 53 && gosaviorold <= 106 || alturatres <= 53 && gosaviorold >= 53 && gosaviorold <= 106 {
		py2 = false
	} else if altura <= 53 && gosaviorold <= 53 || alturadois <= 53 && gosaviorold >= 53 && gosaviorold <= 106 || alturatres <= 53 && gosaviorold >= 106 && gosaviorold <= 159 {
		py3 = false
	} else if (gosaviorold >= 159 || gosaviorold >= 159 || gosaviorold3 >= 159) && altura <= 53 || (gosaviorold >= 159 || gosaviorold >= 159 || gosaviorold3 >= 159) && alturadois <= 53 || (gosaviorold >= 159 || gosaviorold >= 159 || gosaviorold3 >= 159) && alturatres <= 53 {
		py4 = false
	} else {

		pyops.GeoM.Translate(pydestroyer, 0)
		pyops2 := &ebiten.DrawImageOptions{}
		pyops3 := &ebiten.DrawImageOptions{}
		pyops4 := &ebiten.DrawImageOptions{}
		pyops2.GeoM.Translate(pydestroyer+53, 0)
		pyops3.GeoM.Translate(pydestroyer+106, 0)
		pyops4.GeoM.Translate(pydestroyer+159, 0)
		if py1 == true {
			screen.DrawImage(pyinvader, pyops)
		}
		if py2 == true {
			screen.DrawImage(pyinvader, pyops2)
		}
		if py3 == true {
			screen.DrawImage(pyinvader, pyops3)
		}
		if py4 == true {
			screen.DrawImage(pyinvader, pyops4)
		}
	}
	quaopts := &ebiten.DrawImageOptions{}
	quaoptsdois := &ebiten.DrawImageOptions{}
	quaoptstres := &ebiten.DrawImageOptions{}
	quaopts.GeoM.Translate(gosaviorold, altura)
	quaoptsdois.GeoM.Translate(gosaviorold2, alturadois)
	quaoptstres.GeoM.Translate(gosaviorold3, alturatres)
	if varia == true {
		screen.DrawImage(quadrado, quaopts)
		altura = altura - 5
	}
	if variadoi == true {

		screen.DrawImage(quadrado, quaoptsdois)
		alturadois = alturadois - 5
	}
	if variatre == true {

		screen.DrawImage(quadrado, quaoptstres)
		alturatres = alturatres - 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) && varia == false && variadoi == false && variatre == false {
		gosaviorold = gosavior + 35
		varia = true

	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) && varia == true && variadoi == false && variatre == false {
		gosaviorold2 = gosavior + 35
		alturadois = 185
		variadoi = true

	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) && variadoi == true && varia == true {
		gosaviorold3 = gosavior + 35
		alturatres = 185
		variatre = true

	}

	if altura == 0 {
		altura = 185
		varia = false
	}
	if alturadois == 0 {
		alturadois = 185
		variadoi = false
	}
	if alturatres == 0 {
		alturatres = 185
		variatre = false
	}

	if py1 == false && py2 == false && py3 == false && py4 == false {
		ebitenutil.DebugPrint(screen, "vc ganhou")
	}
	return nil
}

func main() {

	//m, _, err := image.Decode(reader)

	//if err != nil {

	//}
	ebiten.Run(jogo, 320, 240, 2, "Go Invader")
	ebiten.IsRunningSlowly()
}
