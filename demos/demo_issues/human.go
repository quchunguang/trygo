package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
)

const (
	SIZE = 300
	N    = 3000
)

func maxValue(A [SIZE][SIZE][3]float32) (m [3]float32) {
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			for c := 0; c < 3; c++ {
				if A[i][j][c] > m[c] {
					m[c] = A[i][j][c]
				}
			}
		}
	}
	return
}

func human() {
	rand.Seed(27)
	var A [SIZE][SIZE][3]float32
	A[SIZE/2][SIZE/2][0] = 100.0
	A[SIZE/2][SIZE/2][1] = 100.0
	A[SIZE/2][SIZE/2][2] = 100.0
	for s := 0; s < N; s++ {
		for i := 1; i < SIZE-1; i++ {
			for j := 1; j < SIZE-1; j++ {
				for c := 0; c < 3; c++ {
					switch rand.Intn(4) {
					case 0:
						A[i-1][j][c] += A[i][j][c] / 2
					case 1:
						A[i+1][j][c] += A[i][j][c] / 2
					case 2:
						A[i][j-1][c] += A[i][j][c] / 2
					case 3:
						A[i][j+1][c] += A[i][j][c] / 2
					}
					A[i][j][c] /= 2
				}
			}
		}
	}

	m := maxValue(A)
	img := image.NewNRGBA(image.Rect(0, 0, SIZE, SIZE))
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			img.Set(i, j, color.RGBA{
				R: uint8(A[i][j][0] / m[0] * 255),
				G: uint8(A[i][j][1] / m[1] * 255),
				B: uint8(A[i][j][2] / m[2] * 255),
				A: 255,
			})
		}
	}

	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
