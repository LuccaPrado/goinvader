package main

//imports necessários
import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image"
	"image/color"
	_ "image/png"
	"os"
	"strconv"
)

//declarando variaveis
var (
	Gosavior    float64 = 0
	gosaviorold float64 = 0
	pydestroyer float64 = 0
	pyimage1    float64 = 0
	pyimage2    float64 = 54 //add 1 px de espaço a cada imagem
	pyimage3    float64 = 107
	pyimage4    float64 = 160
	pyimage5    float64 = 213
	varia       bool    = false
	altura      float64 = 185
	py1         bool    = true
	py2         bool    = true
	py3         bool    = true
	py4         bool    = true
	py5         bool    = true
	tiro        int     = 0
	contatiro   string  = "0" //só p seguir padrão
	pytura      float64 = 0
	nivel       int     = 1
)

//Loop principal
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
	if altura <= 54 && gosaviorold <= pyimage2 {
		py1 = false
	} else if altura <= pytura && gosaviorold < pyimage3 && gosaviorold > pyimage2 {
		py2 = false
	} else if altura <= pytura && gosaviorold > pyimage3 && gosaviorold < pyimage4 {
		py3 = false
	} else if gosaviorold >= pyimage4 && altura <= pytura && gosaviorold < pyimage5 {
		py4 = false
	} else if gosaviorold >= pyimage5 && altura <= pytura {
		py5 = false
	} else {
		//configura e desenha os pys
		pyops.GeoM.Translate(pydestroyer, 0)
		pyops2 := &ebiten.DrawImageOptions{}
		pyops3 := &ebiten.DrawImageOptions{}
		pyops4 := &ebiten.DrawImageOptions{}
		pyops5 := &ebiten.DrawImageOptions{}
		pyops2.GeoM.Translate(pyimage2, pytura)
		pyops3.GeoM.Translate(pyimage3, pytura)
		pyops4.GeoM.Translate(pyimage4, pytura)
		pyops5.GeoM.Translate(pyimage5, pytura)
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
		if py5 == true {
			screen.DrawImage(pyinvader, pyops5)
		}
	}
	quaopts := &ebiten.DrawImageOptions{}

	quaopts.GeoM.Translate(gosaviorold, altura)
	//define que atirou
	if varia == true {
		screen.DrawImage(quadrado, quaopts)
		altura = altura - 5
	}
	//atirar
	if ebiten.IsKeyPressed(ebiten.KeyUp) && varia == false {
		atirar()
	}
	//aonde o tiro está
	if altura == 0 {
		altura = 185
		varia = false
	}
	//começo de pontuação?
	contatiro = strconv.Itoa(tiro)
	//checa se vc ja atirou mais que podia, se sim, vc perdeu
	//checa nivel
	if nivel == 1 {
		if tiro > 8 {
			pytura = 185
			ebitenutil.DebugPrint(screen, "vc perdeu e usou so "+contatiro+" tiros ")
			ebitenutil.DebugPrint(screen, "\n\n\n Use a tecla enter para recomecar ou s para sair!")
			if ebiten.IsKeyPressed(ebiten.KeyS) == true {
				os.Exit(0)
			}
			if ebiten.IsKeyPressed(ebiten.KeyEnter) == true {
				reiniciar()
			}
		} else {
			//arrumando por pura vontade
			if tiro >= 3 && tiro <= 5 {
				pytura += 1
			}
			if tiro >= 6 && tiro <= 8 {
				pytura += 1
			}

			//checa se todos os pythons tão vivos
			if py1 == false && py2 == false && py3 == false && py4 == false && py5 == false {

				ebitenutil.DebugPrint(screen, "vc ganhou e usou so "+contatiro+" tiros ")
				ebitenutil.DebugPrint(screen, "\n\n\n Use a tecla enter para nivel 2 ou s para sair!")
				if ebiten.IsKeyPressed(ebiten.KeyS) == true {
					os.Exit(0)
				}
				if ebiten.IsKeyPressed(ebiten.KeyEnter) == true {
					nivel = 2
					reiniciar()
				}
			}
		}
	} else if nivel == 2 {
		if tiro > 8 {
			pytura = 185
			ebitenutil.DebugPrint(screen, "vc perdeu e usou so "+contatiro+" tiros ")
			ebitenutil.DebugPrint(screen, "\n\n\n Use a tecla enter para recomecar ou s para sair!")
			if ebiten.IsKeyPressed(ebiten.KeyS) == true {
				os.Exit(0)
			}
			if ebiten.IsKeyPressed(ebiten.KeyEnter) == true {
				reiniciar()
			}
		} else {
			//arrumando por pura vontade
			if tiro >= 3 && tiro <= 5 {
				pytura += 1
			}
			if tiro >= 6 && tiro <= 8 {
				pytura += 2
			}
			//checa se todos os pythons tão vivos
			if py1 == false && py2 == false && py3 == false && py4 == false && py5 == false {

				ebitenutil.DebugPrint(screen, "vc ganhou e usou so "+contatiro+" tiros ")
				ebitenutil.DebugPrint(screen, "\n\n\n Use a tecla enter para recomecar ou s para sair!")
				if ebiten.IsKeyPressed(ebiten.KeyS) == true {
					os.Exit(0)
				}
				if ebiten.IsKeyPressed(ebiten.KeyEnter) == true {
					reiniciar()
				}
			}
		}
	}
	return nil
}
func atirar() {
	gosaviorold = Gosavior + 35
	varia = true
	tiro++
}
func reiniciar() {
	//função de reiniciar
	Gosavior = 0
	gosaviorold = 0
	pydestroyer = 0
	pyimage1 = 0
	pyimage2 = 54
	pyimage3 = 107
	pyimage4 = 160
	pyimage5 = 213
	varia = false
	altura = 185
	py1 = true
	py2 = true
	py3 = true
	py4 = true
	py5 = true
	tiro = 0
	contatiro = "0"
	pytura = 0
	if nivel == 2 {
		pytura = 10
	} else {
		nivel = 1
	}

}

//andar para direita
func walkRight() {
	if Gosavior < 260 {
		Gosavior += 3
	}

}

//andar para esquerda
func walkLeft() {
	if Gosavior > 0 {
		Gosavior -= 3
	}

}

//Inicia o jogo
func main() {

	ebiten.Run(jogo, 320, 240, 2, "Go Invader")
	ebiten.IsRunningSlowly() //checa se o jogo esta rodando lento  https://godoc.org/github.com/hajimehoshi/ebiten#IsRunningSlowly
}
