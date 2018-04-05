package parser

import (
	"learngo/crawler/engine"
	"regexp"
	"strconv"
	"learngo/crawler/model"
)

var ageRe = `<td><span class="label">年龄：</span>(\d+)岁</td>`

func ParseProfile(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(ageRe)
	profile := model.Profile{}
	match := re.FindSubmatch(contents)

	if match != nil {
		age, err := strconv.Atoi(string(match[1]))
		if err != nil {
			profile.Age = age
		}
	}
}
