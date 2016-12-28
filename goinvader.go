package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"image"
	"image/color"
	_ "image/png"
	"os"
)

var (
	Gosavior    float64 = 0
	gosaviorold float64 = 0
	pydestroyer float64 = 5
	varia       bool    = false
	variadoi    bool    = false
	variatre    bool    = false
	altura      float64 = 185
	py1         bool    = true
	py2         bool    = true
	py3         bool    = true
	py4         bool    = true
	tiro        int     = 0
	pytura      float64 = 0
)

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
		//n tratando erro pq n é para dar erro
	}

	defer readerpython.Close()

	screen.Fill(color.Black)
	p, _, err := image.Decode(readerpython)
	pyinvader, _ := ebiten.NewImageFromImage(p, ebiten.FilterNearest)
	//ebitenutil.DebugPrint(screen, "ta funfando :3")
	quadrado, _ := ebiten.NewImage(5, 15, ebiten.FilterNearest)

	quadrado.Fill(color.White)

	pyops := &ebiten.DrawImageOptions{}
	jogadorops := &ebiten.DrawImageOptions{}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		walkRight()
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		walkLeft()
	}
	jogadorops.GeoM.Translate(Gosavior, 185)
	screen.DrawImage(jogador, jogadorops)
	if altura <= 54 && gosaviorold <= 53 {
		py1 = false
	} else if altura <= pytura && gosaviorold < 106 && gosaviorold > 53 {
		py2 = false
	} else if altura <= pytura && gosaviorold > 106 && gosaviorold < 159 {
		py3 = false
	} else if gosaviorold >= 159 && altura <= pytura {
		py4 = false
	} else {

		pyops.GeoM.Translate(pydestroyer, 0)
		pyops2 := &ebiten.DrawImageOptions{}
		pyops3 := &ebiten.DrawImageOptions{}
		pyops4 := &ebiten.DrawImageOptions{}
		pyops2.GeoM.Translate(pydestroyer+53, pytura)
		pyops3.GeoM.Translate(pydestroyer+106, pytura)
		pyops4.GeoM.Translate(pydestroyer+159, pytura)
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

	quaopts.GeoM.Translate(gosaviorold, altura)

	if varia == true {
		screen.DrawImage(quadrado, quaopts)
		altura = altura - 5
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) && varia == false {
		gosaviorold = Gosavior + 35
		varia = true
		tiro++

	}

	if altura == 0 {
		altura = 185
		varia = false
	}

	if tiro >= 3 && tiro <= 5 {
		pytura = 10
	}
	if tiro >= 6 && tiro <= 8 {
		pytura = 15
	}
	if tiro >= 9 && tiro <= 13 {
		pytura = 30
	}
	if tiro >= 13 && tiro <= 15 {
		pytura = 45
	}
	if tiro >= 16 && tiro <= 20 {
		pytura = 55
	}
	if tiro > 20 {
		pytura = 185
		ebitenutil.DebugPrint(screen, "vc perdeu")
	} else {

		if py1 == false && py2 == false && py3 == false && py4 == false {
			ebitenutil.DebugPrint(screen, "vc ganhou")
		}
	}
	return nil
}

func walkRight() {
	Gosavior += 3
}

func walkLeft() {
	Gosavior -= 3
}
func main() {

	//m, _, err := image.Decode(reader)

	//if err != nil {

	//}
	ebiten.Run(jogo, 320, 240, 2, "Go Invader")
	ebiten.IsRunningSlowly()
}
