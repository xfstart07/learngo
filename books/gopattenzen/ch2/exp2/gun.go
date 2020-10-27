package main

import "fmt"

// 里氏替换原则

func main() {
	sanMao := Snipper{}
	gun := Aug{}
	sanMao.setGun(gun)
	sanMao.killEnemy()
}

// 枪的接口
type AbstractGun interface {
	shoot()
}

type Handgun struct {
}

func (h Handgun) shoot() {
	fmt.Println("手枪射击...")
}

type Rifle struct {
}

func (r Rifle) shoot() {
	fmt.Println("步枪射击...")
}

type MachineGun struct {
}

func (m MachineGun) shoot() {
	fmt.Println("机枪扫射...")
}

type Aug struct {
}

func (a Aug) zoomOut() {
	fmt.Println("通过望远镜察看敌人...")
}

func (a Aug) shoot() {
	fmt.Println("Aug射击...")
}

type Snipper struct {
	Aug
}

func (s *Snipper) setGun(aug Aug) {
	s.Aug = aug
}

func (s Snipper) killEnemy() {
	s.zoomOut()
	s.shoot()
}
