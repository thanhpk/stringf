exports.Format = function(s, m) {
	var ESCCHAR = '{';
	var ESCCHAREND = '}';
	var i = 0;
	var output = "";

	while (i < s.length) {
		if (s[i] == ESCCHAR) {
			var j = i + 1;
			while (j <= s.length && s[j] == ESCCHAR) {
				j++;
				if ((j - i) % 2 == 0 ) {
					output += ESCCHAR;
				}
			}
			if ((j - i) %2 != 0) {
				var param = "";
				while (j < s.length && s[j] != ESCCHAREND) {
					param += s[j];
					j++;
				}
				if (param.length > 0) {
					output += m[param];
				}

				if (j == s.length) {
					return output;
				} else if (s[j] == ESCCHAREND) {
					j++;
				}
			}
			i = j;
		}
		if (i < s.length) {
			output += s[i];
		}
		i++;
	}
	return output;
};
