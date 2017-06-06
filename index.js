exports.Format = function(s, m) {
	var ESCCHAR = '#';
	var i = 0;
	var output = "";

	while (i < s.length) {
		if (s[i] == ESCCHAR) {
			var j = i + 1;
			while (j <= s.length && s[j] == ESCCHAR) {
				j = j + 1;
				if ((j - i) % 2 == 0 ) {
					output = output + ESCCHAR;
				}
			}
			if ((j - i) %2 != 0) {
				var param = "";
				while (j < s.length && s[j] != ' ') {
					param = param + s[j];
					j++;
				}
				if (param.length > 0) {
					output = output + m[param];
				}
				if (j == s.length) {
					return output;
				}
			}
			i = j;
		}
		output = output + s[i];
		i++;
	}
	return output;
}
