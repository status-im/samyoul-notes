# Priorities

- [Personal](#personal)
- [Team](#team)
- [Reading List](#reading-list)
- [Resolved](./RESOLVED.md)

## Personal

1. [Emoji Reactions](#emoji-reactions)
1. [Add nim tests for nim-status](#add-nim-tests-for-nim-status)
1. [Pending Transactions](#pending-transactions)
1. [`TRANSACTION_COMMANDS` Specs](#transaction_commands-specs)
1. [SNT Reactions](#snt-reactions)
1. [Messengers acquires a lock for very long time](#messengers-acquires-a-lock-for-very-long-time)
1. [Protocol Specs](#protocol-specs)

---

### Emoji Reactions

[#7118 - SNT and emoji reactions**](https://github.com/status-im/status-react/issues/7118)

- **Notes:**
   - [Andre] all the way from the protocol level to the communication with the react app
   - [Hester] Maciej prepared designs for this, you can ask him anything you need about what the expected behavior is supposed to be
     - [UI Designs](https://www.figma.com/file/aS1ct66VQ6V0cio7vSqS8UoG/Chat?node-id=1239%3A0)
   - [Hester] Also, fyi, in planning this is just Emoji reactions, the original and full feature plan is to include SNT reactions, only for this we first need account contracts so you can set rules for transaction signing

### Add nim tests for nim-status

[nim-status#9- add nim tests for nim-status](https://github.com/status-im/nim-status/issues/9)

- [Add tests & documentation - accounts methods (delete import, derive/generate new accounts, identicon, alias)](https://github.com/status-im/nim-status/issues/15)
- [Add tests & documentation - Mailservers methods](https://github.com/status-im/nim-status/issues/18)

### Pending Transactions

[#1834 - Show pending transactions](https://github.com/status-im/status-go/issues/1834)

- **Dependencies:**
  - [`TRANSACTION_COMMANDS` Specs](#3-transaction_commands-specs)
- **Todos:**
  - [ ] Write specs
    - https://github.com/ethereum/devp2p/blob/master/caps/les.md
    - https://openethereum.github.io/wiki/Light-Ethereum-Subprotocol-(LES)
    - https://discuss.status.im/search?q=ultra%20light%20client
  - [ ] Implement
- **Impacts:**
  - [status-react#3997 - Display pending requests in transactions history](https://github.com/status-im/status-react/issues/3997)
  - [status-react#9976 - Show pending transactions](https://github.com/status-im/status-react/issues/9976)

### `TRANSACTION_COMMANDS` Specs

- Transaction Commands
  - [status-im/team-core#7 - Allow for smooth sending and receiving of SNT while maintaining privacy](https://github.com/status-im/team-core/pull/7)
    - Original specs for `TRANSACTION_COMMANDS` 
  - https://github.com/status-im/specs/blob/master/docs/stable/6-payloads.md
    - Where specs for `TRANSACTION_COMMANDS` needs to go

### SNT Reactions

[#7118 - SNT and emoji reactions**](https://github.com/status-im/status-react/issues/7118)

- **Notes:**
  - [See Emoji reactions](#4-emoji-reactions)

### Messengers acquires a lock for very long time

1) [#1936 - Messengers acquires a lock in some methods for very long time](https://github.com/status-im/status-go/issues/1936)


### Protocol Specs

Specs for [status-react#10384 - Protocol specs](https://github.com/status-im/status-react/issues/10384)

---

## Team

[Top voted features for v1.5+](https://discuss.status.im/t/roadmap-planning/1399/38)

---

## Reading List

- [ ] [Why ‘Stimbus’ or Organisational Focus and Alignment](https://discuss.status.im/t/why-stimbus-or-organisational-focus-and-alignment/1753)
- [ ] [How Stimbus or what is needed to start this work](https://discuss.status.im/t/how-stimbus-or-what-is-needed-to-start-this-work/1754)
