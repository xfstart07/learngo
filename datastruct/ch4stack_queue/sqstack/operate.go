// Author: xufei
// Date: 2019-07-26

package sqstack

func InitStack() *SqStack {
	stack := new(SqStack)
	stack.top = -1
	return stack
}

func Push(stack *SqStack, e int) bool {
	if stack.top == MAXSIZE {
		return false
	}

	stack.top++
	stack.data[stack.top] = e
	return true
}

func Pop(stack *SqStack, e *int) bool {
	if stack.top == -1 {
		return false
	}
	*e = stack.data[stack.top]
	stack.top--
	return true
}
