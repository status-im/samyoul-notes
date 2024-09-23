# 2024-09-23
## Issues
- [Data collection and terms&conditions are the first major impresisons of opening the client #21293](https://github.com/status-im/status-mobile/issues/21293)
	- Chu pointed out this issue
	- Discussed internally
	- Further reading suggests we may want to address this issue to the product council

## Reviews
- [# Feat/log and panic #5849](https://github.com/status-im/status-go/pull/5849#pullrequestreview-2321567843) - #approved 
	- Added a couple of points about `defer` stack order and checked when should the `common.LogOnPanic()` function be called in the `defer` stack (first?).
	- Had some [additional conversation](https://github.com/status-im/status-go/pull/5849#discussion_r1771068007) about the order in which panics and defers interact
		- See https://go.dev/play/p/A2ik-YSGsWW
		- This is potentially a problem
## Ad Hoc

- Send Arwen the Cross IFT Suggestions [[#Cross IFT Suggestions]]
- Did a massive file removal from the notes:
	- `git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch attachments/2024-08/2024-08-08_access.log' --prune-empty --tag-name-filter cat -- --all`
### Legal Repo
- New explicit and accurate name: "status software legal documents"
  - https://github.com/status-im/status-software-legal-documents
- Use a flat directory
- Added the Legal team as `CODEOWNERS`
  - https://github.com/status-im/status-software-legal-documents/blob/master/.github/CODEOWNERS
  - Asked Eric if he would be able to add as necessary any other legal team members to the group you created? https://github.com/orgs/status-im/teams/legal
- Added a branch rule on `master` that requires all PRs to be reviewed by the code owners.
- Opened the first PR, requires approval from Legal team:
  - https://github.com/status-im/status-software-legal-documents/pull/1
- Asked if someone could open another PR to create a `terms-of-use.md` file?
  - Just to check that I've set the permissions up correctly and it works as expected.
  - After I can remove myself as an `admin`.

---
# 2024-09-20

## Ad Hoc

### Cross IFT Suggestions

- **Cross IFT Buddies:**
	- Opt-in buddy system for teams to learn about each other; pilot phase followed by broader rollout.
- **Collaboration Hub:**
	- Forum (or forum topic) to propose and curate collaboration and integration opportunities across IFT projects. Open to all CCs.
- **Knowledge Shares:**
	- Regular team presentations to share learnings and achievements, fostering transparency and awareness of other teams' activities.
- **Team Swap:**
	- Short-term skill exchanges between teams (20%–40% involvement), enabling cross-team development and direct participation in other projects.

---
# 2024-09-19
## Ad Hoc
### New legal repo
I created a new repo for the privacy policy single source of truth:
- https://github.com/status-im/status-software-privacy-policy
- I'll log this in a document also as everyone will need to know:
	- We have a single source of truth for privacy policies (and potentially any other legal policy documents)
	- When our platform applications need to reference the privacy policy they should pull from this repo.
	- The input for this repo will be communicated by the Legal team
	- We don't make changes to this repo without the Legal team approving it.

---
# 2024-09-16

## Schedule
### **All IFT Town Hall** IFT Strategy Alignment

- Monday, 16 September⋅14:00 – 15:00
- Big organisation wide structure
- https://streamyard.com/watch/evJU5WHQhSw4
- https://www.notion.so/IFT-Townhall-Recordings-Archive-995e820fd1d14ba3bff7f001a09b7e66
- Part 3: Documentation:
	- Where can I learn more about the Rhythm of Business? Learn more about the ROB here: https://www.notion.so/IFT-Rhythm-of-Business-e61b16873554470d892b20d70fbf014f#cbb05629eeea47e69d555db7badd6d38
	- Where can I find the IFT Annual strategy? Read the 2024-2025 IFT strategy here: https://www.notion.so/33838150ef0b42b3993aa2a963e15503?pvs=24#17f148a5a00c44869ed566c0d3fe04d8
- Documentation - Go deeper:
	- NCT Deepdive: https://www.notion.so/IFT-NCT-Process-2024-2025-f6faa60ab23c421d8a7fd339d835054f#7ee3cb93bb8f4ea2a7a56d1433ddf6d2 
	- Project Steering Committee Workspace: The document is open for feedback today and tomorrow (16, 17 September). The goal of the September meeting is to establish a baseline and kick-off the work of the committee. Notion, here: https://www.notion.so/Workspace-Project-Steering-Committee-fd0af02def0c459cb9cf4de7f68dc5bf. Document, here: https://docs.google.com/document/d/1TYWvktmtgCtZykGmGiFxQzKTAyr6V55N81ZIUM3bvOU/edit 
	- IFT Steering Committee Workspace: The document will be open for feedback next week (23, 24 September). Narratives will be finalised that same week by the committee. All IFT Studio Teams set commitments and tasks between 30 September and 4 October. Notion, here: https://www.notion.so/Workspace-IFT-Steering-Committee-967af9f4f4754a1fa98b18971d7c4715
- __“I still have questions!”__
	- If you still have questions on decision-making that are relevant to the whole org, add them to this document: https://docs.google.com/document/d/1zieqnRMtplw5XLxMypmhFZdNoPjrmAcLkIKZzjQsNz8/edit. I’ll personally go through every relevant question before the October Townhall (14 October).
### Wallet pentesting testing discussion
- Monday, 16 September⋅15:00 – 16:00
- [Transcript](https://docs.google.com/document/d/1VSy2nGP42Wq51MrvYk7Pi_9tRgKa8-udzAB7F1eRe6k/edit#heading=h.ovdducvs67nx)
- Summary
	- In this meeting, the team discussed the status of wallet security testing. They acknowledged that the wallet was released without a formal audit, unlike previous releases, and decided to conduct internal pentesting before seeking external certification. They considered using a familiar security firm, Least Authority, for an external audit of a specific wallet version. The team addressed concerns about which code version to test and agreed to finalize a timeline. Next steps include reaching out to the firm for quotes and scheduling further meetings to discuss progress
- [Least Authority audit of Keycard integration of Status Desktop](https://github.com/status-im/least-authority-audit/blob/master/initial-report/Least%20Authority-%20Status%20Desktop%20Client%20and%20Keycard-go%20Initial%20Audit%20Report.pdf)
- 