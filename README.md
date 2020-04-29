# 2020-04-28

## Pulls

- [x] [#1957 - Refactor/waku tidy](https://github.com/status-im/status-go/pull/1957) - `commits`

## Reviews

- [x] [Status Principles Webinar](https://www.youtube.com/watch?v=rBBpZxZQPy8) - `watched`
- [x] [#69 - Send images](https://github.com/status-im/specs/issues/69) - `read`

## Issues

- [x] [#1937 - Compress public key for chat](https://github.com/status-im/status-go/issues/1937) - `prototype`
  - [Initial draft](https://github.com/status-im/status-go/issues/1937#issuecomment-620603814)
  
---

# 2020-04-27

## Schedule

- [x] 14:00 BST - Vac PM
  - https://meet.google.com/xzi-onqc-gdx
  - https://notes.status.im/Kt7iRwy1TeCfNzcJGnt3Ag?both 
- [x] 15:00 BST - On the nature of the refactor: waku/whisper
  - https://meet.google.com/dcu-dnot-jak
  - [Code duplication between waku and whisper packages](https://codeclimate.com/github/status-im/status-go/pull/1950)
  - [#1945 - Database package location refactor](https://github.com/status-im/status-go/issues/1945)
    - Decided that the `whisper` package will be left along from now. Any changes needed for Waku will only affect the `waku` package.
    - Need to add a DEPRECATION message in the `whisper` package 

## Reviews

- [x] [#1947 - prepare for waku/1](https://github.com/status-im/status-go/pull/1947) - `feedback`, `approved`
- [x] [Code duplication between waku and whisper packages](https://codeclimate.com/github/status-im/status-go/pull/1950)

---

# 2020-04-24

## Schedule

- [x] 14:00 BST - Core Retro
  - https://meet.google.com/hqw-nsfk-xoq
  - [Summary](https://notes.status.im/tl3SGbKHTGC8kL_we_aeng?both)

## Pulls

- [x] [#1949 - Waku package README.md documentation](https://github.com/status-im/status-go/pull/1949) - `commits`
- [x] [#1951 - Refactor moved message functionality from filter to message](https://github.com/status-im/status-go/pull/1951) - `merged`

## Reviews

- [ ] [Code duplication between waku and whisper packages](https://codeclimate.com/github/status-im/status-go/pull/1950)

---

# 2020-04-23

## Ad Hoc

- [x] [File format lint fails are the worst](https://ci.status.im/blue/organizations/jenkins/status-go%2Fprs%2Ftests/detail/PR-1950/4/pipeline)
  - Note : The solution for this fail was insane, see https://github.com/status-im/status-go/pull/1950/commits/04413b9e31e79230f55531cd34cd1481225da993

## Pulls

- [x] [#1949 - Waku package README.md documentation](https://github.com/status-im/status-go/pull/1949) - `commits`
- [x] [#1950 - Refactor/waku.doc](https://github.com/status-im/status-go/pull/1950) - `merged`

## Reviews

- `status-im/specs`
  - [x] [#54 - Early draft: Waku mode spec](https://github.com/status-im/specs/pull/54) - `read`
  - [x] [#130 - Transaction history management](https://github.com/status-im/specs/issues/103) - `read`, `discussed`

---

# 2020-04-22

## Pulls

- [x] [#1949 - Waku package README.md documentation](https://github.com/status-im/status-go/pull/1949) - `draft`

## Reviews

- [x] [#1948 - Set chat active on being re-invited](https://github.com/status-im/status-go/pull/1948/files) - `feedback`

---

# 2020-04-21

## Schedule

- [x] 11:30 BST - 1:1 with Hester
  - https://meet.google.com/fef-vjgd-zxb
  - Points
    - [Open UI position](https://status.im/our_team/open_positions.html?gh_jid=2123002)
    - [Example of Discuss discussion](https://discuss.status.im/t/a-case-for-breaking-backward-compatibility-in-v1/1279)
- [x] 14:00 BST - 1:1 with Andre
  - https://meet.google.com/pga-udvu-wpk
  - Points (spec all the things)
    - https://github.com/status-im/status-react/issues/10384
    - https://github.com/status-im/specs/issues/69
    - https://github.com/status-im/status-go/issues/1834

## Reviews

- `vacp2p/specs`
  - [x] [#113 - Upgrade waku specs to waku/1](https://github.com/vacp2p/specs/pull/113#pullrequestreview-397262991) - `approved`
  - [x] [#41 - Waku 0: Document forward compatibility](https://github.com/vacp2p/specs/issues/41) - `read`
- `status-im/specs`
  - [x] [#61 - Replace forked Whisper usage with Waku/0 in specs](https://github.com/status-im/specs/issues/61) - `read`
  - [x] [#73 - Remove certain specs, refer to Waku](https://github.com/status-im/specs/issues/73) - `read`
- [x] https://discuss.status.im/t/roadmap-planning/1399/38 - `read`

---

# 2020-04-20

## Schedule

- [x] 10:30 BST - 1:1 with Andrea, introduction call and `status-go` discussion
  - https://meet.google.com/gmx-csjy-ett
- [x] 13:00 BST - Status Devs Meeting
  - Location - https://jitsi.status.im/core-dev
  - Notes
    - https://notes.status.im/core-dev-call-28
    - https://github.com/status-im/pm/issues/50
    - https://discuss.status.im/t/stable-status-specs-and-sips/1599
  - Video - https://www.youtube.com/watch?v=pvswK5dR1mc

## Pulls

- [#1946 - Update team collaboration forum](https://github.com/status-im/status-go/pull/1946) - `merged`
- [#1931 - Waku handshake RLP key variable initialisation](https://github.com/status-im/status-go/pull/1931) - `merged`

## Issues

- [#1944 - Add waku package README.md docs](https://github.com/status-im/status-go/issues/1944) - `created`
- [#1945 - Database package location refactor](https://github.com/status-im/status-go/issues/1945) - `created`

## Reviews

- [x] [#1943 - Don't pass contacts without custom fields to the client](https://github.com/status-im/status-go/pull/1943#pullrequestreview-396473829) - `approved`

---

# 2020-04-17

## Schedule

- [x] 14:00 BST - 1:1 with Andre
  - https://meet.google.com/cgu-rszi-veu

---

# 2020-04-16

## Schedule

- [x] 13:45 BST - 1:1 with Petty
  - https://meet.google.com/onc-hnou-sbs
  - Points
    - [Petty's nim status project](https://github.com/status-im/nim-stratus-console)
    - [Miro - Diagram designing application](https://miro.com/app/board/o9J_kxw_3_4=/)
    - https://github.com/status-im/status-security/blob/master/drafts/authentication-and-authorization.md
    - [HackMD Features](https://hackmd.io/features)

---

# 2020-04-15

## Reviewed

- [x] [V1.5+ Objective Ranking](https://docs.google.com/spreadsheets/d/14N74hACVxG6X7j5WqUEqlYKnEbNyH9Po7rpcgiBKntw/edit#gid=0) - `voted`

---

# 2020-04-14

## Schedule

- [x] 12:15 BST - 1:1 with Oskarth (text chat)
  - https://meet.google.com/onc-hnou-sbs
  - Points
    - https://github.com/status-im/specs/issues/73
    - https://github.com/status-im/specs/issues/61

---

# ...

TODO Add missing days

---

# 2020-04-06

## Schedule

- [x] 12:15 BST - 1:1 with Jakub discussing `status-go`
  - https://meet.google.com/aqr-tdun-dfg
  - Points
    - https://github.com/status-im/jakubgs-notes
    - https://fleets.status.im/
    - https://github.com/status-im/infra-role-bootstrap/tree/master/files/keys
    - https://github.com/status-im/infra-role-bootstrap/blob/447080f5af6fcf8aa1a99f1cceaac3835e608fa2/defaults/main.yml#L102
    - https://github.com/status-im/infra-eth-cluster
    - https://consul.status.im/
    - https://consul.status.im/ui/do-ams3/nodes?filter=mail
    - https://consul.status.im/ui/do-ams3/nodes/mail-01.do-ams3.eth.prod
    - https://grafana.status.im/
    - https://grafana.status.im/d/wV77E-4mz/statusd-metrics?orgId=1
    - https://grafana.status.im/d/gxQG_R1Zk/status-peers?orgId=1&refresh=5m
    - https://github.com/status-im/status-go/blob/develop/MAILSERVER.md
    - https://github.com/status-im/status-go/tree/develop/_assets/compose/mailserver
- [x] 13:30 BST - 1:1 with Volodymyr discussing `status-go` and desktop implementation of `status-react`
  - Points
    - https://github.com/status-im/status-react/blob/e7a213fe95f24c18cffd6d13d08db6aa81b3a6b1/modules/react-native-status/ios/RCTStatus/RCTStatus.m#L171
    - https://github.com/status-im/status-react/blob/develop/modules/react-native-status/desktop/CMakeLists.txt

---

# 2020-04-01

## Schedule

- [x] 12:00 BST - 1:1 with Ceri 
  - https://meet.google.com/mtp-qqtw-bpg
  - On boarding day :D

---
