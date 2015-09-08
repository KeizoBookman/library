package view

import (
	"encoding/json"
	"github.com/KeizoBookman/library/seat"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("template/index.tmpl")
	if err != nil {
		ViewFail("index")
		return
	}
	n := strings.Index(r.URL.Path, "index.cgi")
	if n == -1 {
		ViewFail("path")
	}
	data := struct{ Path string }{r.URL.Path[:n]}
	err = tmpl.Execute(w, data)
	if err != nil {
		ViewFail(err.Error())
	}

}

func List(w http.ResponseWriter, r *http.Request) {
	var data interface{}
	tmpl, err := template.ParseFiles("template/list.tmpl")
	if err != nil {
		ViewFail("list")
		return
	}
	err = tmpl.Execute(w, data)
}

func Search(w http.ResponseWriter, r *http.Request) {
	var data interface{}
	tmpl, err := template.ParseFiles("template/search.tmpl")
	if err != nil {
		ViewFail("search")
		return
	}
	err = tmpl.Execute(w, data)
}

func NewSeat(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/seat.tmpl")
	if err != nil {
		ViewFail("new:")
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		ViewFail("new:")
		return
	}
	r.ParseForm()

	c := seat.Character{}
	c.Name.Character = r.FormValue("character")
	c.Name.Player = r.FormValue("player")
	c.Parsonality.Detail = r.FormValue("detail")
	c.Parsonality.Crest.Pos = r.FormValue("pos")
	c.Parsonality.Crest.Figure = r.FormValue("figure")
	c.Parsonality.Goal = r.FormValue("goals")
	c.Parsonality.Feeling = r.FormValue("feeling")

	age, err := strconv.Atoi(r.FormValue("age"))
	c.Creation.Age = age
	c.Creation.Sex = r.FormValue("sex")
	height, _ := strconv.Atoi(r.FormValue("height"))
	c.Creation.Height = height
	weight, _ := strconv.Atoi(r.FormValue("weight"))
	c.Creation.Weight = weight
	c.Creation.Colors.Eye = r.FormValue("eye")
	c.Creation.Colors.Hair = r.FormValue("hair")
	c.Creation.Colors.Skin = r.FormValue("skin")
	c.Creation.Colors.Other = r.FormValue("other")

	lv, _ := strconv.Atoi(r.FormValue("lv"))
	c.Level = lv
	exp, _ := strconv.Atoi(r.FormValue("exp"))
	c.Exp = exp

	c.Tribe.Name = r.FormValue("tribe")
	c.Tribe.Privilege = r.FormValue("privilege")
	//c.Tribe.Type = r.FormValue("type")

	c.Mask = r.FormValue("mask")
	c.Feature = r.FormValue("feature")
	//degree
	agile, _ := strconv.Atoi(r.FormValue("agile"))
	c.Ability.Agile = agile
	intelligence, _ := strconv.Atoi(r.FormValue("Intelligence"))
	c.Ability.Intelligence = intelligence
	mind, _ := strconv.Atoi(r.FormValue("mind"))
	c.Ability.Mind = mind
	luck, _ := strconv.Atoi(r.FormValue("luck"))
	c.Ability.Lucky = luck
	pow, _ := strconv.Atoi(r.FormValue("power"))
	c.Ability.Power = pow
	/*
		c.Amendment.Action
		c.Amendment.Avoid
		c.Amendment.Damage
		c.Amendment.Guard
		c.Amendment.Hit
		c.Amendment.Hp
		c.Amendment.Invok
		c.Amendment.Penetration
		c.Amendment.Resist
	*/

	s := seat.Seat{}
	s.Character = c
	file, err := os.OpenFile("public/store", os.O_RDWR, 0666)
	if err != nil {
		ViewFail("critical--- " + err.Error())
		return
	}

	bytes, err := json.Marshal(c)
	if err != nil {
		ViewFail("Can not Marshal--- " + err.Error())
		return
	}

	_, err = file.Write(bytes)
	if err != nil {
		ViewFail("Can not Write---" + err.Error())
		return
	}
}
