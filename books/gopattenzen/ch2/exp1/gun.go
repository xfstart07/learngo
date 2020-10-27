package main

import "fmt"

// 里氏替换原则

func main() {
	sanMao := Soldier{}
	gun := Rifle{}
	sanMao.setGun(&gun)
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

type Soldier struct {
	gun AbstractGun
}

func (s *Soldier) setGun(gun AbstractGun) {
	s.gun = gun
}

func (s Soldier) killEnemy() {
	s.gun.shoot()
}
