# 2020-06-10

## Schedule

- [x] 11:30 BST - Rajeev / Sam 1:1
   - https://meet.google.com/ejo-csxx-ixi
   - [specs#123 - Clarification on Status node roles](https://github.com/status-im/specs/issues/123)
- [x] 16:00 BST - Samuel / André 1:1
   - https://meet.google.com/rij-fgdg-xxe
   - Notes
     - [#1834 - Show pending transactions](https://github.com/status-im/status-go/issues/1834)
     - [#1937 - Compress public key for chat](https://github.com/status-im/status-go/issues/1937)
     - Review feedback
     - Stimbus

## Posts

- [x] [Full Status Archive Nodes - A path to independence and true decentralisation](https://discuss.status.im/t/full-status-archive-nodes-a-path-to-independence-and-true-decentralisation/1747) - `created`

## Reviews

- [x] [vacp2p/specs#142 - rename mailserver to history node](https://github.com/vacp2p/specs/pull/142) - `approved`
- [x] [status-im/specs#105 - mentions](https://github.com/status-im/specs/pull/105) - `feedback`

---

# 2020-06-09

## Schedule

- [x] 12:00 BST - Samuel / Terry - Catch up
   - https://meet.google.com/acp-ftky-cqe

## Pulls

- [x] [#1990 - Feature/key compression](https://github.com/status-im/status-go/pull/1990) - `draft`, `commits`
   - I had a lot of problems with `vendor make`
     - I needed to install `goware/modvendor` directly, `modvendor-install` was failing
   - Package unit testing caused a lot of problems because C libs were/are throwing errors.
     - For details [see here](attachments/2020-06/2020-06-09_unit-test-error-log.txt)
   - The linter is capricious and likely evil, I spent too much time trying to appease it
     - The linter seems to want every non-native package to be listed separately grouped by repo

## Reviews

- [x] [vacp2p#141 - Add rate limit by size](https://github.com/acp2p/specs/pull/141) - `reapproved`
- [x] [#1989 - Implement rate limiting by bytes](https://github.com/status-im/status-go/pull/1989) - `approved`
- [x] [#1987 - expose getContactByID method](https://github.com/status-im/status-go/pull/1987) - `approved`

---

# 2020-06-08

## Schedule

- [x] 13:00 BST - Town Hall
   - https://www.youtube.com/c/Statusim/live

## Issues

- [x] [#1834 - Show pending transactions](https://github.com/status-im/status-go/issues/1834)
   - Sync with Jarrad about costing for experimentation and rollout plan discussed with Jakub
- [ ] [#1937 - Compress public key for chat](https://github.com/status-im/status-go/issues/1937)
   - Implement binding functions for compressing and uncompressing public keys
- [x] [#1936 - Messengers acquires a lock in some methods for very long time](https://github.com/status-im/status-go/issues/1936)
   - Begin review of code for improvements

## Reviews

- [x] [vacp2p#141 - Add rate limit by size](https://github.com/acp2p/specs/pull/141) - `feedback`, `discussed`, `approved`
- [x] [#105 - mentions](https://github.com/status-im/specs/pull/105) - `feedback`

---

# 2020-06-05

## Schedule

- [x] 14:00 BST - Core Retro
  - https://meet.google.com/hqw-nsfk-xoq
  - [Notes](https://notes.status.im/mj_3vzVQQHiQC0C7DOiZSA?both)
    - Got go ahead for our own infura light, provided cost is not too high ... Need to get costings
      - [x] Samuel to schedule exploratory meeting
- [x] 15:00 BST - Nix in Status
   - https://meet.google.com/kwv-arcf-jqm
   - https://www.youtube.com/watch?v=m4sv2M9jRLg
   - https://drive.google.com/file/d/1Tt5R7QOubudGiSuZIGxuFWB1OYgcThcL
- [x] 16:00 BST - Jakub / Samuel
   - https://meet.google.com/tgb-bjqu-ehr
   - https://notes.status.im/pending-tx-infra-costs
   - https://github.com/status-im/jakubgs-notes/blob/master/research/pending_tx_infra.md

## Issues

- [x] [#1937 - Compress public key for chat](https://github.com/status-im/status-go/issues/1937) - `research`
  - Researching how to create new bindings in the `status-go` binary

## Reviews

- [x] [Adjust index to use `local_chat_id`](https://github.com/status-im/status-go/pull/1986) - `approved`, `feedback`

---

# 2020-06-04

## Schedule

- [x] 12:30 BST - Jarrad and Sam Compressed Keys
   - https://meet.google.com/dnv-otdx-jmp
   - Notes
     - [#1937 - Compress public key for chat](https://github.com/status-im/status-go/issues/1937)
       - Treat this improvement for aesthetic reasons
       - https://github.com/multiformats/multiformats
     - [#1834 - Show pending transactions](https://github.com/status-im/status-go/issues/1834)
       - Implement an archive node network.
       - Light node implementations / querying state
         - https://slock.it/incubed/
         - https://slock.it/incubed-sdk/
         - https://thegraph.com/

## Issues

- [x] [#1834 - Show pending transactions](https://github.com/status-im/status-go/issues/1834) - `discussed`
- [x] [#1937 - Compress public key for chat](https://github.com/status-im/status-go/issues/1937) - `discussed`, `research`
  - [Multiformat public keys](https://github.com/status-im/status-go/issues/1937#issuecomment-638963337)

## Pulls

- [x] [#1984 - Added basic history of the protocol package](https://github.com/status-im/status-go/pull/1984) - `merged` 

---

# 2020-06-03

## Issues

- [x] [#1937 - Compress public key for chat](https://github.com/status-im/status-go/issues/1937) - `research`, `discussed`
   - [Theoretical min and max base58 encoded compressed key lengths](https://github.com/status-im/status-go/issues/1937#issuecomment-638286734)
- [x] [vacp2p/specs#137 - Remove / deprecate Batch Ack packet type](https://github.com/vacp2p/specs/issues/137) - `created`
- [x] [vacp2p/specs#138 - Remove MessageResponse version](https://github.com/vacp2p/specs/issues/138) - `created`
- [x] [vacp2p/specs#139 - Remove confirmations-enabled from status-options](https://github.com/vacp2p/specs/issues/139) - `created`
- [x] [vacp2p/specs#140 - Change name of status-options light-node to is-light-node](https://github.com/vacp2p/specs/issues/140) - `created`

## Pulls

- [x] [status-im/specs#132 - Remove personal pronouns](https://github.com/status-im/specs/pull/132/) - `commits`, `merged`
- [x] [#1984 - Added basic history of the protocol package](https://github.com/status-im/status-go/pull/1984) - `commits`, `submitted` 

## Reviews

- [x] [go-maven-resolver#3 - Add error handling](https://github.com/status-im/go-maven-resolver/pull/3) - `feedback`, `approved`
- [x] [vacp2p/specs#135 - feature/waku-rpc](https://github.com/vacp2p/specs/pull/135) - `feedback`, `approved`
- [x] [#1985 - use EncodeToString to print []byte variables](https://github.com/status-im/status-go/pull/1985) - `approved`

---

# 2020-06-02

## Schedule

- [x] 14:00 BST - Samuel / André 1:1
   - Async ~~https://meet.google.com/pga-udvu-wpk~~
   - [#1834 - Show pending transactions](https://github.com/status-im/status-go/issues/1834)
     - Do light nodes help provide a solution for this

## Issues

- [x] [status-im/specs#134 - Implement Passive Voice / General Grammar Linter](https://github.com/status-im/specs/issues/134) - `created`

## Pulls

- [x] [status-im/specs#132 - Remove personal pronouns](https://github.com/status-im/specs/pull/132/) - `commits`, `resubmitted` 

## Reviews

- [x] [go-maven-resolver#2 - add request retries for non-404 errors](https://github.com/status-im/go-maven-resolver/pull/2) - `approved`

---

# 2020-06-01

## Schedule

- [x] 13:00 BST - Bi-weekly Core Dev Call
   - https://meet.google.com/uyn-yfrq-uka
   - [Notes](https://notes.status.im/core-dev-call_31?both)
   - [Youtube - Core Dev Call #31 - June 1, 2020](https://www.youtube.com/watch?v=yEKMJ1cdfRg)
- [x] 14:00 BST - Discussion of decentralised incoming pending transactions
  - [#1834 - Show pending transactions](https://github.com/status-im/status-go/issues/1834)
  - [Youtube - Part of the conversation](https://youtu.be/yEKMJ1cdfRg?t=2784)

## Issues

- [x] [#1834 - Show pending transactions](https://github.com/status-im/status-go/issues/1834) - `discussed`
- [x] [status-im/pm#56 - Core Dev Call 31 Agenda](https://github.com/status-im/pm/issues/56) - `discussed`

## Pulls

- [x] [status-im/specs#132 - Remove personal pronouns](https://github.com/status-im/specs/pull/132/) - `discussed`

## Reviews

- [x] [go-maven-resolver#2 - add request retries for non-404 errors](https://github.com/status-im/go-maven-resolver/pull/2) - `feedback`

---