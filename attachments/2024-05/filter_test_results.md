<details>
  <summary><strong>0 ms</strong> ✅ TestSetupRouteValidationMapsV2</summary>
  <ul>
    <li>
      <details>
        <summary><em>0 ms</em> <strong>✅ passed</strong> Mixed_locked_amounts</summary>
        <pre>=== RUN   TestSetupRouteValidationMapsV2/Mixed_locked_amounts<br/>--- PASS: TestSetupRouteValidationMapsV2/Mixed_locked_amounts (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><em>0 ms</em> <strong>✅ passed</strong> All_amounts_locked</summary>
        <pre>=== RUN   TestSetupRouteValidationMapsV2/All_amounts_locked<br/>--- PASS: TestSetupRouteValidationMapsV2/All_amounts_locked (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><em>0 ms</em> <strong>✅ passed</strong> No_amounts_locked</summary>
        <pre>=== RUN   TestSetupRouteValidationMapsV2/No_amounts_locked<br/>--- PASS: TestSetupRouteValidationMapsV2/No_amounts_locked (0.00s)<br/></pre>
      </details>
    </li>
  </ul>
</details>
<details>
  <summary><strong>0 ms</strong> ✅ TestCalculateTotalRestAmountV2</summary>
  <ul>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>✅ passed</em> Multiple_paths_with_varying_amounts</summary>
        <pre>=== RUN   TestCalculateTotalRestAmountV2/Multiple_paths_with_varying_amounts<br/>--- PASS: TestCalculateTotalRestAmountV2/Multiple_paths_with_varying_amounts (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>✅ passed</em> Single_path</summary>
        <pre>=== RUN   TestCalculateTotalRestAmountV2/Single_path<br/>--- PASS: TestCalculateTotalRestAmountV2/Single_path (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>✅ passed</em> No_paths</summary>
        <pre>=== RUN   TestCalculateTotalRestAmountV2/No_paths<br/>--- PASS: TestCalculateTotalRestAmountV2/No_paths (0.00s)<br/></pre>
      </details>
    </li>
  </ul>
</details>
<details>
  <summary><strong>0 ms</strong> ✅ TestIsValidForNetworkComplianceV2</summary>
  <ul>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>✅ passed</em> Valid_route_with_required_chain_IDs_included</summary>
        <pre>=== RUN   TestIsValidForNetworkComplianceV2/Valid_route_with_required_chain_IDs_included<br/>--- PASS: TestIsValidForNetworkComplianceV2/Valid_route_with_required_chain_IDs_included (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>✅ passed</em> Invalid_route_with_excluded_chain_ID</summary>
        <pre>=== RUN   TestIsValidForNetworkComplianceV2/Invalid_route_with_excluded_chain_ID<br/>--- PASS: TestIsValidForNetworkComplianceV2/Invalid_route_with_excluded_chain_ID (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>✅ passed</em> Route_missing_required_chain_ID</summary>
        <pre>=== RUN   TestIsValidForNetworkComplianceV2/Route_missing_required_chain_ID<br/>--- PASS: TestIsValidForNetworkComplianceV2/Route_missing_required_chain_ID (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>✅ passed</em> Valid_route_with_multiple_included_chain_IDs</summary>
        <pre>=== RUN   TestIsValidForNetworkComplianceV2/Valid_route_with_multiple_included_chain_IDs<br/>--- PASS: TestIsValidForNetworkComplianceV2/Valid_route_with_multiple_included_chain_IDs (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>✅ passed</em> Invalid_route_missing_one_of_the_required_chain_IDs</summary>
        <pre>=== RUN   TestIsValidForNetworkComplianceV2/Invalid_route_missing_one_of_the_required_chain_IDs<br/>--- PASS: TestIsValidForNetworkComplianceV2/Invalid_route_missing_one_of_the_required_chain_IDs (0.00s)<br/></pre>
      </details>
    </li>
  </ul>
</details>
<details>
  <summary><strong>0 ms</strong> ❌ TestHasSufficientCapacityV2</summary>
  <ul>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>✅ passed</em> Sufficient_capacity_with_multiple_paths</summary>
        <pre>=== RUN   TestHasSufficientCapacityV2/Sufficient_capacity_with_multiple_paths<br/>--- PASS: TestHasSufficientCapacityV2/Sufficient_capacity_with_multiple_paths (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>✅ passed</em> Insufficient_capacity</summary>
        <pre>=== RUN   TestHasSufficientCapacityV2/Insufficient_capacity<br/>--- PASS: TestHasSufficientCapacityV2/Insufficient_capacity (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>❌ failed</em> Exact_capacity_match</summary>
        <pre>=== RUN   TestHasSufficientCapacityV2/Exact_capacity_match<br/>filter_test.go:248: <br/>Error Trace: /Users/samuel/go/src/github.com/Samyoul/status-go/services/wallet/router/filter_test.go:248<br/>Error: Not equal: <br/>expected: true<br/>actual  : false<br/>Test: TestHasSufficientCapacityV2/Exact_capacity_match<br/>--- FAIL: TestHasSufficientCapacityV2/Exact_capacity_match (0.00s)<br/></pre>
      </details>
    </li>
  </ul>
</details>
<details>
  <summary><strong> ❌ TestFilterNetworkComplianceV2</strong></summary>
  <ul>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>❌ failed</em> Mixed_routes_with_valid_and_invalid_paths</summary>
        <pre>=== RUN   TestFilterNetworkComplianceV2/Mixed_routes_with_valid_and_invalid_paths<br/>filter_test.go:512: <br/>Error Trace: /Users/samuel/go/src/github.com/Samyoul/status-go/services/wallet/router/filter_test.go:512<br/>Error: Not equal: <br/>expected: [][]*router.PathV2{[]*router.PathV2{(*router.PathV2)(0xc0004f7790), (*router.PathV2)(0xc0004f7860)}, []*router.PathV2{(*router.PathV2)(0xc0004f7930), (*router.PathV2)(0xc0004f7a00), (*router.PathV2)(0xc0004f7ad0)}}<br/>actual: [][]*router.PathV2{[]*router.PathV2{(*router.PathV2)(0xc0004f71e0), (*router.PathV2)(0xc0004f72b0)}}<br/><br/>Diff:<br/>--- Expected<br/>+++ Actual<br/>@@ -1,2 +1,2 @@<br/>-([][]*router.PathV2) (len=2) {<br/>+([][]*router.PathV2) (len=1) {<br/>([]*router.PathV2) (len=2) {<br/>@@ -92,139 +92,2 @@<br/>})<br/>-},<br/>-([]*router.PathV2) (len=3) {<br/>-(*router.PathV2)({<br/>-BridgeName: (string) "",<br/>-From: (*params.Network)({<br/>-ChainID: (uint64) 1,<br/>-ChainName: (string) "",<br/>-RPCURL: (string) "",<br/>-OriginalRPCURL: (string) "",<br/>-FallbackURL: (string) "",<br/>-OriginalFallbackURL: (string) "",<br/>-BlockExplorerURL: (string) "",<br/>-IconURL: (string) "",<br/>-NativeCurrencyName: (string) "",<br/>-NativeCurrencySymbol: (string) "",<br/>-NativeCurrencyDecimals: (uint64) 0,<br/>-IsTest: (bool) false,<br/>-Layer: (uint64) 0,<br/>-Enabled: (bool) false,<br/>-ChainColor: (string) "",<br/>-ShortName: (string) "",<br/>-TokenOverrides: ([]params.TokenOverride) <nil>,<br/>-RelatedChainID: (uint64) 0<br/>-}),<br/>-To: (*params.Network)(<nil>),<br/>-FromToken: (*token.Token)(<nil>),<br/>-AmountIn: (*hexutil.Big)(<nil>),<br/>-AmountInLocked: (bool) false,<br/>-AmountOut: (*hexutil.Big)(<nil>),<br/>-SuggestedPriorityFees: (*router.PriorityFees)(<nil>),<br/>-TxBaseFee: (*hexutil.Big)(<nil>),<br/>-TxPriorityFee: (*hexutil.Big)(<nil>),<br/>-TxGasAmount: (uint64) 0,<br/>-TxBonderFees: (*hexutil.Big)(<nil>),<br/>-TxTokenFees: (*hexutil.Big)(<nil>),<br/>-TxL1Fee: (*hexutil.Big)(<nil>),<br/>-ApprovalRequired: (bool) false,<br/>-ApprovalAmountRequired: (*hexutil.Big)(<nil>),<br/>-ApprovalContractAddress: (*common.Address)(<nil>),<br/>-ApprovalBaseFee: (*hexutil.Big)(<nil>),<br/>-ApprovalPriorityFee: (*hexutil.Big)(<nil>),<br/>-ApprovalGasAmount: (uint64) 0,<br/>-ApprovalL1Fee: (*hexutil.Big)(<nil>),<br/>-EstimatedTime: (router.TransactionEstimation) 0,<br/>-requiredTokenBalance: (*big.Int)(<nil>),<br/>-requiredNativeBalance: (*big.Int)(<nil>)<br/>-}),<br/>-(*router.PathV2)({<br/>-BridgeName: (string) "",<br/>-From: (*params.Network)({<br/>-ChainID: (uint64) 2,<br/>-ChainName: (string) "",<br/>-RPCURL: (string) "",<br/>-OriginalRPCURL: (string) "",<br/>-FallbackURL: (string) "",<br/>-OriginalFallbackURL: (string) "",<br/>-BlockExplorerURL: (string) "",<br/>-IconURL: (string) "",<br/>-NativeCurrencyName: (string) "",<br/>-NativeCurrencySymbol: (string) "",<br/>-NativeCurrencyDecimals: (uint64) 0,<br/>-IsTest: (bool) false,<br/>-Layer: (uint64) 0,<br/>-Enabled: (bool) false,<br/>-ChainColor: (string) "",<br/>-ShortName: (string) "",<br/>-TokenOverrides: ([]params.TokenOverride) <nil>,<br/>-RelatedChainID: (uint64) 0<br/>-}),<br/>-To: (*params.Network)(<nil>),<br/>-FromToken: (*token.Token)(<nil>),<br/>-AmountIn: (*hexutil.Big)(<nil>),<br/>-AmountInLocked: (bool) false,<br/>-AmountOut: (*hexutil.Big)(<nil>),<br/>-SuggestedPriorityFees: (*router.PriorityFees)(<nil>),<br/>-TxBaseFee: (*hexutil.Big)(<nil>),<br/>-TxPriorityFee: (*hexutil.Big)(<nil>),<br/>-TxGasAmount: (uint64) 0,<br/>-TxBonderFees: (*hexutil.Big)(<nil>),<br/>-TxTokenFees: (*hexutil.Big)(<nil>),<br/>-TxL1Fee: (*hexutil.Big)(<nil>),<br/>-ApprovalRequired: (bool) false,<br/>-ApprovalAmountRequired: (*hexutil.Big)(<nil>),<br/>-ApprovalContractAddress: (*common.Address)(<nil>),<br/>-ApprovalBaseFee: (*hexutil.Big)(<nil>),<br/>-ApprovalPriorityFee: (*hexutil.Big)(<nil>),<br/>-ApprovalGasAmount: (uint64) 0,<br/>-ApprovalL1Fee: (*hexutil.Big)(<nil>),<br/>-EstimatedTime: (router.TransactionEstimation) 0,<br/>-requiredTokenBalance: (*big.Int)(<nil>),<br/>-requiredNativeBalance: (*big.Int)(<nil>)<br/>-}),<br/>-(*router.PathV2)({<br/>-BridgeName: (string) "",<br/>-From: (*params.Network)({<br/>-ChainID: (uint64) 3,<br/>-ChainName: (string) "",<br/>-RPCURL: (string) "",<br/>-OriginalRPCURL: (string) "",<br/>-FallbackURL: (string) "",<br/>-OriginalFallbackURL: (string) "",<br/>-BlockExplorerURL: (string) "",<br/>-IconURL: (string) "",<br/>-NativeCurrencyName: (string) "",<br/>-NativeCurrencySymbol: (string) "",<br/>-NativeCurrencyDecimals: (uint64) 0,<br/>-IsTest: (bool) false,<br/>-Layer: (uint64) 0,<br/>-Enabled: (bool) false,<br/>-ChainColor: (string) "",<br/>-ShortName: (string) "",<br/>-TokenOverrides: ([]params.TokenOverride) <nil>,<br/>-RelatedChainID: (uint64) 0<br/>-}),<br/>-To: (*params.Network)(<nil>),<br/>-FromToken: (*token.Token)(<nil>),<br/>-AmountIn: (*hexutil.Big)(<nil>),<br/>-AmountInLocked: (bool) false,<br/>-AmountOut: (*hexutil.Big)(<nil>),<br/>-SuggestedPriorityFees: (*router.PriorityFees)(<nil>),<br/>-TxBaseFee: (*hexutil.Big)(<nil>),<br/>-TxPriorityFee: (*hexutil.Big)(<nil>),<br/>-TxGasAmount: (uint64) 0,<br/>-TxBonderFees: (*hexutil.Big)(<nil>),<br/>-TxTokenFees: (*hexutil.Big)(<nil>),<br/>-TxL1Fee: (*hexutil.Big)(<nil>),<br/>-ApprovalRequired: (bool) false,<br/>-ApprovalAmountRequired: (*hexutil.Big)(<nil>),<br/>-ApprovalContractAddress: (*common.Address)(<nil>),<br/>-ApprovalBaseFee: (*hexutil.Big)(<nil>),<br/>-ApprovalPriorityFee: (*hexutil.Big)(<nil>),<br/>-ApprovalGasAmount: (uint64) 0,<br/>-ApprovalL1Fee: (*hexutil.Big)(<nil>),<br/>-EstimatedTime: (router.TransactionEstimation) 0,<br/>-requiredTokenBalance: (*big.Int)(<nil>),<br/>-requiredNativeBalance: (*big.Int)(<nil>)<br/>-})<br/>-}<br/>-}<br/>+([][]*router.PathV2) <nil><br/><br/>Test: TestFilterNetworkComplianceV2/Mixed_routes_with_valid_and_invalid_paths<br/>--- FAIL: TestFilterNetworkComplianceV2/Mixed_routes_with_valid_and_invalid_paths (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>✅ passed</em> All_valid_routes</summary>
        <pre>=== RUN   TestFilterNetworkComplianceV2/All_valid_routes<br/>--- PASS: TestFilterNetworkComplianceV2/All_valid_routes (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>❌ failed</em> All_invalid_routes</summary>
        <pre>=== RUN   TestFilterNetworkComplianceV2/All_invalid_routes<br/>filter_test.go:512: <br/>Error Trace: /Users/samuel/go/src/github.com/Samyoul/status-go/services/wallet/router/filter_test.go:512<br/>Error: Not equal: <br/>expected: [][]*router.PathV2{}<br/>actual: [][]*router.PathV2(nil)<br/><br/>Diff:<br/>--- Expected<br/>+++ Actual<br/>@@ -1,3 +1,2 @@<br/>-([][]*router.PathV2) {<br/>-}<br/>+([][]*router.PathV2) <nil><br/><br/>Test: TestFilterNetworkComplianceV2/All_invalid_routes<br/>--- FAIL: TestFilterNetworkComplianceV2/All_invalid_routes (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>❌ failed</em> Empty_routes</summary>
        <pre>=== RUN   TestFilterNetworkComplianceV2/Empty_routes<br/>filter_test.go:512: <br/>Error Trace: /Users/samuel/go/src/github.com/Samyoul/status-go/services/wallet/router/filter_test.go:512<br/>Error: Not equal: <br/>expected: [][]*router.PathV2{}<br/>actual: [][]*router.PathV2(nil)<br/><br/>Diff:<br/>--- Expected<br/>+++ Actual<br/>@@ -1,3 +1,2 @@<br/>-([][]*router.PathV2) {<br/>-}<br/>+([][]*router.PathV2) <nil><br/><br/>Test: TestFilterNetworkComplianceV2/Empty_routes<br/>--- FAIL: TestFilterNetworkComplianceV2/Empty_routes (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>✅ passed</em> No_locked_amounts</summary>
        <pre>=== RUN   TestFilterNetworkComplianceV2/No_locked_amounts<br/>--- PASS: TestFilterNetworkComplianceV2/No_locked_amounts (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>❌ failed</em> Single_route_with_mixed_valid_and_invalid_paths</summary>
        <pre>=== RUN   TestFilterNetworkComplianceV2/Single_route_with_mixed_valid_and_invalid_paths<br/>filter_test.go:512: <br/>Error Trace: /Users/samuel/go/src/github.com/Samyoul/status-go/services/wallet/router/filter_test.go:512<br/>Error: Not equal: <br/>expected: [][]*router.PathV2{[]*router.PathV2{(*router.PathV2)(0xc0004feea0), (*router.PathV2)(0xc0004fef70)}}<br/>actual: [][]*router.PathV2(nil)<br/><br/>Diff:<br/>--- Expected<br/>+++ Actual<br/>@@ -1,95 +1,2 @@<br/>-([][]*router.PathV2) (len=1) {<br/>-([]*router.PathV2) (len=2) {<br/>-(*router.PathV2)({<br/>-BridgeName: (string) "",<br/>-From: (*params.Network)({<br/>-ChainID: (uint64) 1,<br/>-ChainName: (string) "",<br/>-RPCURL: (string) "",<br/>-OriginalRPCURL: (string) "",<br/>-FallbackURL: (string) "",<br/>-OriginalFallbackURL: (string) "",<br/>-BlockExplorerURL: (string) "",<br/>-IconURL: (string) "",<br/>-NativeCurrencyName: (string) "",<br/>-NativeCurrencySymbol: (string) "",<br/>-NativeCurrencyDecimals: (uint64) 0,<br/>-IsTest: (bool) false,<br/>-Layer: (uint64) 0,<br/>-Enabled: (bool) false,<br/>-ChainColor: (string) "",<br/>-ShortName: (string) "",<br/>-TokenOverrides: ([]params.TokenOverride) <nil>,<br/>-RelatedChainID: (uint64) 0<br/>-}),<br/>-To: (*params.Network)(<nil>),<br/>-FromToken: (*token.Token)(<nil>),<br/>-AmountIn: (*hexutil.Big)(<nil>),<br/>-AmountInLocked: (bool) false,<br/>-AmountOut: (*hexutil.Big)(<nil>),<br/>-SuggestedPriorityFees: (*router.PriorityFees)(<nil>),<br/>-TxBaseFee: (*hexutil.Big)(<nil>),<br/>-TxPriorityFee: (*hexutil.Big)(<nil>),<br/>-TxGasAmount: (uint64) 0,<br/>-TxBonderFees: (*hexutil.Big)(<nil>),<br/>-TxTokenFees: (*hexutil.Big)(<nil>),<br/>-TxL1Fee: (*hexutil.Big)(<nil>),<br/>-ApprovalRequired: (bool) false,<br/>-ApprovalAmountRequired: (*hexutil.Big)(<nil>),<br/>-ApprovalContractAddress: (*common.Address)(<nil>),<br/>-ApprovalBaseFee: (*hexutil.Big)(<nil>),<br/>-ApprovalPriorityFee: (*hexutil.Big)(<nil>),<br/>-ApprovalGasAmount: (uint64) 0,<br/>-ApprovalL1Fee: (*hexutil.Big)(<nil>),<br/>-EstimatedTime: (router.TransactionEstimation) 0,<br/>-requiredTokenBalance: (*big.Int)(<nil>),<br/>-requiredNativeBalance: (*big.Int)(<nil>)<br/>-}),<br/>-(*router.PathV2)({<br/>-BridgeName: (string) "",<br/>-From: (*params.Network)({<br/>-ChainID: (uint64) 3,<br/>-ChainName: (string) "",<br/>-RPCURL: (string) "",<br/>-OriginalRPCURL: (string) "",<br/>-FallbackURL: (string) "",<br/>-OriginalFallbackURL: (string) "",<br/>-BlockExplorerURL: (string) "",<br/>-IconURL: (string) "",<br/>-NativeCurrencyName: (string) "",<br/>-NativeCurrencySymbol: (string) "",<br/>-NativeCurrencyDecimals: (uint64) 0,<br/>-IsTest: (bool) false,<br/>-Layer: (uint64) 0,<br/>-Enabled: (bool) false,<br/>-ChainColor: (string) "",<br/>-ShortName: (string) "",<br/>-TokenOverrides: ([]params.TokenOverride) <nil>,<br/>-RelatedChainID: (uint64) 0<br/>-}),<br/>-To: (*params.Network)(<nil>),<br/>-FromToken: (*token.Token)(<nil>),<br/>-AmountIn: (*hexutil.Big)(<nil>),<br/>-AmountInLocked: (bool) false,<br/>-AmountOut: (*hexutil.Big)(<nil>),<br/>-SuggestedPriorityFees: (*router.PriorityFees)(<nil>),<br/>-TxBaseFee: (*hexutil.Big)(<nil>),<br/>-TxPriorityFee: (*hexutil.Big)(<nil>),<br/>-TxGasAmount: (uint64) 0,<br/>-TxBonderFees: (*hexutil.Big)(<nil>),<br/>-TxTokenFees: (*hexutil.Big)(<nil>),<br/>-TxL1Fee: (*hexutil.Big)(<nil>),<br/>-ApprovalRequired: (bool) false,<br/>-ApprovalAmountRequired: (*hexutil.Big)(<nil>),<br/>-ApprovalContractAddress: (*common.Address)(<nil>),<br/>-ApprovalBaseFee: (*hexutil.Big)(<nil>),<br/>-ApprovalPriorityFee: (*hexutil.Big)(<nil>),<br/>-ApprovalGasAmount: (uint64) 0,<br/>-ApprovalL1Fee: (*hexutil.Big)(<nil>),<br/>-EstimatedTime: (router.TransactionEstimation) 0,<br/>-requiredTokenBalance: (*big.Int)(<nil>),<br/>-requiredNativeBalance: (*big.Int)(<nil>)<br/>-})<br/>-}<br/>-}<br/>+([][]*router.PathV2) <nil><br/><br/>Test: TestFilterNetworkComplianceV2/Single_route_with_mixed_valid_and_invalid_paths<br/>--- FAIL: TestFilterNetworkComplianceV2/Single_route_with_mixed_valid_and_invalid_paths (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>✅ passed</em> Routes_with_duplicate_chain_IDs</summary>
        <pre>=== RUN   TestFilterNetworkComplianceV2/Routes_with_duplicate_chain_IDs<br/>--- PASS: TestFilterNetworkComplianceV2/Routes_with_duplicate_chain_IDs (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>✅ passed</em> Minimum_and_maximum_chain_IDs</summary>
        <pre>=== RUN   TestFilterNetworkComplianceV2/Minimum_and_maximum_chain_IDs<br/>--- PASS: TestFilterNetworkComplianceV2/Minimum_and_maximum_chain_IDs (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><strong>120 ms</strong> <em>❌ failed</em> Large_number_of_routes</summary>
        <pre>=== RUN   TestFilterNetworkComplianceV2/Large_number_of_routes<br/>filter_test.go:512: <br/>Error Trace: /Users/samuel/go/src/github.com/Samyoul/status-go/services/wallet/router/filter_test.go:512<br/>Error: Not equal: <br/>expected: [][]*router.PathV2{[]*router.PathV2{(*router.PathV2)(0xc000769380), (*router.PathV2)(0xc000769450)}, []*router.PathV2{(*router.PathV2)(0xc000769520), (*router.PathV2)(0xc0007695f0)},... truncated}<br/><br/>Diff:<br/>--- Expected<br/>+++ Actual<br/>@@ -1,830 +1,2 @@<br/>-([][]*router.PathV2) (len=1000) {<br/>-([]*router.PathV2) (len=2) {<br/>(*router.PathV2)({<br/>-BridgeName: (string) "",<br/>-From: (*params.Network)({<br/>-ChainID: (uint64) 1,<br/>-ChainName: (string) "",<br/>-RPCURL: (string) "",<br/>-OriginalRPCURL: (string) "",<br/>-FallbackURL: (string) "",<br/>-OriginalFallbackURL: (string) "",<br/>-BlockExplorerURL: (string) "",<br/>-IconURL: (string) "",<br/>-NativeCurrencyName: (string) "",<br/>-NativeCurrencySymbol: (string) "",<br/>-NativeCurrencyDecimals: (uint64) 0,<br/>-IsTest: (bool) false,<br/>-Layer: (uint64) 0,<br/>-Enabled: (bool) false,<br/>-ChainColor: (string) "",<br/>-ShortName: (string) "",<br/>-TokenOverrides: ([]params.TokenOverride) <nil>,<br/>-RelatedChainID: (uint64) 0<br/>-}),<br/>-To: (*params.Network)(<nil>),<br/>-FromToken: (*token.Token)(<nil>),<br/>-AmountIn: (*hexutil.Big)(<nil>),<br/>-AmountInLocked: (bool) false,<br/>-AmountOut: (*hexutil.Big)(<nil>),<br/>-SuggestedPriorityFees: (*router.PriorityFees)(<nil>),<br/>-TxBaseFee: (*hexutil.Big)(<nil>),<br/>-TxPriorityFee: (*hexutil.Big)(<nil>),<br/>-TxGasAmount: (uint64) 0,<br/>-TxBonderFees: (*hexutil.Big)(<nil>),<br/>-TxTokenFees: (*hexutil.Big)(<nil>),<br/>-TxL1Fee: (*hexutil.Big)(<nil>),<br/>-ApprovalRequired: (bool) false,<br/>-ApprovalAmountRequired: (*hexutil.Big)(<nil>),<br/>-ApprovalContractAddress: (*common.Address)(<nil>),<br/>-ApprovalBaseFee: (*hexutil.Big)(<nil>),<br/>-ApprovalPriorityFee: (*hexutil.Big)(<nil>),<br/>-ApprovalGasAmount: (uint64) 0,<br/>-ApprovalL1Fee: (*hexutil.Big)(<nil>),<br/>-EstimatedTime: (router.TransactionEstimation) 0,<br/>-requiredTokenBalance: (*big.Int)(<nil>),<br/>-requiredNativeBalance: (*big.Int)(<nil>)<br/>}),<br/>-(*router.PathV2)({<br/>-BridgeName: (string) "",<br/>-From: (*params.Network)({<br/>-ChainID: (uint64) 1001,<br/>-ChainName: (string) "",<br/>-RPCURL: (string) "",<br/>-OriginalRPCURL: (string) "",<br/>-FallbackURL: (string) "",<br/>-OriginalFallbackURL: (string) "",<br/>-BlockExplorerURL: (string) "",<br/>-IconURL: (string) "",<br/>-NativeCurrencyName: (string) "",<br/>-NativeCurrencySymbol: (string) "",<br/>-NativeCurrencyDecimals: (uint64) 0,<br/>-IsTest: (bool) false,<br/>-Layer: (uint64) 0,<br/>-Enabled: (bool) false,<br/>-ChainColor: (string) "",<br/>-ShortName: (string) "",<br/>-TokenOverrides: ([]params.TokenOverride) <nil>,<br/>-RelatedChainID: (uint64) 0<br/>-}),<br/>-To: (*params.Network)(<nil>),<br/>-FromToken: (*token.Token)(<nil>),<br/>-AmountIn: (*hexutil.Big)(<nil>),<br/>-AmountInLocked: (bool) false,<br/>-AmountOut: (*hexutil.Big)(<nil>),<br/>-SuggestedPriorityFees: (*router.PriorityFees)(<nil>),<br/>-TxBaseFee: (*hexutil.Big)(<nil>),<br/>-TxPriorityFee: (*hexutil.Big)(<nil>),<br/>-TxGasAmount: (uint64) 0,<br/>-TxBonderFees: (*hexutil.Big)(<nil>),<br/>-TxTokenFees: (*hexutil.Big)(<nil>),<br/>-TxL1Fee: (*hexutil.Big)(<nil>),<br/>-ApprovalRequired: (bool) false,<br/>-ApprovalAmountRequired: (*hexutil.Big)(<nil>),<br/>-ApprovalContractAddress: (*common.Address)(<nil>),<br/>-ApprovalBaseFee: (*hexutil.Big)(<nil>),<br/>-ApprovalPriorityFee: (*hexutil.Big)(<nil>),<br/>-ApprovalGasAmount: (uint64) 0,<br/>-ApprovalL1Fee: (*hexutil.Big)(<nil>),<br/>-EstimatedTime: (router.TransactionEstimation) 0,<br/>-requiredTokenBalance: (*big.Int)(<nil>),<br/>-requiredNativeBalance: (*big.Int)(<nil>)<br/>})<br/>-},...},<br/>+([][]*router.PathV2) (len=991) {<br/>([]*router.PathV2) (len=2) {<br/>Test: TestFilterNetworkComplianceV2/Large_number_of_routes<br/>--- FAIL: TestFilterNetworkComplianceV2/Large_number_of_routes (0.12s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><strong>1.67 s</strong> <em>❌ failed</em> Routes_with_missing_data</summary>
        <pre>=== RUN   TestFilterNetworkComplianceV2/Routes_with_missing_data<br/>--- FAIL: TestFilterNetworkComplianceV2/Routes_with_missing_data (0.00s)<br/>panic: runtime error: invalid memory address or nil pointer dereference [recovered]<br/>panic: runtime error: invalid memory address or nil pointer dereference<br/>[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x100c02abe]<br/>goroutine 88 [running]:<br/>testing.tRunner.func1.2({0x101588300, 0x1032265f0})<br/>/usr/local/Cellar/go/1.20.5/libexec/src/testing/testing.go:1526 +0x24e<br/>testing.tRunner.func1()<br/>/usr/local/Cellar/go/1.20.5/libexec/src/testing/testing.go:1529 +0x39f<br/>panic({0x101588300, 0x1032265f0})<br/>/usr/local/Cellar/go/1.20.5/libexec/src/runtime/panic.go:884 +0x213<br/>github.com/status-im/status-go/services/wallet/router.isValidForNetworkComplianceV2(...)<br/>/Users/samuel/go/src/github.com/Samyoul/status-go/services/wallet/router/filter.go:34<br/>github.com/status-im/status-go/services/wallet/router.filterNetworkComplianceV2({0xc000168360, 0x2, 0x10005dcdf?}, 0x59a?)<br/>/Users/samuel/go/src/github.com/Samyoul/status-go/services/wallet/router/filter.go:24 +0x25e<br/>github.com/status-im/status-go/services/wallet/router.TestFilterNetworkComplianceV2.func3(0xc000503860?)<br/>/Users/samuel/go/src/github.com/Samyoul/status-go/services/wallet/router/filter_test.go:511 +0x3e<br/>testing.tRunner(0xc000502680, 0xc0007f95e0)<br/>/usr/local/Cellar/go/1.20.5/libexec/src/testing/testing.go:1576 +0x10b<br/>created by testing.(*T).Run<br/>/usr/local/Cellar/go/1.20.5/libexec/src/testing/testing.go:1629 +0x3ea<br/></pre>
      </details>
    </li>
  </ul>
</details>
<details>
  <summary><strong>10 ms</strong> ❌ TestFilterCapacityValidationV2</summary>
  <ul>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>❌ failed</em> Sufficient_capacity_with_multiple_paths</summary>
        <pre>=== RUN   TestFilterCapacityValidationV2/Sufficient_capacity_with_multiple_paths<br/>filter_test.go:722: <br/>Error Trace: /Users/samuel/go/src/github.com/Samyoul/status-go/services/wallet/router/filter_test.go:722<br/>Error: Not equal: <br/>expected: [][]*router.PathV2{[]*router.PathV2{(*router.PathV2)(0xc00058dd40), (*router.PathV2)(0xc00058de10)}}<br/>actual: [][]*router.PathV2{[]*router.PathV2{(*router.PathV2)(0xc00058da00), (*router.PathV2)(0xc00058dad0)}, []*router.PathV2{(*router.PathV2)(0xc00058dba0), (*router.PathV2)(0xc00058dc70)}}<br/><br/>Diff:<br/>--- Expected<br/>+++ Actual<br/>@@ -1,2 +1,2 @@<br/>-([][]*router.PathV2) (len=1) {<br/>+([][]*router.PathV2) (len=2) {<br/>([]*router.PathV2) (len=2) {<br/>@@ -29,2 +29,52 @@<br/>abs: (big.nat) (len=1) {<br/>+ (big.Word) 50<br/>} })<br/>+ AmountInLocked: (bool) true,<br/>+ AmountOut: (*hexutil.Big)(<nil>),<br/>+ SuggestedPriorityFees: (*router.PriorityFees)(<nil>),<br/>+ TxBaseFee: (*hexutil.Big)(<nil>),<br/>+ TxPriorityFee: (*hexutil.Big)(<nil>),<br/>+ TxGasAmount: (uint64) 0,<br/>+ TxBonderFees: (*hexutil.Big)(<nil>),<br/>+ TxTokenFees: (*hexutil.Big)(<nil>),<br/>+ TxL1Fee: (*hexutil.Big)(<nil>),<br/>+ ApprovalRequired: (bool) false,<br/>+ ApprovalAmountRequired: (*hexutil.Big)(<nil>),<br/>+ ApprovalContractAddress: (*common.Address)(<nil>),<br/>+ ApprovalBaseFee: (*hexutil.Big)(<nil>),<br/>+ ApprovalPriorityFee: (*hexutil.Big)(<nil>),<br/>+ ApprovalGasAmount: (uint64) 0,<br/>+ ApprovalL1Fee: (*hexutil.Big)(<nil>),<br/>+ EstimatedTime: (router.TransactionEstimation) 0,<br/>+ requiredTokenBalance: (*big.Int)(<nil>),<br/>+ requiredNativeBalance: (*big.Int)(<nil>)<br/>}),<br/>+ (*router.PathV2)({<br/>+ BridgeName: (string) "",<br/>+ From: (*params.Network)({<br/>+ ChainID: (uint64) 2,<br/>+ ChainName: (string) "",<br/>+ RPCURL: (string) "",<br/>+ OriginalRPCURL: (string) "",<br/>+ FallbackURL: (string) "",<br/>+ OriginalFallbackURL: (string) "",<br/>+ BlockExplorerURL: (string) "",<br/>+ IconURL: (string) "",<br/>+ NativeCurrencyName: (string) "",<br/>+ NativeCurrencySymbol: (string) "",<br/>+ NativeCurrencyDecimals: (uint64) 0,<br/>+ IsTest: (bool) false,<br/>+ Layer: (uint64) 0,<br/>+ Enabled: (bool) false,<br/>+ ChainColor: (string) "",<br/>+ ShortName: (string) "",<br/>+ TokenOverrides: ([]params.TokenOverride) <nil>,<br/>+ RelatedChainID: (uint64) 0<br/>+ }),<br/>+ To: (*params.Network)(<nil>),<br/>+ FromToken: (*token.Token)(<nil>),<br/>+ AmountIn: (*hexutil.Big)({<br/>+ neg: (bool) false,<br/>+ abs: (big.nat) (len=1) {<br/>+ (big.Word) 100<br/>+ } }),<br/>+ AmountInLocked: (bool) true,<br/>+ AmountOut: (*hexutil.Big)(<nil>),<br/>+ SuggestedPriorityFees: (*router.PriorityFees)(<nil>),<br/>+ TxBaseFee: (*hexutil.Big)(<nil>),<br/>+ TxPriorityFee: (*hexutil.Big)(<nil>),<br/>+ TxGasAmount: (uint64) 0,<br/>+ TxBonderFees: (*hexutil.Big)(<nil>),<br/>+ TxTokenFees: (*hexutil.Big)(<nil>),<br/>+ TxL1Fee: (*hexutil.Big)(<nil>),<br/>+ ApprovalRequired: (bool) false,<br/>+ ApprovalAmountRequired: (*hexutil.Big)(<nil>),<br/>+ ApprovalContractAddress: (*common.Address)(<nil>),<br/>+ ApprovalBaseFee: (*hexutil.Big)(<nil>),<br/>+ ApprovalPriorityFee: (*hexutil.Big)(<nil>),<br/>+ ApprovalGasAmount: (uint64) 0,<br/>+ ApprovalL1Fee: (*hexutil.Big)(<nil>),<br/>+ EstimatedTime: (router.TransactionEstimation) 0,<br/>+ requiredTokenBalance: (*big.Int)(<nil>),<br/>+ requiredNativeBalance: (*big.Int)(<nil>)<br/>})<br/>+ },<br/>+ ([]*router.PathV2) (len=2) {<br/>+ (*router.PathV2)({<br/>+ BridgeName: (string) "",<br/>+ From: (*params.Network)({<br/>+ ChainID: (uint64) 1,<br/>+ ChainName: (string) "",<br/>+ RPCURL: (string) "",<br/>+ OriginalRPCURL: (string) "",<br/>+ FallbackURL: (string) "",<br/>+ OriginalFallbackURL: (string) "",<br/>+ BlockExplorerURL: (string) "",<br/>+ IconURL: (string) "",<br/>+ NativeCurrencyName: (string) "",<br/>+ NativeCurrencySymbol: (string) "",<br/>+ NativeCurrencyDecimals: (uint64) 0,<br/>+ IsTest: (bool) false,<br/>+ Layer: (uint64) 0,<br/>+ Enabled: (bool) false,<br/>+ ChainColor: (string) "",<br/>+ ShortName: (string) "",<br/>+ TokenOverrides: ([]params.TokenOverride) <nil>,<br/>+ RelatedChainID: (uint64) 0<br/>+ }),<br/>+ To: (*params.Network)(<nil>),<br/>+ FromToken: (*token.Token)(<nil>),<br/>+ AmountIn: (*hexutil.Big)({<br/>+ neg: (bool) false,<br/>+ abs: (big.nat) (len=1) {<br/>+ (big.Word) 50<br/>+ } }),<br/>+ AmountInLocked: (bool) true,<br/>AmountOut: (*hexutil.Big)(<nil>),<br/>@@ -79,6 +181,6 @@<br/>abs: (big.nat) (len=1) {<br/>- (big.Word) 200<br/>- }<br/>- }),<br/>- AmountInLocked: (bool) false,<br/>+ (big.Word) 100<br/>+ } }),<br/>+ AmountInLocked: (bool) true,<br/>AmountOut: (*hexutil.Big)(<nil>),<br/>Test: TestFilterCapacityValidationV2/Sufficient_capacity_with_multiple_paths<br/>--- FAIL: TestFilterCapacityValidationV2/Sufficient_capacity_with_multiple_paths (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>❌ failed</em> Insufficient_capacity</summary>
        <pre>=== RUN   TestFilterCapacityValidationV2/Insufficient_capacity<br/>filter_test.go:722: <br/>Error Trace: /Users/samuel/go/src/github.com/Samyoul/status-go/services/wallet/router/filter_test.go:722<br/>Error: Not equal: <br/>expected: [][]*router.PathV2{}<br/>actual: [][]*router.PathV2(nil)<br/><br/>Diff:<br/>--- Expected<br/>+++ Actual<br/>@@ -1,3 +1,2 @@<br/>-([][]*router.PathV2) {<br/>-}<br/>+([][]*router.PathV2) <nil><br/><br/>Test: TestFilterCapacityValidationV2/Insufficient_capacity<br/>--- FAIL: TestFilterCapacityValidationV2/Insufficient_capacity (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>❌ failed</em> Exact_capacity_match</summary>
        <pre>=== RUN   TestFilterCapacityValidationV2/Exact_capacity_match<br/>filter_test.go:722: <br/>Error Trace: /Users/samuel/go/src/github.com/Samyoul/status-go/services/wallet/router/filter_test.go:722<br/>Error: Not equal: <br/>expected: [][]*router.PathV2{[]*router.PathV2{(*router.PathV2)(0xc0000ad6c0), (*router.PathV2)(0xc0000ad860)}}<br/>actual: [][]*router.PathV2(nil)<br/><br/>Diff:<br/>--- Expected<br/>+++ Actual<br/>@@ -1,105 +1,2 @@<br/>-([][]*router.PathV2) (len=1) {<br/>-([]*router.PathV2) (len=2) {<br/>- (*router.PathV2)({<br/>- BridgeName: (string) "",<br/>- From: (*params.Network)({<br/>- ChainID: (uint64) 1,<br/>- ChainName: (string) "",<br/>- RPCURL: (string) "",<br/>- OriginalRPCURL: (string) "",<br/>- FallbackURL: (string) "",<br/>- OriginalFallbackURL: (string) "",<br/>- BlockExplorerURL: (string) "",<br/>- IconURL: (string) "",<br/>- NativeCurrencyName: (string) "",<br/>- NativeCurrencySymbol: (string) "",<br/>- NativeCurrencyDecimals: (uint64) 0,<br/>- IsTest: (bool) false,<br/>- Layer: (uint64) 0,<br/>- Enabled: (bool) false,<br/>- ChainColor: (string) "",<br/>- ShortName: (string) "",<br/>- TokenOverrides: ([]params.TokenOverride) <nil>,<br/>- RelatedChainID: (uint64) 0<br/>- }),<br/>- To: (*params.Network)(<nil>),<br/>- FromToken: (*token.Token)(<nil>),<br/>- AmountIn: (*hexutil.Big)({<br/>- neg: (bool) false,<br/>- abs: (big.nat) (len=1) {<br/>- (big.Word) 100<br/>- } }),<br/>- AmountInLocked: (bool) true,<br/>- AmountOut: (*hexutil.Big)(<nil>),<br/>- SuggestedPriorityFees: (*router.PriorityFees)(<nil>),<br/>- TxBaseFee: (*hexutil.Big)(<nil>),<br/>- TxPriorityFee: (*hexutil.Big)(<nil>),<br/>- TxGasAmount: (uint64) 0,<br/>- TxBonderFees: (*hexutil.Big)(<nil>),<br/>- TxTokenFees: (*hexutil.Big)(<nil>),<br/>- TxL1Fee: (*hexutil.Big)(<nil>),<br/>- ApprovalRequired: (bool) false,<br/>- ApprovalAmountRequired: (*hexutil.Big)(<nil>),<br/>- ApprovalContractAddress: (*common.Address)(<nil>),<br/>- ApprovalBaseFee: (*hexutil.Big)(<nil>),<br/>- ApprovalPriorityFee: (*hexutil.Big)(<nil>),<br/>- ApprovalGasAmount: (uint64) 0,<br/>- ApprovalL1Fee: (*hexutil.Big)(<nil>),<br/>- EstimatedTime: (router.TransactionEstimation) 0,<br/>- requiredTokenBalance: (*big.Int)(<nil>),<br/>- requiredNativeBalance: (*big.Int)(<nil>)<br/>- }),<br/>- (*router.PathV2)({<br/>- BridgeName: (string) "",<br/>- From: (*params.Network)({<br/>- ChainID: (uint64) 2,<br/>- ChainName: (string) "",<br/>- RPCURL: (string) "",<br/>- OriginalRPCURL: (string) "",<br/>- FallbackURL: (string) "",<br/>- OriginalFallbackURL: (string) "",<br/>- BlockExplorerURL: (string) "",<br/>- IconURL: (string) "",<br/>- NativeCurrencyName: (string) "",<br/>- NativeCurrencySymbol: (string) "",<br/>- NativeCurrencyDecimals: (uint64) 0,<br/>- IsTest: (bool) false,<br/>- Layer: (uint64) 0,<br/>- Enabled: (bool) false,<br/>- ChainColor: (string) "",<br/>- ShortName: (string) "",<br/>- TokenOverrides: ([]params.TokenOverride) <nil>,<br/>- RelatedChainID: (uint64) 0<br/>- }),<br/>- To: (*params.Network)(<nil>),<br/>- FromToken: (*token.Token)(<nil>),<br/>- AmountIn: (*hexutil.Big)({<br/>- neg: (bool) false,<br/>- abs: (big.nat) (len=1) {<br/>- (big.Word) 50<br/>- } }),<br/>- AmountInLocked: (bool) true,<br/>- AmountOut: (*hexutil.Big)(<nil>),<br/>- SuggestedPriorityFees: (*router.PriorityFees)(<nil>),<br/>- TxBaseFee: (*hexutil.Big)(<nil>),<br/>- TxPriorityFee: (*hexutil.Big)(<nil>),<br/>- TxGasAmount: (uint64) 0,<br/>- TxBonderFees: (*hexutil.Big)(<nil>),<br/>- TxTokenFees: (*hexutil.Big)(<nil>),<br/>- TxL1Fee: (*hexutil.Big)(<nil>),<br/>- ApprovalRequired: (bool) false,<br/>- ApprovalAmountRequired: (*hexutil.Big)(<nil>),<br/>- ApprovalContractAddress: (*common.Address)(<nil>),<br/>- ApprovalBaseFee: (*hexutil.Big)(<nil>),<br/>- ApprovalPriorityFee: (*hexutil.Big)(<nil>),<br/>- ApprovalGasAmount: (uint64) 0,<br/>- ApprovalL1Fee: (*hexutil.Big)(<nil>),<br/>- EstimatedTime: (router.TransactionEstimation) 0,<br/>- requiredTokenBalance: (*big.Int)(<nil>),<br/>- requiredNativeBalance: (*big.Int)(<nil>)<br/>- }),<br/>- (*router.PathV2)({<br/>- BridgeName: (string) "",<br/>- From: (*params.Network)({<br/>- ChainID: (uint64) 1,<br/>- ChainName: (string) "",<br/>- RPCURL: (string) "",<br/>- OriginalRPCURL: (string) "",<br/>- FallbackURL: (string) "",<br/>- OriginalFallbackURL: (string) "",<br/>- BlockExplorerURL: (string) "",<br/>- IconURL: (string) "",<br/>- NativeCurrencyName: (string) "",<br/>- NativeCurrencySymbol: (string) "",<br/>- NativeCurrencyDecimals: (uint64) 0,<br/>- IsTest: (bool) false,<br/>- Layer: (uint64) 0,<br/>- Enabled: (bool) false,<br/>- ChainColor: (string) "",<br/>- ShortName: (string) "",<br/>- TokenOverrides: ([]params.TokenOverride) <nil>,<br/>- RelatedChainID: (uint64) 0<br/>- }),<br/>- To: (*params.Network)(<nil>),<br/>- FromToken: (*token.Token)(<nil>),<br/>- AmountIn: (*hexutil.Big)({<br/>- neg: (bool) false,<br/>- abs: (big.nat) (len=1) {<br/>- (big.Word) 50<br/>- } }),<br/>- AmountInLocked: (bool) true,<br/>AmountOut: (*hexutil.Big)(<nil>),<br/>SuggestedPriorityFees: (*router.PriorityFees)(<nil>),<br/>TxBaseFee: (*hexutil.Big)(<nil>),<br/>TxPriorityFee: (*hexutil.Big)(<nil>),<br/>TxGasAmount: (uint64) 0,<br/>TxBonderFees: (*hexutil.Big)(<nil>),<br/>TxTokenFees: (*hexutil.Big)(<nil>),<br/>TxL1Fee: (*hexutil.Big)(<nil>),<br/>ApprovalRequired: (bool) false,<br/>ApprovalAmountRequired: (*hexutil.Big)(<nil>),<br/>ApprovalContractAddress: (*common.Address)(<nil>),<br/>ApprovalBaseFee: (*hexutil.Big)(<nil>),<br/>ApprovalPriorityFee: (*hexutil.Big)(<nil>),<br/>ApprovalGasAmount: (uint64) 0,<br/>ApprovalL1Fee: (*hexutil.Big)(<nil>),<br/>EstimatedTime: (router.TransactionEstimation) 0,<br/>requiredTokenBalance: (*big.Int)(<nil>),<br/>requiredNativeBalance: (*big.Int)(<nil>)<br/>})<br/>+([][]*router.PathV2) <nil><br/><br/>Test: TestFilterCapacityValidationV2/Exact_capacity_match<br/>--- FAIL: TestFilterCapacityValidationV2/Exact_capacity_match (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>✅ passed</em> No_locked_amounts</summary>
        <pre>=== RUN   TestFilterCapacityValidationV2/No_locked_amounts<br/>--- PASS: TestFilterCapacityValidationV2/No_locked_amounts (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>❌ failed</em> Single_route_with_sufficient_capacity</summary>
        <pre>=== RUN   TestFilterCapacityValidationV2/Single_route_with_sufficient_capacity<br/>filter_test.go:722: <br/>Error Trace: /Users/samuel/go/src/github.com/Samyoul/status-go/services/wallet/router/filter_test.go:722<br/>Error: Not equal: <br/>expected: [][]*router.PathV2{[]*router.PathV2{(*router.PathV2)(0xc000614680)}}<br/>actual: [][]*router.PathV2{[]*router.PathV2{(*router.PathV2)(0xc0000adee0)}}<br/><br/>Diff:<br/>--- Expected<br/>+++ Actual<br/>@@ -29,6 +29,6 @@<br/>abs: (big.nat) (len=1) {<br/>- (big.Word) 200<br/>+ (big.Word) 50<br/>} }),<br/>- AmountInLocked: (bool) false,<br/>+ AmountInLocked: (bool) true,<br/>AmountOut: (*hexutil.Big)(<nil>),<br/>Test: TestFilterCapacityValidationV2/Single_route_with_sufficient_capacity<br/>--- FAIL: TestFilterCapacityValidationV2/Single_route_with_sufficient_capacity (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>❌ failed</em> Single_route_with_insufficient_capacity</summary>
        <pre>=== RUN   TestFilterCapacityValidationV2/Single_route_with_insufficient_capacity<br/>filter_test.go:722: <br/>Error Trace: /Users/samuel/go/src/github.com/Samyoul/status-go/services/wallet/router/filter_test.go:722<br/>Error: Not equal: <br/>expected: [][]*router.PathV2{}<br/>actual: [][]*router.PathV2{[]*router.PathV2{(*router.PathV2)(0xc000614750)}}<br/><br/>Diff:<br/>--- Expected<br/>+++ Actual<br/>@@ -1,2 +1,54 @@<br/>-([][]*router.PathV2) {<br/>+([][]*router.PathV2) (len=1) {<br/>([]*router.PathV2) (len=1) {<br/> (*router.PathV2)({<br/> BridgeName: (string) "",<br/> From: (*params.Network)({<br/> ChainID: (uint64) 1,<br/> ChainName: (string) "",<br/> RPCURL: (string) "",<br/> OriginalRPCURL: (string) "",<br/> FallbackURL: (string) "",<br/> OriginalFallbackURL: (string) "",<br/> BlockExplorerURL: (string) "",<br/> IconURL: (string) "",<br/> NativeCurrencyName: (string) "",<br/> NativeCurrencySymbol: (string) "",<br/> NativeCurrencyDecimals: (uint64) 0,<br/> IsTest: (bool) false,<br/> Layer: (uint64) 0,<br/> Enabled: (bool) false,<br/> ChainColor: (string) "",<br/> ShortName: (string) "",<br/> TokenOverrides: ([]params.TokenOverride) <nil>,<br/> RelatedChainID: (uint64) 0<br/> }),<br/> To: (*params.Network)(<nil>),<br/> FromToken: (*token.Token)(<nil>),<br/> AmountIn: (*hexutil.Big)({<br/> neg: (bool) false,<br/> abs: (big.nat) (len=1) {<br/> (big.Word) 50<br/> } }),<br/> AmountInLocked: (bool) true,<br/> AmountOut: (*hexutil.Big)(<nil>),<br/> SuggestedPriorityFees: (*router.PriorityFees)(<nil>),<br/> TxBaseFee: (*hexutil.Big)(<nil>),<br/> TxPriorityFee: (*hexutil.Big)(<nil>),<br/> TxGasAmount: (uint64) 0,<br/> TxBonderFees: (*hexutil.Big)(<nil>),<br/> TxTokenFees: (*hexutil.Big)(<nil>),<br/> TxL1Fee: (*hexutil.Big)(<nil>),<br/> ApprovalRequired: (bool) false,<br/> ApprovalAmountRequired: (*hexutil.Big)(<nil>),<br/> ApprovalContractAddress: (*common.Address)(<nil>),<br/> ApprovalBaseFee: (*hexutil.Big)(<nil>),<br/> ApprovalPriorityFee: (*hexutil.Big)(<nil>),<br/> ApprovalGasAmount: (uint64) 0,<br/> ApprovalL1Fee: (*hexutil.Big)(<nil>),<br/> EstimatedTime: (router.TransactionEstimation) 0,<br/> requiredTokenBalance: (*big.Int)(<nil>),<br/> requiredNativeBalance: (*big.Int)(<nil>)<br/> })<br/> } }<br/>Test: TestFilterCapacityValidationV2/Single_route_with_insufficient_capacity<br/>--- FAIL: TestFilterCapacityValidationV2/Single_route_with_insufficient_capacity (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>❌ failed</em> Empty_routes</summary>
        <pre>=== RUN   TestFilterCapacityValidationV2/Empty_routes<br/>filter_test.go:722: <br/>Error Trace: /Users/samuel/go/src/github.com/Samyoul/status-go/services/wallet/router/filter_test.go:722<br/>Error: Not equal: <br/>expected: [][]*router.PathV2{}<br/>actual: [][]*router.PathV2(nil)<br/><br/>Diff:<br/>--- Expected<br/>+++ Actual<br/>@@ -1,3 +1,2 @@<br/>-([][]*router.PathV2) {<br/>-}<br/>+([][]*router.PathV2) <nil><br/><br/>Test: TestFilterCapacityValidationV2/Empty_routes<br/>--- FAIL: TestFilterCapacityValidationV2/Empty_routes (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>❌ failed</em> Routes_with_duplicate_chain_IDs</summary>
        <pre>=== RUN   TestFilterCapacityValidationV2/Routes_with_duplicate_chain_IDs<br/>filter_test.go:722: <br/>Error Trace: /Users/samuel/go/src/github.com/Samyoul/status-go/services/wallet/router/filter_test.go:722<br/>Error: Not equal: <br/>expected: [][]*router.PathV2{[]*router.PathV2{(*router.PathV2)(0xc0006149c0), (*router.PathV2)(0xc000614a90)}}<br/>actual: [][]*router.PathV2{[]*router.PathV2{(*router.PathV2)(0xc000614820), (*router.PathV2)(0xc0006148f0)}}<br/><br/>Diff:<br/>--- Expected<br/>+++ Actual<br/>@@ -29,6 +29,6 @@<br/>abs: (big.nat) (len=1) {<br/>- (big.Word) 100<br/>+ (big.Word) 50<br/>} }),<br/>- AmountInLocked: (bool) false,<br/>+ AmountInLocked: (bool) true,<br/>AmountOut: (*hexutil.Big)(<nil>),<br/>Test: TestFilterCapacityValidationV2/Routes_with_duplicate_chain_IDs<br/>--- FAIL: TestFilterCapacityValidationV2/Routes_with_duplicate_chain_IDs (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>❌ failed</em> Partial_locked_amounts</summary>
        <pre>=== RUN   TestFilterCapacityValidationV2/Partial_locked_amounts<br/>filter_test.go:722: <br/>Error Trace: /Users/samuel/go/src/github.com/Samyoul/status-go/services/wallet/router/filter_test.go:722<br/>Error: Not equal: <br/>expected: [][]*router.PathV2{[]*router.PathV2{(*router.PathV2)(0xc000614dd0), (*router.PathV2)(0xc000614ea0), (*router.PathV2)(0xc000614f70)}}<br/>actual: [][]*router.PathV2{[]*router.PathV2{(*router.PathV2)(0xc000614b60), (*router.PathV2)(0xc000614c30), (*router.PathV2)(0xc000614d00)}}<br/><br/>Diff:<br/>--- Expected<br/>+++ Actual<br/>@@ -29,3 +29,3 @@<br/>abs: (big.nat) (len=1) {<br/>- (big.Word) 100<br/>+ (big.Word) 50<br/>} })<br/>@@ -78,7 +78,5 @@<br/>neg: (bool) false,<br/>- abs: (big.nat) (len=1) {<br/>- (big.Word) 100<br/>- }<br/>+ abs: (big.nat) <nil> }),<br/>- AmountInLocked: (bool) false,<br/>+ AmountInLocked: (bool) true,<br/>AmountOut: (*hexutil.Big)(<nil>),<br/>@@ -129,3 +127,3 @@<br/>abs: (big.nat) (len=1) {<br/>- (big.Word) 200<br/>+ (big.Word) 100<br/>} })<br/>Test: TestFilterCapacityValidationV2/Partial_locked_amounts<br/>--- FAIL: TestFilterCapacityValidationV2/Partial_locked_amounts (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>✅ passed</em> Mixed_networks_with_sufficient_capacity</summary>
        <pre>=== RUN   TestFilterCapacityValidationV2/Mixed_networks_with_sufficient_capacity<br/>--- PASS: TestFilterCapacityValidationV2/Mixed_networks_with_sufficient_capacity (0.00s)<br/></pre>
      </details>
    </li>
    <li>
      <details>
        <summary><strong>0 ms</strong> <em>❌ failed</em> Mixed_networks_with_insufficient_capacity</summary>
        <pre>=== RUN   TestFilterCapacityValidationV2/Mixed_networks_with_insufficient_capacity<br/>filter_test.go:722: <br/>Error Trace: /Users/samuel/go/src/github.com/Samyoul/status-go/services/wallet/router/filter_test.go:722<br/>Error: Not equal: <br/>expected: [][]*router.PathV2{}<br/>actual: [][]*router.PathV2{[]*router.PathV2{(*router.PathV2)(0xc000615380), (*router.PathV2)(0xc000615450)}}<br/><br/>Diff:<br/>--- Expected<br/>+++ Actual<br/>@@ -1,2 +1,104 @@<br/>-([][]*router.PathV2) {<br/>+([][]*router.PathV2) (len=1) {<br/>([]*router.PathV2) (len=2) {<br/> (*router.PathV2)({<br/> BridgeName: (string) "",<br/> From: (*params.Network)({<br/> ChainID: (uint64) 1,<br/> ChainName: (string) "",<br/> RPCURL: (string) "",<br/> OriginalRPCURL: (string) "",<br/> FallbackURL: (string) "",<br/> OriginalFallbackURL: (string) "",<br/> BlockExplorerURL: (string) "",<br/> IconURL: (string) "",<br/> NativeCurrencyName: (string) "",<br/> NativeCurrencySymbol: (string) "",<br/> NativeCurrencyDecimals: (uint64) 0,<br/> IsTest: (bool) false,<br/> Layer: (uint64) 0,<br/> Enabled: (bool) false,<br/> ChainColor: (string) "",<br/> ShortName: (string) "",<br/> TokenOverrides: ([]params.TokenOverride) <nil>,<br/> RelatedChainID: (uint64) 0<br/> }),<br/> To: (*params.Network)(<nil>),<br/> FromToken: (*token.Token)(<nil>),<br/> AmountIn: (*hexutil.Big)({<br/> neg: (bool) false,<br/> abs: (big.nat) (len=1) {<br/> (big.Word) 50<br/> } }),<br/> AmountInLocked: (bool) true,<br/> AmountOut: (*hexutil.Big)(<nil>),<br/> SuggestedPriorityFees: (*router.PriorityFees)(<nil>),<br/> TxBaseFee: (*hexutil.Big)(<nil>),<br/> TxPriorityFee: (*hexutil.Big)(<nil>),<br/> TxGasAmount: (uint64) 0,<br/> TxBonderFees: (*hexutil.Big)(<nil>),<br/> TxTokenFees: (*hexutil.Big)(<nil>),<br/> TxL1Fee: (*hexutil.Big)(<nil>),<br/> ApprovalRequired: (bool) false,<br/> ApprovalAmountRequired: (*hexutil.Big)(<nil>),<br/> ApprovalContractAddress: (*common.Address)(<nil>),<br/> ApprovalBaseFee: (*hexutil.Big)(<nil>),<br/> ApprovalPriorityFee: (*hexutil.Big)(<nil>),<br/> ApprovalGasAmount: (uint64) 0,<br/> ApprovalL1Fee: (*hexutil.Big)(<nil>),<br/> EstimatedTime: (router.TransactionEstimation) 0,<br/> requiredTokenBalance: (*big.Int)(<nil>),<br/> requiredNativeBalance: (*big.Int)(<nil>)<br/> }),<br/> (*router.PathV2)({<br/> BridgeName: (string) "",<br/> From: (*params.Network)({<br/> ChainID: (uint64) 3,<br/> ChainName: (string) "",<br/> RPCURL: (string) "",<br/> OriginalRPCURL: (string) "",<br/> FallbackURL: (string) "",<br/> OriginalFallbackURL: (string) "",<br/> BlockExplorerURL: (string) "",<br/> IconURL: (string) "",<br/> NativeCurrencyName: (string) "",<br/> NativeCurrencySymbol: (string) "",<br/> NativeCurrencyDecimals: (uint64) 0,<br/> IsTest: (bool) false,<br/> Layer: (uint64) 0,<br/> Enabled: (bool) false,<br/> ChainColor: (string) "",<br/> ShortName: (string) "",<br/> TokenOverrides: ([]params.TokenOverride) <nil>,<br/> RelatedChainID: (uint64) 0<br/> }),<br/> To: (*params.Network)(<nil>),<br/> FromToken: (*token.Token)(<nil>),<br/> AmountIn: (*hexutil.Big)({<br/> neg: (bool) false,<br/> abs: (big.nat) (len=1) {<br/> (big.Word) 100<br/> } }),<br/> AmountInLocked: (bool) true,<br/> AmountOut: (*hexutil.Big)(<nil>),<br/> SuggestedPriorityFees: (*router.PriorityFees)(<nil>),<br/> TxBaseFee: (*hexutil.Big)(<nil>),<br/> TxPriorityFee: (*hexutil.Big)(<nil>),<br/> TxGasAmount: (uint64) 0,<br/> TxBonderFees: (*hexutil.Big)(<nil>),<br/> TxTokenFees: (*hexutil.Big)(<nil>),<br/> TxL1Fee: (*hexutil.Big)(<nil>),<br/> ApprovalRequired: (bool) false,<br/> ApprovalAmountRequired: (*hexutil.Big)(<nil>),<br/> ApprovalContractAddress: (*common.Address)(<nil>),<br/> ApprovalBaseFee: (*hexutil.Big)(<nil>),<br/> ApprovalPriorityFee: (*hexutil.Big)(<nil>),<br/> ApprovalGasAmount: (uint64) 0,<br/> ApprovalL1Fee: (*hexutil.Big)(<nil>),<br/> EstimatedTime: (router.TransactionEstimation) 0,<br/> requiredTokenBalance: (*big.Int)(<nil>),<br/> requiredNativeBalance: (*big.Int)(<nil>)<br/> }) } }<br/>Test: TestFilterCapacityValidationV2/Mixed_networks_with_insufficient_capacity<br/>--- FAIL: TestFilterCapacityValidationV2/Mixed_networks_with_insufficient_capacity (0.00s)<br/></pre>
      </details>
    </li>
  </ul>
</details>
