package _1

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Door         int
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	return strings.TrimSpace(s)[:3]
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {
	b := sage{
		"Buddha",
		"The belief of no beliefs",
	}

	g := sage{
		"Gandhi",
		"Be the change",
	}

	m := sage{
		"Martin Luther King",
		"Hatred never ceases with hatred but with love alone is hatred",
	}

	c := car{
		"Ford",
		"F150",
		2,
	}

	sages := []sage{b, g, m}
	cars := []car{c}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatal(err)
	}
}
