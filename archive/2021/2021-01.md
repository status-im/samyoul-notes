# 2021-01-29

## AdHoc
- [x] Preparation for anonymous metrics and scoping go implementation

## Schedule
- [x] 09:00 GMT - Pairing `status-react` (Andrea & Samuel)
  - https://meet.google.com/zrn-ovoh-tdk
- [x] 13:00 GMT - Sync anonymous metrics
  - https://meet.google.com/boj-gnpj-qoc
  - https://notes.status.im/anonymous-metrics
- [x] 14:00 GMT - Core Retro
  - https://meet.google.com/hqw-nsfk-xoq
  - https://notes.status.im/core-retro-2021-01-29?both

---

# 2021-01-28

## PRs
- [x] [#2131 - Adding Crypto On Ramp API Endpoint to Wallet Service](https://github.com/status-im/status-go/pull/2131) - `commits`, `submitted`

## Reviews
- [x] [#2136 - Make sure tokens balance can be fetched](https://github.com/status-im/status-go/pull/2136) - `approved`, `feedback`

## Schedule
- [x] 09:00 GMT - Pairing `status-react` (Andrea & Samuel)
  - https://meet.google.com/xdm-ecdg-fyq

---

# 2021-01-27

## AdHoc
- [x] Anonymous Metrics info from Shivek
  - [Background overview](https://discuss.status.im/t/waku-based-analytics/2082)
  - [System Overview](https://miro.com/app/board/o9J_lY7ZiUI=/)
  - [Data we want to collect](https://docs.google.com/spreadsheets/d/1dGX57QbnRrSGHdz49vG0JevEmbPQbgRdQ0ROe1ZULRQ/edit#gid=0)
  - [Hester's notes](https://notes.status.im/anonymous-metrics)
  - [Frontend implementation draft](https://notes.status.im/s/Gae7Ocb28)

## PRs
- [x] [#2131 - Adding Crypto On Ramp API Endpoint to Wallet Service](https://github.com/status-im/status-go/pull/2131) - `commits`

## Schedule
- [x] 13:30 GMT - Fiat <> Crypto Onramp Sync
  - https://meet.google.com/pzc-efwz-nyj
  - [Notes](https://notes.status.im/eTU27ajnSteRWYCoCceYCQ?both)
- [x] 14:00 GMT - Andrea & Samuel 1-to-1
  - https://meet.google.com/rhy-trwc-ybb
  - Intro to building out new features in `status-react`

---

# 2021-01-26

## AdHoc
- [x] Setting up `status-react`
- [x] Learning the ways of Clojure 

## PRs
- [x] [Crypto On Ramps](https://github.com/status-im/crypto-on-ramps) - `commits`
  - Changed json from `kebab-case` to `camelCase`
- [x] [#2131 - Adding Crypto On Ramp API Endpoint to Wallet Service](https://github.com/status-im/status-go/pull/2131) - `created`, `commits`

## Reviews
- [x] [#2128 - Confirmation setting fixed for mailservers](https://github.com/status-im/status-go/pull/2128) - `approved`, `feedback`
- [x] [#2129 - Confirmations processing fixed on client side](https://github.com/status-im/status-go/pull/2129) - `approved`
- [x] [#2130 - Add back notification code](https://github.com/status-im/status-go/pull/2130) - `approved`

## Schedule
- [x] status react set up (Andrea & Samuel)
  - https://meet.google.com/cjk-pdjc-xew

---

# 2021-01-25

## PRs
- [x] [Crypto On Ramps](https://github.com/status-im/crypto-on-ramps) - `commits`
  - Addresses [status-react#11644 - Add list of website where you can buy crypto](https://github.com/status-im/status-react/issues/11644)

## Reviews
- [x] [#2126 - Pass topics to mailserver cursor](https://github.com/status-im/status-go/pull/2126) - `approved`
- [x] [#2125 - fix "Pending" status for messages in private group chat](https://github.com/status-im/status-go/pull/2125) - `approved`, `discussed`
- [x] [#2122 - Add export/import methods](https://github.com/status-im/status-go/pull/2122) - `approved`
- [x] [#2121 - Clean topics that we don't listen to](https://github.com/status-im/status-go/pull/2121) - `approved`
- [x] [#2127 - Fix segmentation failure on mailserver](https://github.com/status-im/status-go/pull/2127) - `approved`
- [x] [#2120 - Use personal topic for push notification registration](https://github.com/status-im/status-go/pull/2120) - `approved`
- [x] [#2119 - Skip wrapping emojis in private group chats](https://github.com/status-im/status-go/pull/2119) - `feedback`

## Schedule
- [x] 14:30 BST - Q1 Core Priorities (Andrea and Sam)
  - https://meet.google.com/mnb-bnzo-ydj?
  - **Crypto onboarding**
    - https://github.com/status-im/status-react/issues/11644
    - https://www.figma.com/file/XUehMnhyD1FGcWzvGz6SXqvh/Wallet?node-id=10672%3A43638

---

# All the intervening days

- Sleep training my daughter

---

# 2021-01-12

## Pulls
- [x] [#2100 - Expand Local Notifications](https://github.com/status-im/status-go/pull/2100) - `commits`, `merged`
- [x] [Animated Image Manipulation](https://github.com/status-im/animated-image-manipulation) - `commits`
  - **Issues :**
    - Re-encoding resized gifs give an error `fail to encode gif lzw: input byte too large for the litWidth`
    - Resizer changes the image type from `image.Paletted` to `image.RGBA64`, (which may be related to the above issue)

---

# 2021-01-11

## Pulls
- [x] [#2100 - Expand Local Notifications](https://github.com/status-im/status-go/pull/2100) - `commits`, `submitted`
- [x] Continued work on basic gif manipulation repo, for researching gif profile images and gif type chat messages

## Reviews
- [x] [#2112 - Cache waku messages](https://github.com/status-im/status-go/pull/2112) - `approved`
- [x] [#2111 - Send abridged version of history with messages in group chats](https://github.com/status-im/status-go/pull/2111) - `approved`

---

# 2021-01-08

## Pulls
- [x] [#2100 - Expand Local Notifications](https://github.com/status-im/status-go/pull/2100) - `commits`
- [x] Began work on basic gif manipulation repo, for researching gif profile images and gif type chat messages

---

# 2021-01-07

# Reviews
- [x] [#2110 - Fix communities migration](https://github.com/status-im/status-go/pull/2110) - `approved`

# Issues
- [x] [status#56 - Profile picture privacy is worse than Signal](https://github.com/status-im/status/issues/56) - `discussed`, `research`
  - [Proposed solution for private only profile images](https://github.com/status-im/status/issues/56#issuecomment-756222460)

# Pulls
- [x] [Todos To Docs](https://github.com/status-im/todo-to-docs) - `commits`
- [x] [#2100 - Expand Local Notifications](https://github.com/status-im/status-go/pull/2100) - `commits`

# Ad Hoc
- [x] Completed list of known technical debt in `status-go`
  - [Document Here](https://notes.status.im/cLQWkUEbTQmsIvUGAyj95A?view)
  - Now can be done automatically with the help of the [Todos To Docs](https://github.com/status-im/todo-to-docs)

---

# 2021-01-06

// TODO

---

# 2021-01-05

// TODO

---

# 2021-01-04

// TODO
