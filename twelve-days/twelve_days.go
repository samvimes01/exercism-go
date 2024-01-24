package twelve

var days = []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}

var gifts = []string{"a Partridge in a Pear Tree", "two Turtle Doves", "three French Hens", "four Calling Birds", "five Gold Rings", "six Geese-a-Laying", "seven Swans-a-Swimming", "eight Maids-a-Milking", "nine Ladies Dancing", "ten Lords-a-Leaping", "eleven Pipers Piping", "twelve Drummers Drumming"}

func Verse(i int) string {
	verse := "On the " + days[i-1] + " day of Christmas my true love gave to me: "
	for j := i; j > 0; j-- {
		verse += gifts[j-1]
		if j == 1 {
			verse += "."
		} else if j == 2 {
			verse += ", and "
		} else {
			verse += ", "
		}
	}
	return verse
}

func Song() string {
	song := ""
	for i := 1; i < 12; i++ {
		song += Verse(i) + "\n"
	}
	return song + Verse(12)
}
