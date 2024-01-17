// Solution for twofer task.
package twofer

// ShareWith based on name return different string.
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	} else {
		return "One for " + name + ", one for me."
	}
}
