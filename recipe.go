package archnemesis

type Recipe []string

var RecipeList = map[string]Recipe{
	//T1
	"Mana Siphoner":     {"Concecrator", "Dynamo"},
	"Heralding Minions": {"Arcane Buffer", "Dynamo"},
	"Rejugenating":      {"Gargantuan", "Vampiric"},
	"Assassin":          {"Vampiric", "Deadeye"},
	"Drought Bringer":   {"Deadeye", "Malediction"},
	"Hexer":             {"Echoist", "Chaosweaver"},
	"Necromancer":       {"Overcharged", "Bombardier"},
	"Executioner":       {"Frenzied", "Berserker"},
	"Mirror Image":      {"Echoist", "Soul Conduit"},
	"Flame Strider":     {"Flameweaver", "Hasted"},
	"Frost Strider":     {"Frostweaver", "Hasted"},
	"Storm Rider":       {"Stormweaver", "Hasted"},
	"Evocationist":      {"Flameweaver", "Frostweaver", "Stormweaver"},
	"Corrupter":         {"Chaosweaver", "Bloodletter"},
	"Entangler":         {"Bloodletter", "Toxic"},
	"Ice Prison":        {"Permafrost", "Sentinel"},
	"Magma Barrier":     {"Bonebreaker", "Incendiary"},

	//T2
	"Invulnerable":       {"Sentinel", "Concecrator", "Juggernaut"},
	"Temporal Bubble":    {"Juggernaut", "Arcane Buffer", "Hexer"},
	"Trickster":          {"Assassin", "Echoist", "Overcharged"},
	"Empowering Minions": {"Gargantuan", "Necromancer", "Executioner"},
	"Crystal Skinned":    {"Rejugenating", "Berserker", "Permafrost"},
	"Empowered Elements": {"Chaosweaver", "Evocationist", "Steel-Invused"},
	"Effigy":             {"Hexer", "Corrupter", "Malediction"},
	"Treant Horde":       {"Steel-Invused", "Toxic", "Sentinel"},
	"Corpse Detonator":   {"Necromancer", "Incendiary"},
	"Soul Eater":         {"Necromancer", "Soul Conuit", "Gargantuan"},

	// T3
	"Lunaris-Touched":    {"Invulnerable", "Empowering Minions", "Frost Strider"},
	"Solaris-Touched":    {"Invulnerable", "Empowering Minions", "Magma Barrier"},
	"Arakaali-Touched":   {"Assassin", "Entangler", "Corpse Detonator"},
	"Brine King-Touched": {"Heralding Minions", "Storm Rider", "Ice Prison"},
	"Tukohama-Touched":   {"Executioner", "Bonebreaker", "Magma Barrier"},
	"Abberath-Touched":   {"Rejugenating", "Frenzied", "Flame Strider"},
	"Shakari-Touched":    {"Drought Bringer", "Entangler", "Soul Eater"},

	// T4
	"Innocence-Touched": {"Lunaris-Touched", "Solaris-Touched", "Mana Siphoner", "Mirror Image"},
	"Kitava-Touched":    {"Tukohama-Touched", "Abberath-Touched", "Corrupter", "Corpse Detonator"},
}
