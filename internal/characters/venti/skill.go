package venti

import (
	"github.com/genshinsim/gcsim/internal/frames"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
)

var skillPressFrames []int
var skillHoldFrames []int

func (c *char) Skill(p map[string]int) action.ActionInfo {
	ai := combat.AttackInfo{
		ActorIndex:   c.Index,
		Abil:         "Skyward Sonnett",
		AttackTag:    combat.AttackTagElementalArt,
		ICDTag:       combat.ICDTagNone,
		ICDGroup:     combat.ICDGroupDefault,
		Element:      attributes.Anemo,
		Durability:   50,
		Mult:         skillPress[c.TalentLvlSkill()],
		HitWeakPoint: true,
	}

	act := action.ActionInfo{
		Frames:          frames.NewAbilFunc(skillPressFrames),
		AnimationLength: skillPressFrames[action.InvalidAction],
		CanQueueAfter:   skillPressFrames[action.ActionDash], // earliest cancel
		Post:            skillPressFrames[action.ActionDash], // earliest cancel
		State:           action.SkillState,
	}

	cd := 360
	cdstart := 21
	hitmark := 51
	if p["hold"] != 0 {
		cd = 900
		cdstart = 34
		hitmark = 74
		ai.Mult = skillHold[c.TalentLvlSkill()]

		act = action.ActionInfo{
			Frames:          frames.NewAbilFunc(skillHoldFrames),
			AnimationLength: skillHoldFrames[action.InvalidAction],
			CanQueueAfter:   skillHoldFrames[action.ActionHighPlunge], // earliest cancel
			Post:            skillHoldFrames[action.ActionHighPlunge], // earliest cancel
			State:           action.SkillState,
		}
	}

	c.Core.QueueAttack(ai, combat.NewDefCircHit(4, false, combat.TargettableEnemy), 0, hitmark, c.c2)
	c.Core.QueueParticle("venti", 3, attributes.Anemo, hitmark+100)

	c.SetCDWithDelay(action.ActionSkill, cd, cdstart)

	return act
}