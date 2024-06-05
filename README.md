# 2024-06-04

## Pulls

- [Removing torrent from mobile build #5257](https://github.com/status-im/status-go/pull/5257) - `commits`
  - Implemented build OS conditional build instructions
  - Added more complex build conditions to exclude OSes
  - 😱 Debugging why OS conditional builds are not working 😱

## Reviews

- [fix missing value of keyuid for old mobile user #5203](https://github.com/status-im/status-go/pull/5203) `approved`, `feedback`
  - Related to this PR there seems to be some discrepancy between what Code Climate judges as having test coverage and what IntelliJ IDEA reports.
  - From the case of this PR it seems that Code Climate is less accurate and reports that some statements are not covered when they are.

## Schedule

- [x] 12:00 – 13:00 : mobile-core planning
  - https://meet.google.com/azr-ppob-obc
  - 🎉 Announcement - Ícaro's 100% guaranteed promotion to `🤘 mobilecore 🤘` lead 🎉
  - Discussion and review of ongoing release V2.29 cycle
    - Discussed Waku issues that are causing major blockers
  - Stand-ups
    - https://www.notion.so/Week-Beginning-af2a460c8d0d43f290fe9730241b2323
  - Planning session

---

# 2024-06-03

## Pulls

- [Removing torrent from mobile build #5257](https://github.com/status-im/status-go/pull/5257) - `commits`, `scoping`
  - Removed TorrentManager from `handleImportedMessages`
  - Removed TorrentManager from `checkIfIMemberOfCommunity`
  - Moved archive related funcs into `ArchiveManager`
    - To be honest once I started this work I quickly realised how pointless it is as archiving functionality and torrent seeding functionality are really entwined. So I'm keeping the code I've done, but it is a bit pointless without spending a lot of time untangling torrenting and archiving. I'm just going to make an interface for all the functions that are used publicly from `TorrentManager`. I think that this will be the fast way to move on from this issue. I don't like this work anymore, there is a lot of work to do elsewhere and torrent is a rabbit hole filled with canned worms.
  - Created `TorrentContract` and `ArchiveContract` interfaces
  - Created nil structs for Mobile use (or close to nil as possible)
  - Mapped all 2nd level `TorrentManager` usage
    - https://github.com/status-im/status-go/pull/5257#issuecomment-2143956872
    - Follow on from [`TorrentManager` API Usage](https://github.com/status-im/status-go/pull/5257#issuecomment-2142881649)

## Schedule

- [x] 11:30 – 12:00 : Melanie <> Sam
  - https://meet.google.com/xmr-vjog-zj
  - Discussion of EthCC
  - And to be honest just a very delightful catchup ❤️
- [x] 13:00 – 13:45 : bi-weekly status-go call
  - https://meet.google.com/gbq-tyqe-vju
  - Stand-ups
  - Belal was introduced to the group (Hey Belal)
  - Discussion of how to ensure compatibility between V1 and V2 key derivation and generation
  - Action points / Key Decisions
    - Activating Code Climate diff test coverage as REQUIRED
    - Agreement to activate audio transcription to ensure we have low effort, yet (reasonably) accurate notes for our calls
