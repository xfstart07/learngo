// Author: xufei
// Date: 2020/3/3

package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/chromedp/cdproto/page"

	"github.com/chromedp/chromedp"
)

func main() {
	err := printPDF("https://juejin.im/post/5e0c698d5188253aaf656925", "output.pdf")
	if err != nil {
		fmt.Println(err)
	}
}

func printPDF(url string, toPath string) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			buf, _, err = page.PrintToPDF().Do(ctx)
			return err
		}),
	})

	if err != nil {
		return fmt.Errorf("chromedp Run failed, %w", err)
	}

	if err := ioutil.WriteFile(toPath, buf, 0644); err != nil {
		return fmt.Errorf("write to file failed, %w", err)
	}

	return nil
}
