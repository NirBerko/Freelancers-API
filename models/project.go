package models

type Project struct {
	DBBase
	Title       string
	Description string
	Skills      map[int]Skill
}
