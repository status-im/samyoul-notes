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
  