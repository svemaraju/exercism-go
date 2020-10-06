package space

type Planet string

func Age(age float64, p Planet) float64 {
	if p == "Earth" {
		return age / (31557600)
	} else if p == "Mercury" {
		return age / (0.2408467 * 31557600)
	} else if p == "Venus" {
		return age / (0.61519726 * 31557600)
	} else if p == "Mars" {
		return age / (1.8808158 * 31557600)
	} else if p == "Jupiter" {
		return age / (11.862615 * 31557600)
	} else if p == "Saturn" {
		return age / (29.447498 * 31557600)
	} else if p == "Uranus" {
		return age / (84.016846 * 31557600)
	} else if p == "Neptune" {
		return age / (164.79132 * 31557600)
	} 
	return 0
}
