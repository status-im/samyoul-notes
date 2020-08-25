# 2020-08-25

## Pulls

- [x] [status-im/specs#145 - Added EmojiReaction payload types](https://github.com/status-im/specs/pull/145) - `commits`, `merged`
- [x] [status-im/specs#146 - Updated release commit data](https://github.com/status-im/specs/pull/147/files) - `submitted`

## Reviews

- [x] [status-im/go-ethereum#90 - Merge patches 0000, 0040, and 0044 #1772](https://github.com/status-im/go-ethereum/pull/90) - `approved`
- [x] [#2023 - [#11046] Add local contact names](https://github.com/status-im/status-go/pull/2023) - `feedback`
  - Rename "Nickname" to "LocalNickname" or add documentation to make explicit that a "Nickname" is only for local use
    - This will make a distinction from "DisplayNames" in the case these are implemented. 
- [x] [#2013 - Feature/group chat invitation](https://github.com/status-im/status-go/pull/2013) - `feedback`

---

# 2020-08-24

## Issues

- [x] [status-react#11047 - Add profile photo](https://github.com/status-im/status-react/issues/11047) - `discussed`
  - https://notes.status.im/vk29m0ZjT1WClNR1BR8msg?both

## Schedule

- [x] []()
// Add Dev call meeting notes

---

# 2020-08-21

## Pulls

- [x] [Status Avatars (Prototype)](https://github.com/status-im/avatar) - `created`

---

# 2020-08-20

## Schedule
// Add Referal meeting

- [x] [status-react#11047 - Add profile photo](https://github.com/status-im/status-react/issues/11047) - `discussed`
// Add spec v2

---

# 2020-08-19

## Issues

- [x] [status-react#11047 - Add profile photo](https://github.com/status-im/status-react/issues/11047) - `discussed`
// Add spec v2

## Schedule
// Add Andre meeting

---

# 2020-08-18

## Issues

- [x] [status-react#11047 - Add profile photo](https://github.com/status-im/status-react/issues/11047) - `discussed`, `scoped`, `draft`
   - [Created specs v1](https://notes.status.im/CQ-GGYmAR3aM8qUABgedWg?both)

## Reviews

- [x] [#2013 - Feature/group chat invitation](https://github.com/status-im/status-go/pull/2013) - `feedback`
- [x] [#2009 - Add topic and contact code propagation](https://github.com/status-im/status-go/pull/2009) - `discussed`
- [x] [#2017 - Fix mark messages seen](https://github.com/status-im/status-go/pull/2017) - `discussed`
- [x] [#2018 - Handle wallet initialization](https://github.com/status-im/status-go/pull/2018) - `approved`

---

# 2020-08-17

## Reviews

- [x] [#2009 - Add topic and contact code propagation](https://github.com/status-im/status-go/pull/2009) - `approved`, `feedback`

## Schedule

- [x] 14:30 BST - Referral program implementation sync
  - https://meet.google.com/nfd-pare-jke

---

# 2020-08-14

## Schedule

- [x] 14:00 BST - Core Retro
   - https://meet.google.com/hqw-nsfk-xoq
   - **Notes :**
     - https://notes.status.im/uqCj7Q6lTIqslxKUlmxELQ?both

---

# 2020-08-13

## Pulls

- [x] [#2016 - Fixes a minor rebase quirk](https://github.com/status-im/status-go/pull/2016) - `submitted`, `merged`

---

# 2020-08-12

- [x] 14:00 BST - Kick-off 'In-app notifications'
   - https://meet.google.com/pxe-obgh-udo
   - **Notes :**
     - Epic [In-app notifications](https://github.com/status-im/status-react/issues/10547)
     - [Feature Document](https://docs.google.com/document/d/162zAXBtTPQeKKvxccW7ukC8I2a2ZhBgTOVDPT-JWvOA/edit#)
- [x] 15:00 BST - Kick off 'Profile photo'
   - https://meet.google.com/edn-eonp-csm
   - **Notes :**
     - Epic [Add profile photo](https://github.com/status-im/status-react/issues/11047)
     - [Feature Document](https://docs.google.com/document/d/1FeTNuluM7rF7kAIkc38rE1p7gTwirZHhPC-Q7P0FTwM/edit#heading=h.3o4bb8dnf05)
     - [Miro notifiable events](https://miro.com/app/board/o9J_kn6l1NQ=/)
     - [Designs](https://www.figma.com/file/aS1ct66VQ6V0cio7vSqS8UoG/Chat?node-id=1405%3A2407)
     - Discussion of user display names:
       - The group agreed that adding only profile photos would be seen as delivering half a feature, and that the other half of this feature is allowing the user the ability to chose their own display name.
       - Need to write document detailing
         - Threat modeling for display name
         - Show rationale for why a display name would be good
         - Add an element of uniqueness, `<display_name> + <first_8_chars_of_pubkey>`, "Jupiter King 0xeb18d99c..."
- [x] 16:00 BST - Referral program implementation sync
   - https://meet.google.com/nfd-pare-jke 

---

# 2020-08-11

## Schedule

- [x] 14:00 BST - Samuel / André 1:1
   - https://meet.google.com/pga-udvu-wpk
   - **Notes :**
     - Contact name functionality
     - Chat in Status
     - Alice is Missing
     - Drivethrough RPG
- [x] 14:30 BST - Kick-off 'Status chat inside dapps'
   - https://meet.google.com/umi-xeiu-hso
   - **Notes :**
     - **Feature Epic** - [Status chat inside dapps](https://github.com/status-im/status-react/issues/11044)
     - **Waku in browser solutions**
       - https://github.com/status-im/murmur
       - [Waku Typescript lib](https://github.com/vacp2p/waku-ts)
     - **Pre-existing proposals**
       - Embeddable read only feeds
       - [Status APIs](https://notes.status.im/z0AMSnl-S7mTqQR9FfR3ng#)
       - [`status-react` - Chat api](https://github.com/status-im/status-react/pull/10910)
     - **Private key management**
       - https://walletconnect.org/
     - **Potential use cases**
       - https://gov.yearn.finance/t/idea-move-to-gasless-voting/1039/14
       - https://github.com/status-im/gasless-democracy

---

# 2020-08-10

## Schedule

- [x] 13:00 BST - Bi-weekly Core Dev Call
   - https://meet.google.com/uyn-yfrq-uka
   - **Notes :**
     - https://boardgamegeek.com/boardgame/159473/quartermaster-general
     - https://github.com/status-im/status-react/issues/8512#issuecomment-670151370
     - https://3box.io/products/profiles
     - https://github.com/status-im/status-react/pull/10910
- [x] 15:00 BST - Kick-off 'Local names'
   - https://meet.google.com/pwz-spxv-cxj
   - **Notes :**
     - https://github.com/status-im/status-react/issues/11046
     - https://www.figma.com/file/TNCyHKtR3sx5EL6YznFWUa4O/Profile?node-id=4066%3A7400

## Issues

- [x] [Add local contact names](https://github.com/status-im/status-react/issues/11046) - `discussed`, `scoping`
  - [Added responses to team concerns](https://github.com/status-im/status-react/issues/11046#issuecomment-670623942)     

---

# 2020-08-07

## Reviews

- [x] [#2013 - Feature/group chat invitation](https://github.com/status-im/status-go/pull/2013) - `feedback`

---

# 2020-08-06

- [x] 13:00 BST - Roadmap planning workshop - Next steps
   - https://meet.google.com/rrx-rjwg-qpg
   - https://docs.google.com/spreadsheets/d/1w6ohP7RQY6am4Vpac-43YUuBIIENIe_h09GfESCPy1k/edit#gid=1462389450

---

# 2020-08-05

## Schedule

- [x] 16:00 BST - Samuel / André 1:1
   - https://meet.google.com/rij-fgdg-xxe

---

# 2020-08-04

## Schedule

- [x] 15:00 - 17:00 BST - Mission Sync Hester / Samuel
  - https://meet.google.com/exz-bqvi-hwa
  - **Notes**
    - [Roadmap discuss](https://discuss.status.im/t/retention-focus-and-roadmaps/1826/29)
    - [feature map. Feasibility vs impact](https://drive.google.com/file/d/1ZljKYBi5wZ8MvFJpfhgZ0wKoga2We7Yu/view)
    - [Feature votes](https://cdn.discordapp.com/attachments/701847687674331177/740209062544932935/Screenshot_2020-07-27_at_16.19.18.png)
    - [Market Research - Community manager interviews](https://docs.google.com/spreadsheets/d/1xGgysilDM_rUhr6GkrBVKJgq37MkZpzWDNfPzB4FL-g/edit#gid=0)
    - [Telegram widgets](https://core.telegram.org/widgets/posts)
    - [Contract Notifications Service](https://github.com/status-im/contract-notifier)
---

# 2020-08-03

## Schedule

- [x] 14:30 BST - Referral program implementation sync
   - https://meet.google.com/nfd-pare-jke

