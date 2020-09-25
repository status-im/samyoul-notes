# Priorities

- [Personal](#personal)
- [Team](#team)
- [Reading List](#reading-list)
- [Resolved](./RESOLVED.md)

## Personal

1. [1.8 - Add profile photo](#add-profile-photo)
1. [Improved device syncing](#improved-device-syncing)
1. [Pending Transactions](#pending-transactions)
1. [`TRANSACTION_COMMANDS` Specs](#transaction_commands-specs)
1. [SNT Reactions](#snt-reactions)
1. [Messengers acquires a lock for very long time](#messengers-acquires-a-lock-for-very-long-time)
1. [Protocol Specs](#protocol-specs)
1. [1.8 - Status chat inside dapps](#status-chat-inside-dapps)

---

### Improved device syncing

- **References**
  - [How desktop_computer Desktop & Mobile should connect?](https://discuss.status.im/t/how-desktop-mobile-should-connect/1668)
  - [Figma design for sync](https://www.figma.com/file/Mr3rqxxgKJ2zMQ06UAKiWL/%F0%9F%92%AC-Chat%E2%8E%9CDesktop?node-id=1944%3A0)
- **Action Points**
  - Make process flow for QR code sync
  - Make process flow unsyncron
  - Make process flow for installation of sync multiple
  - make process flow for unsyncron for multiple devices

### Add profile photo

[status-react#11047 - Add profile photo](https://github.com/status-im/status-react/issues/11047)

- **Notes :**
  - Avatars
    - [Avatar introduction](https://notes.status.im/vk29m0ZjT1WClNR1BR8msg?both)
    - [Avatar renderer experiment / prototype](https://github.com/status-im/avatar)
  - Profile Images
    - [Research on vector based profile images](https://notes.status.im/0MWyryTPQleMWe31K0L7gg?both)
    - [Profile images spec V2](https://notes.status.im/oUChlPB3Q2aUPLYBSjq6MQ?both)
    - [Profile images spec V1](https://notes.status.im/CQ-GGYmAR3aM8qUABgedWg?both)

### Status chat inside dapps

[#11044 - Status chat inside dapps](https://github.com/status-im/status-react/issues/11044)

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
- [ ] [The Keys to Privacy at Status](https://docs.google.com/document/d/1r8tHGYiWw1__uy8b-T8FMrf-VACXJeV4yPMheEpRXsQ/edit#heading=h.foffmq12en67)
- [ ] [User growth and retention - comment](https://discuss.status.im/t/user-growth-and-retention/1782/23)
- [ ] [On remote communication](https://discuss.status.im/t/on-remote-communication/1819)
