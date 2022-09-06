package player

type Stats struct {
	Health          int
	Score           int
	CurrentLocation string
}

type Player struct {
	PlayerInfo Stats
}

func Setup() Stats {
	var stats Stats
	stats.Health = 10
	stats.Score = 0
	stats.CurrentLocation = "entrance"
	return stats
}
