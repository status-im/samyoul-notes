# `newSuggestedRoutes()`

A constructor function used to create and initialize a struct that represents a collection of suggested routes for a cross-chain transaction. This function aggregates relevant data for each route to determine the best possible paths for a transaction.

## Purpose
The primary purpose of `newSuggestedRoutes` is to assemble a structured view of all potential transaction paths based on the criteria cost, speed, network preferences, and other operational parameters.

### Builds Routes:
Collects and organises routes that meet given criteria.

### Filters and Selects Routes:
Applies business logic or algorithms to select the most efficient or cost-effective routes.

### Provides Contextual Data:
Enriches routes with additional data such as cost estimates, gas fees, and estimated transaction times.

## Component Analysis
### `amountIn *big.Int`
The amount of tokens or currency to be transferred, influencing which paths are viable based on the maximum transferable amounts and balance restrictions.

### `candidates []*Path`
A list of potential paths that have been pre-determined as possible routes for the transaction.

### `fromLockedAmount map[uint64]*hexutil.Big`
A map indicating amounts locked or reserved on specific networks, possibly affecting the viability of certain paths.

### `filterRoutes` and [`findBest`](#findbest)
Helper functions that might be used within `newSuggestedRoutes` to filter out infeasible paths and then select the optimal path based on the criteria. These functions represent core logic for determining the best routes from a set of candidates.

TODO Additional analysis is required for both of these.

---

## `findBest`

Aims to find the most cost-effective route among a list of potential routes, where each route is a slice of `Path` structs.
Each `Path` has an associated cost stored in `Cost`, which is of type `big.Float`.
The function searches for the route with the minimum total cost.

### Function Operation
- 1: Initialization:
  - `best`: A slice of `*Path` that will store the best route found during the function's execution.
  - `bestCost`: A `big.Float` initialized to positive infinity `(math.Inf(1))`, which ensures any real cost comparison will initially be lower.
    - Possibly a hacky solution for the problem. It is probably totally fine.
- 2: Route Evaluation Loop:
  - The outer loop iterates over each route in the provided 2D slice `routes`.
  - For each `route`, it initializes currentCost to zero.
- 3: Path Cost Accumulation:
  - The inner loop iterates over each `Path` within the current `route`.
  - It accumulates the cost of all paths in the route into `currentCost`.
- 4: Comparison and Update:
  - After calculating the total cost for a route, it compares this cost with the current `bestCost`.
  - If `currentCost` is less than `bestCost` `(currentCost.Cmp(bestCost) == -1)`, it updates `bestCost` to this new lower cost and sets `best` to the current route.
- 5: Return Value:
  - After processing all `routes`, it returns the route with the lowest total cost as `best`. 

### Review
The function is straightforward, and finds the route with the lowest cost. However, I've highlighted a few potential improvements to consider:

- Precision and Performance:
  - While using `big.Float` for cost calculations ensures precision, which is essential in financial applications, operations on `big.Float` are computationally more expensive than on native float types.
  - We should confirm that this level of precision is necessary.
- Error Handling:
  - There is no error handling in this function.
  - If the route data structure is guaranteed to be well-formed (non-nil and correctly initialised), this is fine. Otherwise, we may want to add validation and/or recover from potential panics due to nil references.
- Optimisation:
  - If performance is critical and the number of routes is large, we should consider parallel processing techniques.
  - Divide the route slice and process each segment in a separate goroutine, then combine the results.
  - This requires careful handling to avoid race conditions when updating `best` and `bestCost`.
- Flexibility:
  - Currently, the function only minimises cost. If we need to optimise other metrics (e.g. speed, shortest path, etc)
  - We could consider passing in a comparator function as an argument.
