# 2024-07-03

## Issues

- [Check Validity of fromLockedAmount in SuggestedRoutesV2 `#5227`](https://github.com/status-im/status-go/issues/5227) - `closed`

## PRs

- [Added test suite for validateInputData() `#5323`](https://github.com/status-im/status-go/pull/5323) - `commits` `merged`
  - blended my validation with latest develop
  - MERGED
- [chore: skip keycard code paths that crash on intel MacOS `#15194`](https://github.com/status-im/status-desktop/pull/15194)
  - Tested Sid's latest work on v2.29 Intel Mac build for Desktop
  - [It worked!](https://github.com/status-im/status-desktop/pull/15194#issuecomment-2205598645)

## Schedule

- 12:00 – 13:00 : ❄🔥 Snow Blowers (status-go flaky tests) Sync Call
  - https://meet.google.com/cvh-crfd-uji
  - Discussion of a number of issues with testing and ideological decisions negatively impacting on productivity.
  - Discussion of API wrapper test exemptions
    - [Remove API/wrappers from test coverage `#5466`](https://github.com/status-im/status-go/issues/5466)

## Ad Hoc

- Review of latest changes to the Status Internal Comms Policy

---

# 2024-07-02

# PRs

- [Added test suite for validateInputData() `#5323`](https://github.com/status-im/status-go/pull/5323) - `commits`
  - integrated the new errors into test logic
  - Resolved invalid recursive type aliasing 
    - Additionally, I've resolved a panic when AmountIn and/or AmountOut is nil

## Schedules

- 12:00 – 13:00 : mobile-core planning
  - https://meet.google.com/azr-ppob-obc
  - [Call notes](https://www.notion.so/Planning-Call-Notes-4892895dbeff4fdfbb92e7a0b06b3a0eO)
  - [Stand-up document](https://www.notion.so/Week-Beginning-cbcad6aa2c2b4f7886e70ed016088d3a)

---

# 2024-07-01

## Schedule

- 13:00 – 13:45 : bi-weekly status-go call
  - https://meet.google.com/gbq-tyqe-vju
  - [Call notes](https://www.notion.so/status-go-Meeting-Notes-e3d2e2cb2cce46968634ffd673d932ea)

## Ad Hoc

- Organising my week
- Working on EthCC goals and plans
