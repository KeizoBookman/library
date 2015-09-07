package seat

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
	Mask        int
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
	Type      string
	Privilege string
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
	Hand1     Arm
	Hand2     Arm
	Armor     Arm
	Ornament  Arm
	Ornament2 Arm
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
