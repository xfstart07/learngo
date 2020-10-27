// Author: xufei
// Date: 2019/4/30

package car

import "fmt"

type Benz struct {
	Color  Color
	Wheels Wheels
	Speed  Speed
}

func (b Benz) Drive() error {
	fmt.Println(fmt.Sprintf("驾驶 %s 奔驰, 行驶速度 %v...", b.Color, b.Speed))
	return nil
}

func (Benz) Stop() error {
	fmt.Println("奔驰停车...")
	return nil
}
