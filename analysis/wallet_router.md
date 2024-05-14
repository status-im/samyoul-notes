# Constants and Types
- `EstimateUsername` and `EstimatePubKey` are used as placeholder data in estimation calculations.
- `SendType` is an enum that categorises the types of transaction that the router supports
  - e.g. transfers, ENS (Ethereum Name Service) operations, purchasing stickers, etc.

# Utility Functions
- `IsCollectiblesTransfer()`: Checks if the transaction involves ERC721 or ERC1155 collectibles.
- `FetchPrices()`: Fetches current token prices from a market manager service.
- `FindToken()`: Finds a token within a given blockchain network, handling both ERC20 tokens and collectibles.
- `isAvailableBetween()`: Checks if a certain transaction type is permissible between two networks.
- `EstimateGas()`: Estimates the gas required for transactions such as ENS registrations or transfers.

# Structs
- `Node` and `Graph`
  - Are used for constructing a graph of potential transaction paths, representing routes for transferring tokens across networks.
- `Path`
  - Represents a transaction path including details like from/to networks, fees, and estimated costs.
- `Router`
  - Manages routing of transactions using different bridges defined for various transaction types.
  - [Usage analysis](./wallet_router_usage.md)
  - **Functions:**
    - `requireApproval()`
      - Determines if token approval is required before executing a transaction over a bridge.
    - `getBalance()` and `getERC1155Balance()`
      - Fetch balance of ERC20 or ERC1155 tokens respectively.
    - `suggestedRoutes()`
      - Suggests possible transaction routes based on constraints like disabled or preferred chain IDs, and calculates fees and estimated transaction times.
  - **Exclusive Utility Functions:**
    - `containsNetworkChainID()`
      - Determines if a specified blockchain network's Chain ID is present within a given list of Chain IDs
    - `newSuggestedRoutes()`
      - newSuggestedRoutes function is a crucial part of the route selection and optimization process routing transactions across potentially complex network paths.
      - [Detailed Analysis of `newSuggestedRoutes()`](./wallet_router_newSuggestedRoutes.md)

# Transaction Handling
- The code handles transactions through various bridges which abstract the different transaction types.
  - e.g. simple transfers, ERC721, ERC1155, and cross-chain transfers.
  - TODO need to explore down this route
- A method for calculating optimal transaction routes (`suggestedRoutes`) evaluates multiple paths, considering various factors like fees, token prices, and network constraints.
  - TODO need to explore down this route

# Error Handling and Data Integrity
- There is significant error handling to ensure that only valid transactions are processed.
  - For example, when finding tokens or fetching prices, errors lead to early termination with descriptive messages.
- Use of concurrency (`sync.Mutex`, `async.NewAtomicGroup`) suggests an approach of handling multiple transactions or operations concurrently.
  - TODO need to explore more on this.

# Observations and Potential Improvements
- Error Reporting:
  - Detailed errors and/or logging could help with debugging.
- Optimisation:
  - The graph-building and route optimisation process could potentially be optimised for performance, especially when dealing with a large number of possible paths.
  - Knapsack, mD Knapsack, Travelling Salesman
  - TODO need to explore more on this.