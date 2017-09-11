package stringf

// Format format s using map m
// Keys of map m can be anything but must not contains spaces
func Format(s string, m map[string]string) string {
	ESCCHAR := byte('{');
	ESCCHAREND := byte('}')
	i := 0
	output := ""
	for i < len(s) {
		if s[i] == ESCCHAR {
			j := i + 1
			for j < len(s) && s[j] == ESCCHAR {
				j++
				if (j - i) % 2 == 0 {
					output += string(ESCCHAR)
				}
			}
			if (j - i) % 2 != 0 {
				param := ""
				for j < len(s) && s[j] != ESCCHAREND {
					param += string(s[j])
					j++
				}
				if len(param) > 0 {
					output += m[param]
				}

				if j == len(s) {
					return output
				} else if s[j] == ESCCHAREND {
					j++
				}
			}
			i = j
		}
		if i < len(s) {
			output += string(s[i])
		}
		i++
	}
	return output
}
