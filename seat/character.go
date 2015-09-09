package seat

import (
	"net/http"
	"strconv"
)

type Seat struct {
	Character Character
}

type Character struct {
	Name        Name
	Parsonality Parsonality
	Creation    Creation
	Level       int
	Exp         int
	Tribe       Tribe
	Mask        string
	Feature     string
	Degrees     []Degree
	Ability     Ability
	Amendment   Amendment
	Guard       Guard
	Equipment   Equipment
	Tarrents    []Tarrent
}

type Name struct {
	Character string
	Player    string
}

type Parsonality struct {
	Detail  string
	Crest   Crest
	Goal    string
	Feeling string
}

type Crest struct {
	Pos    string
	Figure string
}

type Creation struct {
	Orgnization string
	Age         int
	Sex         string
	Height      int
	Weight      int
	Colors      Colors
}

type Colors struct {
	Eye   string
	Hair  string
	Skin  string
	Other string
}

type Tribe struct {
	Name      string
	Type      int
	Privilege string
}

type Degree struct {
	Name string
	Main bool
	Type string
}

type Ability struct {
	Power        int
	Agile        int
	Intelligence int
	Mind         int
	Lucky        int
}

type Amendment struct {
	Param
}

type Guard struct {
	Strong int
	Hart   int
}

type Damage struct {
	Pysics int
	Magic  int
}

type Equipment struct {
	Hands     []Arm
	Armor     Arm
	Ornaments []Arm
}

type Arm struct {
	Name string
	Pos  string
	Param
	Enhance []string
}

type Tarrent struct {
	Name     string
	Timing   int
	Distance string
	Taget    string
	Cost     string
	Effect   string
}

type Param struct {
	Hit         int
	Avoid       int
	Invok       int
	Resist      int
	Penetration int
	Damage      Damage
	Action      int
	Hp          int
	Guard       Guard
}

func (*Seat) Version() (string, error) {
	return "0.0.1", nil
}

func (c *Character) GetFormValue(r *http.Request) error {
	c.getFormAbility(r)
	c.getFormAmendment(r)
	c.getFormCreation(r)
	c.getFormDegrees(r)
	c.getFormEquipment(r)
	c.getFormExp(r)
	c.getFormFeature(r)
	c.getFormGuard(r)
	c.getFormLevel(r)
	c.getFormMask(r)
	c.getFormName(r)
	c.getFormParsonality(r)
	c.getFormTarrents(r)
	c.getFormTribe(r)

	return nil
}

func (c *Character) getFormAbility(r *http.Request) error {
	agile, err := strconv.Atoi(r.FormValue("agile"))
	if err != nil {
	}
	c.Ability.Agile = agile

	intelligence, err := strconv.Atoi(r.FormValue("Intelligence"))
	if err != nil {
	}
	c.Ability.Intelligence = intelligence

	mind, err := strconv.Atoi(r.FormValue("mind"))
	if err != nil {
	}
	c.Ability.Mind = mind

	luck, err := strconv.Atoi(r.FormValue("luck"))
	if err != nil {
	}
	c.Ability.Lucky = luck

	pow, err := strconv.Atoi(r.FormValue("power"))
	if err != nil {
	}
	c.Ability.Power = pow

	action, err := strconv.Atoi(r.FormValue("action"))
	if err != nil {
	}
	c.Amendment.Action = action
	return nil
}

func (c *Character) getFormAmendment(r *http.Request) error {
	c.Amendment.Avoid = 0
	c.Amendment.Damage.Pysics = 0
	c.Amendment.Damage.Magic = 0
	c.Amendment.Hit = 0
	c.Amendment.Hp = 0
	c.Amendment.Invok = 0
	c.Amendment.Penetration = 0
	c.Amendment.Resist = 0
	c.Amendment.Guard.Strong = 0
	c.Amendment.Guard.Hart = 0
	return nil
}

func (c *Character) getFormCreation(r *http.Request) error {
	age, err := strconv.Atoi(r.FormValue("age"))
	if age > -1 {
	}
	c.Creation.Age = age

	c.Creation.Sex = r.FormValue("sex")
	height, err := strconv.Atoi(r.FormValue("height"))
	if err != nil {
	}
	c.Creation.Height = height

	weight, err := strconv.Atoi(r.FormValue("weight"))
	if err != nil {
	}
	c.Creation.Weight = weight

	c.Creation.Colors.Eye = r.FormValue("eye")
	c.Creation.Colors.Hair = r.FormValue("hair")
	c.Creation.Colors.Skin = r.FormValue("skin")
	c.Creation.Colors.Other = r.FormValue("other")
	return nil

}

func (c *Character) getFormDegrees(r *http.Request) error {

	return nil
}

func (c *Character) getFormEquipment(r *http.Request) error {

	return nil
}

func (c *Character) getFormExp(r *http.Request) error {
	exp, err := strconv.Atoi(r.FormValue("exp"))
	if err != nil {
	}
	c.Exp = exp
	return nil
}

func (c *Character) getFormFeature(r *http.Request) error {
	c.Feature = r.FormValue("feature")
	return nil
}

func (c *Character) getFormGuard(r *http.Request) error {

	return nil
}

func (c *Character) getFormLevel(r *http.Request) error {
	lv, err := strconv.Atoi(r.FormValue("lv"))
	if err != nil {
	}
	c.Level = lv
	return nil
}

func (c *Character) getFormMask(r *http.Request) error {
	c.Mask = r.FormValue("mask")
	return nil
}

func (c *Character) getFormName(r *http.Request) error {
	c.Name.Character = r.FormValue("character")
	c.Name.Player = r.FormValue("player")

	return nil
}

func (c *Character) getFormParsonality(r *http.Request) error {
	c.Parsonality.Detail = r.FormValue("detail")
	c.Parsonality.Crest.Pos = r.FormValue("pos")
	c.Parsonality.Crest.Figure = r.FormValue("figure")
	c.Parsonality.Goal = r.FormValue("goals")
	c.Parsonality.Feeling = r.FormValue("feeling")
	return nil
}

func (c *Character) getFormTarrents(r *http.Request) error {
	return nil
}

func (c *Character) getFormTribe(r *http.Request) error {
	c.Tribe.Name = r.FormValue("tribe")
	c.Tribe.Privilege = r.FormValue("privilege")
	//c.Tribe.Type = r.FormValue("type")
	return nil
}
