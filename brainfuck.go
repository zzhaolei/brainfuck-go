package brainfuckgo

import "fmt"

func Brainfuck(code []byte) error {
	var (
		tape       [30000]uint8
		index, ptr int
	)
	for ; index < len(code); index += 1 {
		switch code[index] {
		case '>':
			ptr++
		case '<':
			ptr--
		case '+':
			tape[ptr]++
		case '-':
			tape[ptr]--
		case '.':
			fmt.Printf("%c", tape[ptr])
		case ',':
			fmt.Print("[ascii]>>> ")
			var r byte
			if _, err := fmt.Scanf("%c", &r); err != nil {
				return err
			}
			tape[ptr] = r
		case '[':
			if tape[ptr] == 0 {
				index++
				loop := 1
				for ; loop > 0; index++ {
					switch code[index] {
					case '[':
						loop++
					case ']':
						loop--
					}
				}
			}
		case ']':
			if tape[ptr] != 0 {
				index--
				loop := 1
				for ; loop > 0; index-- {
					switch code[index] {
					case '[':
						loop--
					case ']':
						loop++
					}
				}
			}
		}
	}
	return nil
}
