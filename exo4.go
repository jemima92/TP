package b2

func Ft_profit(prices []int) int {

	maxbenef := 0
	minPrix := prices[0]

	for _, price := range prices[1:] {
		if price < minPrix {
			minPrix = price
		} else {
			profit := price - minPrix
			if profit > maxbenef {
				maxbenef = profit
			}
		}
	}

	return maxbenef
}
