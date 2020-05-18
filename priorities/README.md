# Priorities

## Personal

1) [status-im/specs#61 - Replace forked Whisper usage with Waku/0 in specs](https://github.com/status-im/specs/issues/61)
   - Related :
     - [status-react#5590 - Whisper spamming protection](https://github.com/status-im/status-react/issues/5590#issuecomment-624465899)
     - [status-im/specs#73 - Remove certain specs, refer to Waku](https://github.com/status-im/specs/issues/73)
1) [I’ll open a PR making an explicit reference that Envelope is specified by ABNF. Resolving Proposal 4.](https://discuss.status.im/t/wherefore-art-thou-mailserver-treatise-on-waku-terminology/1664/8?u=samuel)
1) Create VAC Waku spec improvment issue for:
   - [Remove MessageResponse version](https://discuss.status.im/t/wherefore-art-thou-mailserver-treatise-on-waku-terminology/1664/3?u=samuel)
   - [Remove Batch Ack packet type](https://discuss.status.im/t/wherefore-art-thou-mailserver-treatise-on-waku-terminology/1664/10?u=samuel)
1) [#1834 - Show pending transactions](https://github.com/status-im/status-go/issues/1834)
   - Todos:
     - [ ] Write specs
     - [ ] Implement
   - Impacts:
     - [status-react#3997 - Display pending requests in transactions history](https://github.com/status-im/status-react/issues/3997)
1) [SNT and emoji reactions**](https://github.com/status-im/status-react/issues/7118)
   - **Notes:**
     - [Andre] all the way from the protocol level to the communication with the react app
     - [Hester] Maciej prepared designs for this, you can ask him anything you need about what the expected behavior is supposed to be
     - [Hester] Also, fyi, in planning this is just Emoji reactions, the original and full feature plan is to include SNT reactions, only for this we first need account contracts so you can set rules for transaction signing
1) [#1937 - Compress public key for chat](https://github.com/status-im/status-go/issues/1937)
   - [x] [Do tests to see average compressed key length](https://github.com/status-im/status-go/issues/1937#issuecomment-624690407)
1) Specs for [status-react#10384 - Protocol specs](https://github.com/status-im/status-react/issues/10384)
1) [Nimbus on mobile](https://discuss.status.im/t/nimbus-on-mobile/1370)

---

## Team

[Top voted features for v1.5+](https://discuss.status.im/t/roadmap-planning/1399/38)

---

## Resolved

- [x] ~~Discuss post on waku mailserver name,~~ [~~change to `history node`, `persistence node`, `echo node` . Some thing more descriptive.~~](https://github.com/status-im/status-go/pull/1949#discussion_r419615374)
- [x] ~~Discuss post on~~ [~~Database package location refactor~~](https://github.com/status-im/status-go/issues/1945)
- [x] [~~#1949 - Waku package README.md documentation~~](https://github.com/status-im/status-go/pull/1949)
