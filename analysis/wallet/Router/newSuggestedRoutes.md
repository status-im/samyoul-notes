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

### [`filterRoutes`](#filterRoutes) and [`findBest`](#findbest)
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

### Analysis
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

---

## `filterRoutes`

This function aims to filter routes based on "locked amounts" specified for each network (chain ID) and the required `amountIn`. It uses `fromLockedAmount` to constrain two main levels of filtering to determine which routes are viable.

### Function stages:
- 1: Initial Check:
  - If the `fromLockedAmount` map is empty, it immediately returns all the routes as valid because there are no constraints to apply.
- 2: First Level of Filtering (`filteredRoutesLevel1`):
  - It initializes `fromIncluded` and `fromExcluded` maps based on the `fromLockedAmount` values:
    - If the locked amount for a chain ID is zero, that chain ID is marked as excluded.
    - If the locked amount is non-zero, that chain ID needs to be included, and the route must pass through it.
  - For each route, it checks:
    - If any of the route's paths are on an excluded chain ID (`fromExcluded`), the route is marked **not** OK (`routeOk = false`).
    - If a path's chain ID is in `fromIncluded`, it marks that chain ID as having been included in the route.
    - Finally, it ensures that all required (`fromIncluded`) chain IDs are indeed part of the route. If any are missing, the route is marked **not** OK.
  - Only routes that are marked OK are passed to the next level of filtering.
- 3: Second Level of Filtering (`filteredRoutesLevel2`):
  - For each route that passed the first level:
    - 1: checks every path in the route against the `fromLockedAmount` to determine if the remaining routes can satisfy the locked amount conditions.
    - 2: calculates `requiredAmountIn` by subtracting the locked amount for the path's chain ID from the total `amountIn`.
    - 3: then calculates `restAmountIn` by summing the `MaxAmountIn` of all other paths in the route, excluding the current one.
    - 4: If `restAmountIn` (the total potential input from other paths) is sufficient to cover `requiredAmountIn`, it "locks" the amount for that path.
      - It could be worth renaming the word "lock" to something like "allocates" or "assigns"
    - 5: Else, the route is marked **not** OK (`routeOk = false`).
  - Routes that pass this check are added to `filteredRoutesLevel2`.

## Analysis:
### Efficiency Concerns:
- The function performs nested loops over potentially large sets of routes and paths, which can be computationally intensive, especially with the additional checks and conditions within each loop.

### Potential Improvements:
- **Optimise the Inner Loops:**
  - The calculation of restAmountIn could be optimized. Instead of recalculating it for each path in the route, it could be calculated once per route and then adjusted for each path.
- **Memoisation:**
  - If the routes and locked amounts do not change frequently, we could use memoisation to cache results for certain input combinations and possibly improve performance.
- **Concurrency:**
  - Processing routes in parallel could speed up the filtering significantly.

### Maintainability:
The function is complex and handles a lot of logic, which might make it hard to maintain or debug.

It could benefit from being broken down into smaller, more manageable functions that handle specific parts of the logic. Example the setup of `fromIncluded` and `fromExcluded`, or the calculations of `requiredAmountIn` and `restAmountIn`.

TODO - Add a suggested function PR.

### Error Handling:
There is no error handling or validation for the inputs (checking for nil pointers, etc.).