package code

import "math/big"

func filterRoutes(routes [][]*Path, amountIn *big.Int, fromLockedAmount map[uint64]*hexutil.Big) [][]*Path {
	if len(fromLockedAmount) == 0 {
		return routes
	}

	routesAfterNetworkCompliance := NetworkComplianceFilter(routes, fromLockedAmount)
	return CapacityValidationFilter(routesAfterNetworkCompliance, amountIn, fromLockedAmount)
}

func NetworkComplianceFilter(routes [][]*Path, fromLockedAmount map[uint64]*hexutil.Big) [][]*Path {
	var filteredRoutes [][]*Path

	for _, route := range routes {
		if isValidForNetworkCompliance(route, fromLockedAmount) {
			filteredRoutes = append(filteredRoutes, route)
		}
	}
	return filteredRoutes
}

func isValidForNetworkCompliance(route []*Path, fromLockedAmount map[uint64]*hexutil.Big) bool {
	fromIncluded, fromExcluded := setupRouteValidationMaps(fromLockedAmount)

	for _, path := range route {
		if fromExcluded[path.From.ChainID] {
			return false
		}
		if _, ok := fromIncluded[path.From.ChainID]; ok {
			fromIncluded[path.From.ChainID] = true
		}
	}

	for _, included := range fromIncluded {
		if !included {
			return false
		}
	}

	return true
}

func CapacityValidationFilter(routes [][]*Path, amountIn *big.Int, fromLockedAmount map[uint64]*hexutil.Big) [][]*Path {
	var filteredRoutes [][]*Path

	for _, route := range routes {
		if hasSufficientCapacity(route, amountIn, fromLockedAmount) {
			filteredRoutes = append(filteredRoutes, route)
		}
	}
	return filteredRoutes
}

func hasSufficientCapacity(route []*Path, amountIn *big.Int, fromLockedAmount map[uint64]*hexutil.Big) bool {
	totalRestAmount := calculateTotalRestAmount(route)

	for _, path := range route {
		if amount, ok := fromLockedAmount[path.From.ChainID]; ok {
			requiredAmountIn := new(big.Int).Sub(amountIn, amount.ToInt())
			if totalRestAmount.Cmp(requiredAmountIn) < 0 {
				return false
			}
			path.AmountIn = amount
			path.AmountInLocked = true
			totalRestAmount.Sub(totalRestAmount, amount.ToInt())
		}
	}
	return true
}
