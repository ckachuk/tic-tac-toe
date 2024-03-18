package main

import (
	"html/template"
	"io"
	"strconv"
	tictactoe "tic-tac-toe/logic"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type InitGame struct{
	Player1Name string
	Player2Name string
	Board [9]string
}

type HandleChoice struct{
	Board [9] string
}


type Template struct{
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context )error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main(){
	e := echo.New()
	e.Use(middleware.Logger())
	templ:= &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = templ

	game := tictactoe.InitGame("Carlos", "Aldi")

	e.GET("/", func(c echo.Context) error {		
		data:= InitGame{
			Player1Name: game.Player1.Name,
			Player2Name: game.Player2.Name,
			Board: game.Board,
		}

		return c.Render(200, "index", data)
	})
	
	e.POST("/choice/:id", func(c echo.Context) error {
		ichoice,_ := strconv.Atoi(c.Param("id"))
		game.UpdateBoard(int8(ichoice))

		data:= &HandleChoice{ Board: game.Board}
		return c.Render(200, "displayBoard", data)
	})
	
	e.GET("/reset", func(c echo.Context) error {
			game.ResetGame()
			
			return c.Render(200, "displayBoard", &HandleChoice{ Board: game.Board})
	})

	e.Logger.Fatal(e.Start(":8080"))
}