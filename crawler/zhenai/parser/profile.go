// Author: xufei
// Date: 2018/11/23

package parser

import (
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	NameSelector     = "h1.nickName"
	IDSelector       = ".m-userInfo div.id"
	DesSelector      = ".m-content-box .purple-btns .purple"
	DesOtherSelector = ".m-content-box .pink-btns .pink"
)

const (
	IntRe = `([0-9]*)`
)

type ProfileParser struct {
	engine.ParseArgs
}

func (s *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return parseProfile(contents, url, s.Name, s.Gender)
}

func (s *ProfileParser) SerializeParse() (name string, args engine.ParseArgs) {
	return "ProfileParser", s.ParseArgs
}

func NewProfileParser(param engine.ParseArgs) *ProfileParser {
	return &ProfileParser{
		ParseArgs: param,
	}
}

func parseProfile(body []byte, url, name, gender string) engine.ParseResult {
	doc, err := transformDocument(body)
	if err != nil {
		log.Printf("transform document error: %v", err)
		return engine.ParseResult{}
	}

	result := engine.ParseResult{}

	id := doc.Find(IDSelector).Text()
	if id == "" {
		log.Printf("用户ID获取错误，不保存")
		return result
	}

	profile := model.Profile{
		Name:   name,
		Gender: gender,
		URL:    url,
	}
	profile.ID = doc.Find(IDSelector).Text()
	profile.ID = strings.Replace(profile.ID, "ID：", "", -1)

	DesLen := doc.Find(DesSelector).Size()
	doc.Find(DesSelector).Each(func(index int, node *goquery.Selection) {
		if strings.Contains("未婚,离异,丧偶", node.Text()) {
			profile.Marital = node.Text()
		}

		if strings.Contains(node.Text(), "岁") {
			profile.Age = ReInt(node.Text())
		}
		if strings.Contains(node.Text(), "座(") {
			profile.Astrology = node.Text()
		}
		if strings.Contains(node.Text(), "cm") {
			profile.Height = ReInt(node.Text())
		}
		if strings.Contains(node.Text(), "kg") {
			profile.Weight = ReInt(node.Text())
		}
		if strings.Contains(node.Text(), "工作地") {
			profile.WorkCity = strings.Replace(node.Text(), "工作地:", "", -1)
		}
		if strings.Contains(node.Text(), "月收入") {
			profile.Income = node.Text()
		}

		if index == DesLen-1 {
			profile.Education = node.Text()
		}
	})

	doc.Find(DesOtherSelector).Each(func(index int, node *goquery.Selection) {
		if strings.Contains(node.Text(), "族") {
			profile.Nation = node.Text()
		}
		if strings.Contains(node.Text(), "籍贯") {
			profile.Native = node.Text()
		}
		if strings.Contains(node.Text(), "房") {
			profile.IsBuyHouse = node.Text()
		}
		if strings.Contains(node.Text(), "车") {
			profile.IsBuyCar = node.Text()
		}
	})

	result.Items = append(result.Items, engine.Item{
		ID:      profile.ID,
		Payload: profile,
	})

	return result
}

var reInt = regexp.MustCompile(IntRe)

func ReInt(ageString string) int {
	findString := reInt.FindString(ageString)
	val, _ := strconv.Atoi(findString)
	return val
}
