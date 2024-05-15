# 2024-05-15

## Scoping

- Wallet Router
  - Complete analysis for the `findBest` function:
    - [See here for details `findBest`](./analysis/wallet/Router/newSuggestedRoutes.md#findbest)

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
