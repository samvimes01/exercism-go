package allergies

var items = []struct {
	item  string
	value uint
}{
	{"eggs", 1}, {"peanuts", 2}, {"shellfish", 4}, {"strawberries", 8}, {"tomatoes", 16}, {"chocolate", 32}, {"pollen", 64}, {"cats", 128},
}

func Allergies(allergies uint) []string {
	if allergies > 1024 {
		for allergies > 1024 {
			allergies -= 1024
		}
	}
	if allergies > 512 {
		allergies -= 512
	}
	if allergies > 256 {
		allergies -= 256
	}
	r := []string{}
	for i := len(items) - 1; i >= 0; i-- {
		if items[i].value <= allergies {
			r = append([]string{items[i].item}, r...)
			allergies -= items[i].value
			if items[i].value <= allergies {
				for items[i].value < allergies {
					allergies -= items[i].value
				}
			}
		}
		if allergies <= 0 {
			break
		}
	}
	return r
}

func AllergicTo(allergies uint, allergen string) bool {
	//panic("Please implement the AllergicTo function")
	als := Allergies(allergies)
	for _, v := range als {
		if allergen == v {
			return true
		}
	}
	return false
}
