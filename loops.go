package loop_temp_struct

type smallStruct struct {
	Num  uint32
	Data []byte
}

// Val is an exported variable ensured to keep functions A and B from being optimized out of existence.
var Val int

func A(vs []smallStruct) int {
	var out int
	for i := 0; i < len(vs); i++ {
		temp := vs[i]
		switch temp.Num {
		case 4:
			out += 41
		case 15:
			out -= 6
		case 589:
			out = 12
		default:
			out = -1
		}
	}
	return out
}

func B(vs []smallStruct) int {
	var out int
	for i := 0; i < len(vs); i++ {
		temp := &vs[i]
		switch temp.Num {
		case 4:
			out += 41
		case 15:
			out -= 6
		case 589:
			out = 12
		default:
			out = -1
		}
	}
	return out
}
