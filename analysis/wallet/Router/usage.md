# Initialisation
The `Router` struct is only initialised in one place in the application, in `wallet/api.go`.

The consumption complexity is very low:

```go
func NewAPI(s *Service) *API {
	router := NewRouter(s)
	return &API{s, s.reader, router}
}

type API struct {
    s      *Service
    reader *Reader
    router *Router
}
```

# Method and field calls

The `API.router` field is only called in two places:

## `GetSuggestedRoutes()`

This function is a straight through wrapper of `Router.suggestedRoutes()`:

```go
func (api *API) GetSuggestedRoutes(
	ctx context.Context,
	sendType SendType,
	addrFrom common.Address,
	addrTo common.Address,
	amountIn *hexutil.Big,
	tokenID string,
	disabledFromChainIDs,
	disabledToChaindIDs,
	preferedChainIDs []uint64,
	gasFeeMode GasFeeMode,
	fromLockedAmount map[uint64]*hexutil.Big,
) (*SuggestedRoutes, error) {
	log.Debug("call to GetSuggestedRoutes")
	return api.router.suggestedRoutes(ctx, sendType, addrFrom, addrTo, amountIn.ToInt(), tokenID, disabledFromChainIDs,
		disabledToChaindIDs, preferedChainIDs, gasFeeMode, fromLockedAmount)
}
```

## `CreateMultiTransaction()`

This function simply makes a read call to `Router.bridges`

```go
func (api *API) CreateMultiTransaction(ctx context.Context, multiTransactionCommand *transfer.MultiTransactionCommand, data []*bridge.TransactionBridge, password string) (*transfer.MultiTransactionCommandResult, error) {
	log.Debug("[WalletAPI:: CreateMultiTransaction] create multi transaction")
	return api.s.transactionManager.CreateMultiTransactionFromCommand(ctx, multiTransactionCommand, data, api.router.bridges, password)
}
```
