var assert = require('assert');
var stringf = require('../index.js');

describe('stringf', function() {
	it('should not format escape characters', function() {
		var str = stringf.Format('hi {{name{name}}', { name: 'Kieu Thanh'});
		assert.equal(str, 'hi {nameKieu Thanh}');
	});

	it('should format parameters', function() {
		var str = stringf.Format("hi {name} ", { name: "thanh" });
		assert.equal(str, 'hi thanh ');
	});
});
