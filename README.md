# 2024-06-07

## Issues

- [Remove torrent dependencies at build time for mobile `#5146`](https://github.com/status-im/status-go/issues/5146) - `closed`
- [Refactor `Messenger` to Delegate Archive Management to `communities.Manager` `#5313`](https://github.com/status-im/status-go/issues/5313) - `created`
- [Check Validity of `fromLockedAmount` in `SuggestedRoutesV2` `#5227`](https://github.com/status-im/status-go/issues/5227) - `scoping`
- [Refactor, Split and Clean Up Exported Functions in `ArchiveManager` `#5316`](https://github.com/status-im/status-go/issues/5316) - `created`

## Pulls

- [Test PR to check that removing torrent code from status-go will give us a smaller APK #20393](https://github.com/status-im/status-mobile/pull/20393) - `created`, `discussed`, `closed`
  - A PR to prove that my work on [Removing torrent from mobile build #5257](https://github.com/status-im/status-go/pull/5257) actually does what we expect it to do.
  - IT DOES!!! https://github.com/status-im/status-mobile/pull/20393
    - 🎉 Android APK **93.7mb** 🎉
    - 🎉 iOS IPA **55.7mb** 🎉
- [Removing torrent from mobile build #5257](https://github.com/status-im/status-go/pull/5257) - `commits`, `merged`
  - Refactored NewArchiveManager to use config pattern
  - MERGED!

---

# 2024-06-06

## Pulls

- [Removing torrent from mobile build #5257](https://github.com/status-im/status-go/pull/5257) - `commits`
  - Readied this PR for review, and requested reviews
  - Gave proper attribution and full context
  - Replaced entirely `LogStdout` with default `logger`
  - Replaced `fmt.Sprintf` usage from `zap.logger` calls
  - Explicitly cast `ManagerSuite.torrentManager`
    - Instead of using the `TorrentContract` interface I've set the field to expicitly declare as `*TorrentManager`. This removes all the random type casting used in the tests. I also removed all the usages of `buildTorrentConfig()` as we build the test suite with the torrent config now.
  - Renamed **ManagerMobile to **ManagerNop
  - Renamed Torrent to Archive
    - I've renamed `TorrentManager` to `ArchiveManager`, `ArchiveManager` to `ArchiveFileManager`, `TorrentContract` to `ArchiveService`, `ArchiveContract` to `ArchiveFileService`. I've also renamed all init functions and struct fields to the appropriate archive-centric naming.
  - Renamed archive files to archive_file
  - Renamed torrent files to archive
  - Fixed failing tests
  - Gave a defence of my approach and acknowledgment of its failings
    - https://github.com/status-im/status-go/pull/5257#issuecomment-2152326864

## Reviews

- [Add timeout to call rpc endpoint #5285](https://github.com/status-im/status-go/pull/5285) - `approved`
- [decreased limits for concurrent requests and rps](https://github.com/status-im/status-go/pull/5253) - `approved`, `feedback`
- [status-go integration tests](https://github.com/status-im/status-go/pull/5302) - `approved`

## Schedule
- [x] 11:00 - 12:00 : Andrea and Samuel 1:1
  - https://meet.google.com/rhy-trwc-ybb
  - <details>
    <summary style="color:red;">*Knock knock*</summary>
      <img src="./attachments/2024-05/200w.gif" title="Dennis Nedry - Jurassic Park" alt="&quot;Ah ah aaaah. You didn't say the magic word.&quot;"/>
    </details>

---

# 2024-06-05

## Pulls

- [Removing torrent from mobile build #5257](https://github.com/status-im/status-go/pull/5257) - `commits`
  - Added pre 1.17 build commands
  - [Had a number of frustrating build fails](https://github.com/status-im/status-go/issues/5146#issuecomment-2149259488)
  - Copy and pasted Andrea's work
    - https://github.com/status-im/status-go/pull/5295
    - FFS Andrea spent 10 minutes looking at this and fixed what I'd spent 2 hours rage quitting about.
  - Fixed torrent tests
    - This was annoying
  - Resolved rebase conflicts :)

## Ad Hoc

- Got gold star approved for EthCC

---

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
