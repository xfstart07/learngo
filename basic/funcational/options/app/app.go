// Author: xufei
// Date: 2019-05-15

package app

type App struct {
	opts Options
}

type Options struct {
	UserAgent string
}

func New(options ...func(*App)) *App {
	c := &App{}

	for _, f := range options {
		f(c)
	}

	return c
}

func UserAgent(ua string) func(*App) {
	return func(c *App) {
		c.opts.UserAgent = ua
	}
}
