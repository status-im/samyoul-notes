# `newSuggestedRoutes()`

A constructor function used to create and initialize a struct that represents a collection of suggested routes for a cross-chain transaction. This function aggregates relevant data for each route to determine the best possible paths for a transaction.

## Purpose and Role
The primary purpose of `newSuggestedRoutes` is to assemble a structured view of all potential transaction paths based on the criteria cost, speed, network preferences, and other operational parameters.

### Aggregates Routes:
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

### `filterRoutes` and `findBest`
Helper functions that might be used within `newSuggestedRoutes` to filter out infeasible paths and then select the optimal path based on the criteria. These functions represent core logic for determining the best routes from a set of candidates.

TODO Additional analysis is required for both of these.
