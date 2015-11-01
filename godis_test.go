package godis

func setUp() *Godis {
	return NewGodis()
}

type Case struct {
	key   string
	value string
}

var cases = []Case{
	{"key1", "value1"},
	{"key2", "value2"},
	{"key 3", "value 3"},    // keys with spaces
	{"மொழி", "தமிழ்"},       // unicode
	{"key1", "new value 1"}, // overwrite a key
	/*{"tested", true},        // boolean value
	{"test_num", 7},         // int value
	{"PI", 3.14},            // float value
	*/
}

var integers = []Case{
	{"int", "234"},
	{"long", "223344231"},
	{"negative", "-554"},
	{"zero", "0"},
}

var floats = []Case{
	{"flt64", "22423234.1223"},
	{"flt32", "443.21"},
	{"fltExp", "123.34e23"},
	{"negative", "-123.34e23"},
	{"zero", "0"},
}

// var strings = []Case{
// 	{"மொழி", "தமிழ்"}, // value with string
// }
