package main

import (
	arch "github.com/pro7ech/archnemesis"
)

func main() {
	arch.Missing("Innocence-Touched", Inventory)
	arch.Missing("Kitava-Touched", Inventory)
}

var Inventory = map[string]int{

	//T0
	"Arcane Buffer": 3,
	"Berserker":     1,
	"Bloodletter":   3,
	"Bombardier":    2,
	"Bonebreaker":   1,
	"Chaosweaver":   3,
	"Concecrator":   1,
	"Deadeye":       2,
	"Dynamo":        2,
	"Echoist":       2,
	"Flameweaver":   2,
	"Frenzied":      0,
	"Frostweaver":   2,
	"Gargantuan":    1,
	"Hasted":        2,
	"Incendiary":    0,
	"Juggernaut":    0,
	"Malediction":   2,
	"Overcharged":   2,
	"Permafrost":    2,
	"Sentinel":      1,
	"Soul Conduit":  2,
	"Steel-Invused": 1,
	"Stormweaver":   1,
	"Toxic":         2,
	"Vampiric":      2,

	//T1
	"Assassin":          0,
	"Corrupter":         1,
	"Drought Bringer":   1,
	"Entangler":         0,
	"Evocationist":      0,
	"Executioner":       0,
	"Flame Strider":     1,
	"Frost Strider":     0,
	"Heralding Minions": 1,
	"Hexer":             1,
	"Ice Prison":        1,
	"Magma Barrier":     1,
	"Mana Siphoner":     0,
	"Mirror Image":      0,
	"Necromancer":       0,
	"Rejugenating":      1,
	"Storm Rider":       0,

	//T2
	"Corpse Detonator":   1,
	"Crystal Skinned":    1,
	"Effigy":             1,
	"Empowered Elements": 1,
	"Empowering Minions": 0,
	"Invulnerable":       0,
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
	"Innocence-Touched": 1,
	"Kitava-Touched":    0,
}

