package main

func removeDuplicates(s string) string {
	if len(s) == 1 {
		return s
	}

	str := []rune(s)

	stack := NewStack()

	for i := 0; i < len(s); i++ {
		now := str[i]
		if stack.Peek() != now {
			stack.Push(now)
		} else {
			for stack.Peek() == now {
				stack.Pop()
			}
		}
	}

	ans := ""
	for !stack.Empty() {
		ans = string(stack.Pop()) + ans
	}

	return ans
}
