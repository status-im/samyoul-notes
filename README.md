# 2024-05-21

## Pulls

- [Router Filter Modularisation #5177](https://github.com/status-im/status-go/pull/5177)
  - Fixes to capacity logic

## Schedule

- [x] SendModal designs review : 10:00 – 12:00
  - https://meet.google.com/qdq-wrzq-mha
  - Meeting notes
    - https://www.notion.so/Wallet-Desktop-Core-Sync-638b257c27c8493bb6e0552e0c64d969
    - https://www.figma.com/design/FkFClTCYKf83RJWoifWgoX/Wallet-v2?node-id=23480-273462&t=fubB8mAUPncc0Upv-0
  - **Attendees:**
    - Sale (Facilitator)
    - Alisher
    - Ben
    - Magnus
    - Michal C
    - Myself
  - A very in depth meeting reviewing the send flows of desktop.
  - **Key Decisions:**
    - Reduce the scope of the Send, Bridge and Swap UI designs for 2.30
    - UI and design to work together
      - Desktop UI should detail what is feasible for the 2.30 release
      - Alisher and Ben are supportive of facilitating a highly iterative framework.
      - Alisher and Ben will adapt sending flows as required
    - Michal and Sale would discuss the best process for improving the UI while at the same time being able to deliver something. Include Noelia
    - Identified distinctions between the concepts of backend `Router` and UI `Router`
      - Backend `Router` is a logic model for sending blockchain tokens. This model contains functionality that manages:
        - Calculating the best path for any given transaction requirements:
          - From address(es)
          - To address(es)
          - From network(s)
          - To network(s)
          - From token type(s)
          - To token type(s)
        - The `Router` is capable of using multiple input tokens and addresses to plan a route to multiple output addresses and tokens.
          - The `Router` determines which swaps and bridges may be required to achieve the goal
      - Backend `Router` can be called by multiple UI components:
        - Bridging
        - Sending
        - Swapping
      - Backend `Router` handles all of these transaction types internally.
      - UI `Router` is a complex high level UI component that allows a user to plan with great detail which tokens, addresses and networks they wish for a transaction (or series of transactions to take).
      - UI `Router` is part of the UI bridging functionality.
- [x] mobile-core retro : 13:00 – 14:00
  - https://meet.google.com/hci-mahx-rfr
  - Meeting document
    - https://docs.google.com/document/d/1wDJQMYj-Q7jqY_JbE0Ev5ZlnbzmGInlKp5RGU-FnLC8/edit

---

# 2024-05-20

## Pulls

- [Router Filter Modularisation #5177](https://github.com/status-im/status-go/pull/5177)
  - Resolved a number of weird corruption issues with my local repo of `status-go`, see the comments from today:
    - https://github.com/status-im/status-go/pull/5177#issuecomment-2120187378
  - Work on resolving some of the logic issues

## Schedule

- [x] 13:00 – 13:45 : bi-weekly status-go call
  - https://meet.google.com/gbq-tyqe-vju
  - Stand-ups
  - Discussion about changes required for Wallet Router to accommodate swap functionality.
  - Discussion about the stability work done on Wallet Router.

---

# 2024-05-17

## Pulls

- [Router Filter Modularisation #5177](https://github.com/status-im/status-go/pull/5177) `created`, `commits`, `research`, `scoping`
  - I've implemented the modularisation of the `filterRoutesV2` function.
  - In addition, I've added tests to check the functionality works as expected:
    - The list is a bit huge, so look at this comment:
      - https://github.com/status-im/status-go/pull/5177#issuecomment-2118233055
  - Some tests are failing, [see here for details](./attachments/2024-05/filter_test_results.md)

# 2024-05-16

## Issues

- [Remove torrent dependencies at build time for mobile #5146](https://github.com/status-im/status-go/issues/5146) `scoped`
  - Identified where the imports are present in the code
    - https://github.com/status-im/status-go/issues/5146#issuecomment-2114657333
  - Identified a viable implementation approach.
    - https://github.com/status-im/status-go/issues/5146#issuecomment-2115244124
  - Further work needs to be done to scope how the restructuring of something as big and important as the Communities Manager.
    - Initialisation of the Communities Manager should be written in dedicated platform-centric files.
      - Calling the Torrent logic should be done via a callback function within the dedicated file.
      - As little code as possible should be in the platform-centric files.
    - Torrent logic should be stripped out and moved into a dedicated file.
      - Add build restrictions to the torrent file.

## Reviews

- https://github.com/status-im/status-go/pull/5159 `approved`, `feedback`
  - A re-review after Sale made some changes in response to all of the feedback given.
  - I've suggested one opinion based change based on function signature complexity, aside from that I've approved the PR.

## Schedule

- [x] 11:00 - 12:00 : Andrea and Samuel 1:1
  - https://meet.google.com/rhy-trwc-ybb
  - <details>
    <summary style="color:red;">Confidential</summary>
      <img src="./attachments/2024-05/200w.gif" title="Dennis Nedry - Jurassic Park" alt="&quot;Ah ah aaaah. You didn't say the magic word.&quot;"/>
    </details>
- [x] 13:00 - 13:30 : mobile-core planning
  - https://meet.google.com/azr-ppob-obc
  - Team stand-ups
  - Reviewed current milestone goals
  - Discussed next release cut timeline
  - Overview of longer term milestone goals
  - Brief discussion about how much I don't care about github issue labels.

---

# 2024-05-15

## Coding
- `filterRoutes` is a very complex function, and it is very difficult to read. I've split it apart and named the levels of filter after what they do rather when they happen.
  - [See here for details](./analysis/wallet/Router/code/filterRoutes.go)

## Reviews

- https://github.com/status-im/status-go/pull/5159
  - I gave a rather thorough review of Sale's raft PR
  - This PR attempts to resolve or more easily identify problems with the Router by reducing the surface area of the code.

## Scoping

- Wallet Router
  - Complete analysis for the `findBest` function:
    - [See here for details `findBest`](./analysis/wallet/Router/newSuggestedRoutes.md#findbest)
  - Completed analysis of the `filterRoutes` function:
    - [See here for details `filterRoutes`](./analysis/wallet/Router/newSuggestedRoutes.md#filterroutes)

---

# 2024-05-14

## Scoping

- Wallet Router
  - I've analysed the main components of the wallet router, see here:
    - [wallet/Router file analysis](./analysis/wallet/Router/README.md)
    - [wallet/Router struct usage](./analysis/wallet/Router/usage.md)
    - [wallet/Router `newSuggestedRoutes()`](./analysis/wallet/Router/newSuggestedRoutes.md)

## Schedule
- I raised an axe to the crew meetings, swung hard and true. Toppled and prostrate they settled motionless. Tonight we revel in the light of their embers!

---

# 2024-05-13

## Scoping

- Wallet Router
  - I began work on reading and building context of the Route `status-go` logic.
  - This work is tricky to track, but I have begun making notes of the logic. I plan to continue to build out a description or documentation detailing the analysis.
- Torrent dependencies in mobile build
  - https://github.com/status-im/status-go/issues/5146
  - Scope work for this issue requires more time, I've identified the imports and dependencies for implementing the backup distribution logic. 


## Schedule
- [x] 10:30 – 11:30 : Discussion and sync with Sale about planning and scoping of Wallet Router work.
  - Discussion giving context on the structure of the Router
  - Discussed work done last week in : https://github.com/status-im/status-desktop/issues/14638
  - We decided that we need to:
    - spend time understanding the code
    - determine what work can be parallelised
- [x] 12:00 – 12:30 : Wallet Router Sync
  - https://meet.google.com/amj-krdp-dta
  - **Attendees:**
    - Alisher
    - Andrea
    - Emil
    - Ivan
    - Jamie
    - Sale
    - Myself
  - **Notes**
    - Discussed what should be priorities
    - Decided the following:
      - Scope the Wallet Router to understand the logic
      - Tests should be writen to validate any changes the Router code, but we do not build tests out without introducing changes.
        - Tests can be written in the context of demonstrating / exploring how code works.
      - Alisher was meant to present the product expectations, but we ran out of time.
- [x] 14:30 – 15:00 : team update
  - https://meet.google.com/vnx-agcb-nao
  - A very sombre meeting, discussing the recent layoffs
