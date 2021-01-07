# 2020-12-23

## Pulls

- [x] [#2100 - Expand Local Notifications](https://github.com/status-im/status-go/pull/2100) - `commits`


## Schedule

- [x] 14:00 GMT - Samuel & Andrea 1-to-1
  - https://meet.google.com/rhy-trwc-ybb

---

# 2020-12-22

## AdHoc

- Quirk that prevented tests from running
  - `xcrun: error: invalid active developer path (/Library/Developer/CommandLineTools), missing xcrun at: /Library/Developer/CommandLineTools/usr/bin/xcrun`
  - [Solution](https://ma.ttias.be/mac-os-xcrun-error-invalid-active-developer-path-missing-xcrun/)

## Pulls

- [x] [#2100 - Expand Local Notifications](https://github.com/status-im/status-go/pull/2100) - `debugging`
  - Resolved the weird bug where adding `MarshalJSON()` to the `Notification` struct would result in a fail
    - Initially I thought that the problem was because the `signal.send()` was handling the `json.Marshalling` in a strange way
    - The real bug was in `signal.send()` and caused by the new `Body.Type` field not being set.
    - This wasn't immediately obvious because `signal.send()` uses a logger to record errors and the logger wasn't printing to the terminal, therefore a failed marshal resulted in an error not about the thing that had actually gone wrong, adding lots of `spew.Dumps()` helped
    - A minor note the `signal` package has some weird structure, for example setting and getting a package var as a kind of global, therefore maintaining an implicit "package state".    

---

# 2020-12-21

## Pulls

- [x] [#2100 - Expand Local Notifications](https://github.com/status-im/status-go/pull/2100) - `debugging`
  - Working on a weird bug where adding `MarshalJSON()` to the `Notification` struct would result in a fail.

## Review

- [x] [#2046 - Communities](https://github.com/status-im/status-go/pull/2046) - `approved`
- [x] [status-im/specs#150 - Feature/org channels](https://github.com/status-im/specs/pull/151) - `feedback`
  - Discussing primarily the functionality related to the `grant` concept
  - [details](https://github.com/status-im/specs/pull/151#discussion_r546714736)
  
---

# 2020-12-18

## Ad Hoc

- [x] Resolve weird issue on my machine where make vendor has started throwing `modvendor: command not found`
  - Caused by an Apple update https://stackoverflow.com/a/59138750/1700106
  - Solution was to rename `./bash_profile` to `./zprofile`
- [x] Compile list of known technical debt in `status-go`
  - [Document Here](https://notes.status.im/cLQWkUEbTQmsIvUGAyj95A)
  - Based on the `//TODO` items in the code
  - Gone through all directories up to `protocol`, need to finish `protocol`, all `protocol` sub directories and all remaining directories

## Review

- [x] [#2046 - Communities](https://github.com/status-im/status-go/pull/2046) - `feedback`
  - Initial review require a deeper review of the core functionality

## Schedule

- [x] 14:00 GMT - Core Retro
  - https://meet.google.com/hqw-nsfk-xoq
  - A review / summary of the year.

---

# 2020-12-17

## Pulls

- [x] [#2100 - Expand Local Notifications](https://github.com/status-im/status-go/pull/2100) - `commits`, `research`
  - https://notes.status.im/notifications-android
  - Currently have a bug that I think is caused by a the jsonMarshal func not working as intented, need to resolve this next before proceeding

---

# 2020-12-16

## Pulls

- [x] [#2051 - Profile Images](https://github.com/status-im/status-go/pull/2051) - `commits`
  - Related to debugging
  - Issue finally resolved by Cammellos
    - [tracking of the messages would deadlock](https://github.com/status-im/status-go/blob/50b17308bde0008daf4c8365782575d4f22b1515/protocol/messenger.go#L306)
    - [Making `EnvelopesMonitor.Add()` deadlock](https://github.com/status-im/status-go/blob/50b17308bde0008daf4c8365782575d4f22b1515/protocol/transport/waku/envelopes.go#L108)
    - [Because `EnvelopesMonitor.handleEventEnvelopeSent()` couldn't unlock](https://github.com/status-im/status-go/blob/50b17308bde0008daf4c8365782575d4f22b1515/protocol/transport/waku/envelopes.go#L165)
    - [Fix was to remove mutex lock from `Messenger.processSentMessages()`](https://github.com/status-im/status-go/blob/50b17308bde0008daf4c8365782575d4f22b1515/protocol/messenger.go#L306)

## Issues

- [x] [#11347 - Profile edit picture UI](https://github.com/status-im/status-react/pull/11347) 
  - [App Freezing when user sends message after changing profile image](https://github.com/status-im/status-react/pull/11347#issuecomment-746252933)
    - Debugging

## Schedule

- [x] 14:00 GMT - Samuel & Andrea 1-to-1
  - https://meet.google.com/rhy-trwc-ybb
  - **Notes :**
    - *Profile Images*
      - [App Freezing when user sends message after changing profile image](https://github.com/status-im/status-react/pull/11347#issuecomment-746252933)
    - *`status-go` notifications improvement*
      - Adding `Notifications` field to `MessageResponse` struct and parsing in the `PublisherSignalHandler`
  