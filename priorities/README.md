# Priorities

## Personal

1) [#1834 - Show pending transactions](https://github.com/status-im/status-go/issues/1834)
  - **Todos:**
    - [ ] Write specs
    - [ ] Implement
  - **Impacts:**
    - [status-react#3997 - Display pending requests in transactions history](https://github.com/status-im/status-react/issues/3997)
     - [status-react#9976 - Show pending transactions](https://github.com/status-im/status-react/issues/9976)
1) [#7118 - SNT and emoji reactions**](https://github.com/status-im/status-react/issues/7118)
   - **Notes:**
     - [Andre] all the way from the protocol level to the communication with the react app
     - [Hester] Maciej prepared designs for this, you can ask him anything you need about what the expected behavior is supposed to be
     - [Hester] Also, fyi, in planning this is just Emoji reactions, the original and full feature plan is to include SNT reactions, only for this we first need account contracts so you can set rules for transaction signing
1) [#1936 - Messengers acquires a lock in some methods for very long time](https://github.com/status-im/status-go/issues/1936)
1) [#1937 - Compress public key for chat](https://github.com/status-im/status-go/issues/1937)
   - **Todos:**
     - [x] [Do tests to see average compressed key length](https://github.com/status-im/status-go/issues/1937#issuecomment-624690407)
       - [Results](https://github.com/status-im/status-go/issues/1937#issuecomment-624920237)
     - [x] [Track down iOS URL intent registry for current build](https://github.com/status-im/status-go/issues/1937#issuecomment-628082382)
       - [See](https://github.com/status-im/status-go/issues/1937#issuecomment-632186000)
1) Specs for [status-react#10384 - Protocol specs](https://github.com/status-im/status-react/issues/10384)
1) [Nimbus on mobile](https://discuss.status.im/t/nimbus-on-mobile/1370)

---

## Team

[Top voted features for v1.5+](https://discuss.status.im/t/roadmap-planning/1399/38)

---

## Reading List

- [ ] [Desktop paths forward, a choose your own adventure?](https://discuss.status.im/t/desktop-paths-forward-a-choose-your-own-adventure/1666)
- [ ] [How :desktop_computer Desktop & Mobile should connect?](https://discuss.status.im/t/how-desktop-mobile-should-connect/1668)
- [ ] [Feature request category](https://discuss.status.im/t/feature-request-category/1698)

---

## Resolved

- [x] Create VAC Waku spec improvement issue(s) for:
  - **Discussed:**
    - [here for discussion](https://discuss.status.im/t/wherefore-art-thou-mailserver-treatise-on-waku-terminology/1664)
  - **Improvements:**
    - [~~Remove MessageResponse version~~](https://discuss.status.im/t/wherefore-art-thou-mailserver-treatise-on-waku-terminology/1664/3?u=samuel)
    - [~~Remove Batch Ack packet type~~](https://discuss.status.im/t/wherefore-art-thou-mailserver-treatise-on-waku-terminology/1664/10?u=samuel)
    - [~~Remove `confimations-enabled` from `status-options`.~~](https://github.com/vacp2p/specs/pull/128#discussion_r427771425)
    - ~~Change name of `status-options` `light-node` to `is-light-node`.~~
- [x] ~~Add to the `status-go/protocol` package README notes about how it was originally a whole repo with managed dependencies.~~ 
  - **PR :** [#1984 - Added basic history of the protocol package](https://github.com/status-im/status-go/pull/1984)
- [x] [~~status-im/specs#61 - Replace forked Whisper usage with Waku/0 in specs~~](https://github.com/status-im/specs/issues/61)
  - **PR :** [#114 - Update/waku replace](https://github.com/status-im/specs/pull/114)
- [x] [~~I’ll open a PR making an explicit reference that Envelope is specified by ABNF. Resolving Proposal 4.~~](https://discuss.status.im/t/wherefore-art-thou-mailserver-treatise-on-waku-terminology/1664/8?u=samuel)
- [x] ~~Discuss post on waku mailserver name,~~ [~~change to `history node`, `persistence node`, `echo node` . Some thing more descriptive.~~](https://github.com/status-im/status-go/pull/1949#discussion_r419615374)
- [x] ~~Discuss post on~~ [~~Database package location refactor~~](https://github.com/status-im/status-go/issues/1945)
- [x] [~~#1949 - Waku package README.md documentation~~](https://github.com/status-im/status-go/pull/1949)
