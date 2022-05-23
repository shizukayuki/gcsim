package beidou

import (
	"github.com/genshinsim/gcsim/internal/frames"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/player/shield"
)

var skillFrames []int

const skillHitmark = 23

func (c *char) Skill(p map[string]int) action.ActionInfo {
	//0 for base dmg, 1 for 1x bonus, 2 for max bonus
	counter := p["counter"]
	if counter >= 2 {
		counter = 2
		c.a4()
	}
	ai := combat.AttackInfo{
		ActorIndex: c.Index,
		Abil:       "Tidecaller (E)",
		AttackTag:  combat.AttackTagElementalArt,
		ICDTag:     combat.ICDTagNone,
		ICDGroup:   combat.ICDGroupDefault,
		StrikeType: combat.StrikeTypeBlunt,
		Element:    attributes.Electro,
		Durability: 50,
		Mult:       skillbase[c.TalentLvlSkill()] + skillbonus[c.TalentLvlSkill()]*float64(counter),
	}
	c.Core.QueueAttack(ai, combat.NewDefCircHit(1, false, combat.TargettableEnemy), skillHitmark, skillHitmark)

	//2 if no hit, 3 if 1 hit, 4 if perfect
	c.Core.QueueParticle("beidou", 2+float64(counter), attributes.Electro, 100)

	if counter > 0 {
		//add shield
		c.Core.Player.Shields.Add(&shield.Tmpl{
			Src:        c.Core.F,
			ShieldType: shield.ShieldBeidouThunderShield,
			Name:       "Beidou Skill",
			HP:         shieldPer[c.TalentLvlSkill()]*c.MaxHP() + shieldBase[c.TalentLvlSkill()],
			Ele:        attributes.Electro,
			Expires:    c.Core.F + 900, //15 sec
		})
	}

	c.SetCDWithDelay(action.ActionSkill, 450, 4)

	return action.ActionInfo{
		Frames:          frames.NewAbilFunc(skillFrames),
		AnimationLength: skillFrames[action.InvalidAction],
		CanQueueAfter:   skillFrames[action.ActionDash], // earliest cancel
		Post:            skillFrames[action.ActionDash], // earliest cancel
		State:           action.SkillState,
	}
}