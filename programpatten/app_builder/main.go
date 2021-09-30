// Author: xufei
// Date: 2020/2/29

package main

import "fmt"

type App struct {
	Name string
}

type AppBuilder []func(*App) error

func (ab *AppBuilder) Register(f ...func(*App) error) {
	*ab = append(*ab, f...)
}

func (ab *AppBuilder) Build(a *App) error {
	for _, f := range *ab {
		if err := f(a); err != nil {
			return err
		}
	}

	return nil
}

// NOTE: App 构建案例
func main() {
	a := &App{}
	ab := &AppBuilder{}

	ab.Register(func(app *App) error {
		app.Name = "Awesome, wow!"
		return nil
	})

	if err := ab.Build(a); err != nil {
		fmt.Println(err)
	}

	println(a.Name)
}
