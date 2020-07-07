# Resolved Priorities

## Reading List

- [x] `2020-07-07` - [Feature request category](https://discuss.status.im/t/feature-request-category/1698) 
- [x] `2020-07-01` - [Desktop paths forward, a choose your own adventure?](https://discuss.status.im/t/desktop-paths-forward-a-choose-your-own-adventure/1666)
- [x] `2020-07-01` - [How :desktop_computer Desktop & Mobile should connect?](https://discuss.status.im/t/how-desktop-mobile-should-connect/1668)

## Tasks

### Compress public key for chat

[#1937 - Compress public key for chat](https://github.com/status-im/status-go/issues/1937)

- **Progress:**
  - [x] [status-go#1990 - Feature/key compression](https://github.com/status-im/status-go/pull/1990)
  - [x] [status-im/specs#137 - Added public key compression specs](https://github.com/status-im/specs/pull/137)
- **Todos:**
  - [x] [Do tests to see average compressed key length](https://github.com/status-im/status-go/issues/1937#issuecomment-624690407)
    - [Results](https://github.com/status-im/status-go/issues/1937#issuecomment-624920237)
  - [x] [Track down iOS URL intent registry for current build](https://github.com/status-im/status-go/issues/1937#issuecomment-628082382)
    - [See](https://github.com/status-im/status-go/issues/1937#issuecomment-632186000)
  - [x] [Research base58 key length](https://github.com/status-im/status-go/issues/1937#issuecomment-638286734)
  - [x] [Research Multiformat for key versioning](https://github.com/status-im/status-go/issues/1937#issuecomment-638963337)
  - [x] Once specs are merged add link to the specs in the `status-go` code
  - [x] Implement multiformat parsing for website
    - See https://github.com/status-im/universal-links-handler/pull/42
- **Impacts**
  - [x] [#10325 - Request compressed keys](https://github.com/status-im/status-react/issues/10325)

---

### Older Items  

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
