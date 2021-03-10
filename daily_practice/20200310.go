package daily_practice

/*
基本计算器
只有加减运算
*/
func calculate(s string) (ans int) {
	ops := []int{1}
	sign := 1
	l := len(s)
	for i := 0; i < l; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = ops[len(ops)-1]
			i++
		case '-':
			sign = -ops[len(ops)-1]
			i++
		case '(':
			ops = append(ops, sign)
			i++
		case ')':
			ops = ops[:len(ops)-1]
			i++
		default:
			num := 0
			for ; i < l && '0' <= s[i] && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			ans += sign * num
		}

	}
	return ans
}
