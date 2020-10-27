// Author: xufei
// Date: 2019/4/30

package main

import "fmt"

// 定义建造者接口
type Builder interface {
	SetName(string) Builder
	SetArms(string) Builder
	Build() *Character
}

// 具体的产品类
type Character struct {
	Name string
	Arms string
}

func (p *Character) SetName(name string) {
	p.Name = name
}

func (p *Character) SetArms(arms string) {
	p.Arms = arms
}

// 具体的建造者，实现了建造接口
type CharacterBuilder struct {
	character *Character
}

func (p *CharacterBuilder) SetName(name string) Builder {
	if p.character == nil {
		p.character = &Character{}
	}
	p.character.SetName(name)

	return p
}

func (p *CharacterBuilder) SetArms(arms string) Builder {
	if p.character == nil {
		p.character = &Character{}
	}
	p.character.SetArms(arms)

	return p
}

func (p *CharacterBuilder) Build() *Character {
	return p.character
}

// Director 导演类

type Director struct {
	Builder
}

// 导演类 根据相同的构建规则 创建出不同的 对象。
func (p Director) Create(name, arms string) *Character {
	return p.Builder.SetName(name).SetArms(arms).Build()
}

func main() {
	var builder Builder = &CharacterBuilder{}
	director := &Director{Builder: builder}

	character := director.Create("leon", "brn")
	fmt.Println(character.Name, character.Arms)
}
