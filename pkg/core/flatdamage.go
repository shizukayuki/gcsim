package core

type FlatDamage struct {
	ActorIndex int
	Abil       string
	Damage     float64
}

func (ai *AttackInfo) AddFlatDmg(m FlatDamage) {
	if m.Abil == "" {
		return
	}
	ai.FlatDmg = append(ai.FlatDmg, m)
}
