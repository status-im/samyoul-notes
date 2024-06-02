# 2024-05-31

## Pulls

- [Removing torrent from mobile build #5257](https://github.com/status-im/status-go/pull/5257) - `commits`, `scoping`
  - I've fully split Manager from `TorrentManager`
    - I've removed any mention or dependency of `TorrentManager` from `Manager`. There is still more work to do, but `Messenger` now communicates directly with a `TorrentManager` rather than asking the `communities` `Manager` to handle it.
    - I've ensured that `LogStdout()` is only called from `TorrentManager` and removed entirely from `Manager`, this functionality seems to be some kind of debug tool specifically for torrent related functionality.
    - Next I need to focus on functions within `Messenger` that call a `TorrentManager` and see how to isolate these from the main flows.
    - I also need fix the tests that are broken.
    - I also need to refactor `torrentClientReady()` so that it is a function of `TorrentManager`, this may allow for pushing more functions into `TorrentManager` which will lead to better torrent encapsulation.
  - Made `torrentClientReady()` belong to `TorrentManager`
  - Unexported all exported funcs not used externally in `TorrentManager`
  - Mapped for restructure torrent specific `Messenger` functionality
    - https://github.com/status-im/status-go/pull/5257#issuecomment-2142460289
  - Mapped entire `TorrentManager` API Usage
    - https://github.com/status-im/status-go/pull/5257#issuecomment-2142881649
    - This will augment the restructure of torrent specific `Messenger` functionality, as `Messenger` is split across multiple go files.
    - Once this is done I can identify 2 key things:
      1. What `Messenger` functionality can be pushed into `TorrentManager`
      2. What the bare minimum `TorrentManager` interface signature needs to be.

---

# 2024-05-30

## Pulls

- [Removing torrent from mobile build #5257](https://github.com/status-im/status-go/pull/5257) - `commits`, `scoping`
  - Resolved TorrentManager dep injection
    - [I mapped out the details of the deps](https://github.com/status-im/status-go/pull/5257#issuecomment-2139207967)
    - I'm not particularly proud of this work. I've just passed in all the deps as vars, there are way too many concerns handled by the TorrentManager that I don't know what is the best approach to removing them. I've even resorted to declaring an 'Publisher' interface to handle all the 'publish()' calls the TorrentManager makes. Next up I need to resolved the communities Manager API, because I've removed all of its end points. The TorrentManager needs to be a lot more simple than it is.
  - Scoping of the TorrentManager API usage
    - [mapping of the API usage here](https://github.com/status-im/status-go/pull/5257#issuecomment-2140527523)
    - This is so tightly wound into the workings of both `Messenger` and `communities`.`Manager`.

## Schedule
- [x] 11:00 - 12:00 : Andrea and Samuel 1:1
  - https://meet.google.com/rhy-trwc-ybb
  - <details>
    <summary style="color:red;">*Magic word*</summary>
      <img src="./attachments/2024-05/200w.gif" title="Dennis Nedry - Jurassic Park" alt="&quot;Ah ah aaaah. You didn't say the magic word.&quot;"/>
    </details>
- [x] 15:00 – 16:00 : mobile-core retro
  - https://meet.google.com/hci-mahx-rfr
  - Cancelled. Release season merits no retrospection!

---

# 2024-05-29

## Documents

- Implementing User Stories
  - https://www.notion.so/About-implementing-user-stories-fc6bc6b81e54461c85cd1b793cae4e22
  - I was given the privilege of previewing and reviewing this stable draft of Status' unified approach to User Stories.

## Issues

- [Remove torrent dependencies at build time for mobile #5146](https://github.com/status-im/status-go/issues/5146)
  - Begun working on this issue.

## Pulls

- [Removing torrent from mobile build #5257](https://github.com/status-im/status-go/pull/5257) - `commits`
  - Initial migration of all torrent dependent code
    - Code moved into new `TorrentManager`. This is BROKEN! The code is not ready to use so don't use it, lots more work to do. Biggest problem is that the torrent management in `Manager` is very tightly coupled to sending encrypting etc. All of that needs to be prised apart
  - Ensured move of all torrent funcs and structs
    - I also ensured that the order of functions matches the original code, to make comparison easier during review.

## Schedule

Note details of this below call were shared in the Snow Blowers' chat and in the general `status-go` Discord / Status channel.

- 12:00 – 13:00 : ❄🔥 Snow Blowers (status-go flaky tests) Sync Call
  - https://meet.google.com/cvh-crfd-uji
  - Igor presented his intent to make diff test coverage at 50% per PR REQUIRED.
  - Attendees:
    - Igor (facilitator)
    - Patryk
    - Myself
  - Summary
    - Discussed the latest nightly flaky results:
      - https://ci.status.im/job/status-go/job/tests-nightly/235/Reports/
      - https://ci.status.im/job/status-go/job/tests-nightly/235/
      - We agreed to target all of the `TestMessengerCommunitiesTokenPermissionsSuite` flakes
        - [See the log file for more details](./attachments/2024-05/2024-05-29_02:48:00_nightly.log)
    - We agreed to activate REQUIRED and considered a number of points that would need to be addressed.
      - Generated files should be ignored
      - Test coverage of 50% should be reasonable and not burdensome while effective at enforcing necessary code quality.
      - We have a mechanism to override the requirement in EXTREME cases
    - QA should be given visibility of this work
    - [Make codeclimate diff coverage test reports Required #5254](https://github.com/status-im/status-go/issues/5254)
      - [fix missing value of keyuid for old mobile user #5203](https://github.com/status-im/status-go/pull/5203) A case where the tests are present and seem to be robust but Codeclimate reports that the test coverage is only 36% (this case is notable because it seems unfair, may need investigation)
      - [refactor_: ExtractTokenCriteria -> extractContractAddressesByChain #5226](https://github.com/status-im/status-go/pull/5226) A case where code changes are made but have no Codeclimate test coverage report
      - [Direct settings updates (API) #5237](https://github.com/status-im/status-go/pull/5237) A case with 0% coverage
      - [Router Filter Modularisation and Test suite #5177](https://github.com/status-im/status-go/pull/5177) A case with 100% coverage 😉

---

# 2024-05-28

## Issues

- [Refactor `fromIncluded` and `fromExcluded` Maps to Slices in Route Validation Logic](https://github.com/status-im/status-go/issues/5241) - `created`
- [Refactor and Simplify `hasSufficientCapacityV2` Logic to Ensure Correct Path Amount Validation](https://github.com/status-im/status-go/issues/5244) - `created`
- [Check Validity of `fromLockedAmount` in `SuggestedRoutesV2`](https://github.com/status-im/status-go/issues/5227) - `updated`

## Pulls

- [Router Filter Modularisation #5177](https://github.com/status-im/status-go/pull/5177) -`merged`
  - Addressed feedback, mostly consisting of follow-on issues for various logic improvements.
    - https://github.com/status-im/status-go/issues/5227
    - https://github.com/status-im/status-go/issues/5241
    - https://github.com/status-im/status-go/issues/5244
  - Note: I would prefer to merge the existing logic first, even if this logic is incorrect, because this is how original logic works.
    - I just want to make sure that we have a clear history of changes from the original into whatever version we end with.
  - I need to rewrite the commit names to follow the new convention.
    - https://github.com/status-im/status-go/pull/5177#issuecomment-2134929345
  - Merged!

## Schedule

- 12:00 – 12:30 : mobile-core planning
  - https://meet.google.com/azr-ppob-obc
  - Stand-ups
    - https://www.notion.so/Week-Beginning-20aeca44ce2045058d472784e02aed45
  - Discussed priorities for this week:
    - 2.29 Release
    - 2.30 priorities:
      - Performance
      - Installation file size, get under 100mb
      - Stablise and polish communities and wallet functionality
      - Help where needed.

---

# 2024-05-27

AFK - Spring public holiday 🚂🚃🚃🚃

---

# 2024-05-24

## Issues

- [Check Validity of fromLockedAmount in SuggestedRoutesV2 #5227](https://github.com/status-im/status-go/issues/5227) `created`
  - Follow-on work to ensure the robustness of the wider Wallet `Router` functionality. 

## Pulls

- [Router Filter Modularisation #5177](https://github.com/status-im/status-go/pull/5177)
  - Finished! 🎉 Ready for full review
  - Big commit implementing additional test cases to make the test suite as comprehensive as I can.
    - That doesn't mean it is fully comprehensive
  - Addressed feedback
  - Resolved insufficient rest test case issue
  - Implemented `testing.T` `t.Logf` test level logging
  - Implemented pure `zap` debug logging for all relevant filter functions.

## Schedule

- [x] 10:30 – 11:30 : Desktop Send modal review
  - https://meet.google.com/pow-qyes-pcv
  - "A meeting to dispel doubts about the design and current state of SendModal."
  - **Attendees:**
    - Michal C (Facilitator)
    - Alisher
    - Ben
    - Brian
    - John (Lea)
    - Sale
    - Myself
  - **Summary:**
    - Discussion clarifying the rules and requirements of the desktop designs for sending tokens.
      - No agreement yet on where to capture the rules
    - Initially some confusion on the distinction between multichain send "without bridge" and "bridging"
      - Confusion as sending over multiple chains requires bridging
      - Clarity eventually given; references to "bridging" is a term for a specific UI action and not an implementation behaviour.
        - ie in the UI context bridging refers to the user pressing a "bridging" button to perform an explicit bridging action. Not the required routing performed by `status-go` Wallet `Router`.
        - ⚠️ Attention. Devs and QAs particularly should be mindful of this.
    - Emphasised that the app should mint `ERC1155` rather than `ERC721` tokens
      - Confirmation that this is the case was suggested.
    - Highlighted that there is an inconsistency with the current desktop app functionality verse the intention of the design.
      - Locking behaviour for values on specific chains does not work as the product design intended
        - Full requirements not yet provided, expected to follow.
        - Additional clarification and validation likely to be needed
      - Discussed that `ERC1155` tokens of the same kind should be grouped together.
        - Unclear if designs and requirements for this feature exist in the form needed by Desktop UI.
        - Dario has mentioned in chat: > "*This is mostly a backend effort with relatively high complexity (define a special kind of token for these "community ERC721 we want to treat as ERC1155" ones, do the grouping backend-side when fetching them, do the un-grouping backend-side when sending, ensure no side effects in activity). I see low chances of swap-team being able to work on + finish that for 2.30, unless it's prioritized over Swap.*"
          - https://discord.com/channels/1210237582470807632/1217175558522278008/1242844033873739837
        - So this intent is / has been expressed somewhere.
    - Brief update from Brian and me.
      - The call expressed a blithesome response to the update that the Wallet Router filter modularisation and test suite was complete and ready for review.

---

# 2024-05-23

## Pulls

- [Router Filter Modularisation #5177](https://github.com/status-im/status-go/pull/5177)
  - Resolved a boat load of test scenario issues
    - I'm super seriously very nearly finished now.
    - However, I do need reviewers to pay super close attention to the test cases and check that the intent of the functions is matched by the cases.
  - The full list of changes is available in the PR commits.

## Schedule
- [x] 11:00 - 12:00 : Andrea and Samuel 1:1
  - https://meet.google.com/rhy-trwc-ybb
  - <details>
    <summary style="color:red;">Confidential</summary>
      <img src="./attachments/2024-05/200w.gif" title="Dennis Nedry - Jurassic Park" alt="&quot;Ah ah aaaah. You didn't say the magic word.&quot;"/>
    </details>
- [x] 13:00 – 13:30 : mobile-core planning
  - https://meet.google.com/azr-ppob-obc
  - Stand-ups
  - Discussion of next week's aims and prioritising for release
  - John Ngei is the 🤘 `mobilecore` 🤘 release manager for `v2.29` 🎉
  - Agreed that we should have a weekly standup document:
    - We will use the document that Mariia created:
      - https://www.notion.so/Stand-ups-74686c0ec06e4f64a130136db9771f84
    - Anyone can create a new weekly document using the magic template button.
      - Template : https://www.notion.so/Week-Beginning-63e5359505d74d3da8b3f93608544317

---

# 2024-05-22

## Pulls

- [Router Filter Modularisation #5177](https://github.com/status-im/status-go/pull/5177)
  - Resolved a boat load of test scenario issues
    - I'm very nearly finished now. However, I do need reviewers to pay super close attention to the test cases and check that the intent of the functions is matched by the cases.
  - The full list of changes is available in the PR commit.

## Scoping

- ₿♢Ξ Revenue, Revenue, Revenue Ξ♢₿
  - https://www.notion.so/Monetisation-6421841be37c4343a4342e6715924c91
  - Did refinement to some of my market analysis and monetisation recommendations for Status

## Ad Hoc

- Confirming details about my EthCC presentation

---

# 2024-05-21

## Pulls

- [Router Filter Modularisation #5177](https://github.com/status-im/status-go/pull/5177)
  - Fixes to capacity logic
    - handling for nil pointers added.
  - updated the tests, quite a considerable amount change on this.

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
