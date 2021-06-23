package Program1

func num2roman(num int)string{
	str:=""
	var sliNum = [] int {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var sliStr = [] string {"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for i := 0; i < len(sliNum); i++ {
		for num >= sliNum[i] {
			num-=sliNum[i]
			str+=sliStr[i]
		}
	}
	return str
}