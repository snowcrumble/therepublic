// Package therepublic 理想国 Salute Plato
package therepublic

// Person 人有心灵和身体两种属性
type Person struct {
	Heart int
	Body  int
}

// Doctor 医生是人
type Doctor Person

// NewBestDoctor 最好的医生身体不行
func NewBestDoctor() Doctor {
	return Doctor{
		Heart: 1,
		Body:  0,
	}
}

// CureAbility 医生的治疗能力
func (d Doctor) CureAbility() int {
	return d.Heart + d.Body
}
