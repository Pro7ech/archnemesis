package main

import (
	arch "github.com/pro7ech/archnemesis"
)

func main() {
	arch.Missing("Innocence-Touched", Inventory)
}

var Inventory = map[string]int{

	//T0
	"Arcane Buffer": 2,
	"Berserker":     2,
	"Bloodletter":   2,
	"Bombardier":    1,
	"Bonebreaker":   2,
	"Chaosweaver":   1,
	"Concecrator":   2,
	"Deadeye":       2,
	"Dynamo":        1,
	"Echoist":       2,
	"Flameweaver":   2,
	"Frenzied":      0,
	"Frostweaver":   2,
	"Gargantuan":    1,
	"Hasted":        2,
	"Incendiary":    0,
	"Juggernaut":    2,
	"Malediction":   2,
	"Overcharged":   1,
	"Permafrost":    2,
	"Sentinel":      2,
	"Soul Conduit":  2,
	"Steel-Invused": 3,
	"Stormweaver":   1,
	"Toxic":         2,
	"Vampiric":      1,

	//T1
	"Assassin":          0,
	"Corrupter":         1,
	"Drought Bringer":   1,
	"Entangler":         0,
	"Evocationist":      0,
	"Executioner":       0,
	"Flame Strider":     1,
	"Frost Strider":     1,
	"Heralding Minions": 1,
	"Hexer":             1,
	"Ice Prison":        1,
	"Magma Barrier":     0,
	"Mana Siphoner":     1,
	"Mirror Image":      1,
	"Necromancer":       1,
	"Rejugenating":      1,
	"Storm Rider":       0,

	//T2
	"Corpse Detonator":   1,
	"Crystal Skinned":    1,
	"Effigy":             1,
	"Empowered Elements": 1,
	"Empowering Minions": 0,
	"Invulnerable":       1,
	"Soul Eater":         0,
	"Temporal Bubble":    1,
	"Treant Horde":       1,
	"Trickster":          1,

	// T3
	"Abberath-Touched":   1,
	"Arakaali-Touched":   1,
	"Brine King-Touched": 1,
	"Lunaris-Touched":    0,
	"Solaris-Touched":    0,
	"Tukohama-Touched":   0,
	"Shakari-Touched":    0,

	// T4
	"Innocence-Touched": 0,
	"Kitava-Touched":    0,
}

