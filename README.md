[[INDEX.md]]
# 2024-10-09
## Schedule
- Status-go guild sync:
	- [[️status-go guild Weekly Sync 2024-10-09]]
## Program
### IFT Buddy program
- Finalised document based on feedback from Ícaro and Arwen:
- [[PROPOSAL Cross-IFT Buddy Program]]

---
# 2024-10-08
## Schedule
- Mobile Core sync and planning:
	- MixPanel via proxy
	- Critical path errors should be reported via sentry.
	- Wallet integrations asking permissions like Gifs
	- [ ] Create feature request for ask permissions to expose IP addresses in wallet in GitHub. (✅ Added)
	- [ ] Set up a call with Waku to establish a process for requesting specific items: (✅ Added)
		- Stable test net
		- Store node:
			- See what is the connected node
			- have the option to select our own node
			- Why was it removed previously?
	- Set up some crews:
		- Get Volo involved and introduced
		- a quick and simple idea is to create teams in GH, that is crew teams. So that QAs or anybody can tag the crew in issues.
## Program

### Minor Update on Legal Docs in the Apps

- 🦸 Legal super powers:
  - The `Legal` Github team now have `CODEOWNERS` power over the whole repo and their approval is required before any legal documents are updated.
  - Thank you :pray: @siddarthkay for pressing the On button. :grimacing: 
  - https://github.com/status-im/status-software-legal-documents/pull/4#issuecomment-2398830820
- 🧑‍⚖️ Source of Truth:
  - After some discussion with @siddarthkay and @jakubgs, we've established a feasible method of simply maintaining versions of the legal documents in the repos. AND we can ensure that repo versions of the legal documents are identical to the source of truth documents.
  - [Feature Request: Use Git Submodule for Legal Documents Comparison #5](https://github.com/status-im/status-software-legal-documents/issues/5)
  - Next steps are mentioned in the issue but will require involvement from all platform teams.
- 💪 `terms-of-use.md`:
  - Thank you @flexsurfer for this PR introducing the `terms-of-use.md` document to the repo:
  - This resolves issue [Create `terms-of-use.md` file with correct contents #2](https://github.com/status-im/status-software-legal-documents/issues/2)
  - @e.nguye101 this PR now requires your approval ✅ [Create terms-of-use.md #4](https://github.com/status-im/status-software-legal-documents/pull/4)
- 🚧 `privacy-policy.md`:
  - This PR is still open and requires someone (ideally a non-mobile dev for cross Status participation and visibility) [Add the Latest Privacy Policy to `privacy-policy.md` #3](https://github.com/status-im/status-software-legal-documents/issues/3)
  - Any volunteers? ( @shyolghul @0x70atryk @helium__ @anthony3992 )

---
# 2024-10-07
## Schedule
- James from Status Chain
## Issues
- [Feature Request: CI for Ensuring No Diff Between Platform and Legal Docs](https://github.com/status-im/status-software-legal-documents/issues/5) #created 
---
# 2024-10-04
## Schedule
- Cross Mobile and Status Chain Sync
	- Cyp's bit:
		- [AURA tokenomics](https://www.notion.so/AURA-tokenomics-10d8f96fb65c80299329ced4eafbd58b)
		- Some suggestions from me:
			- Niche use case a "Social L2 Network":
			- ENS+ - Identity attestations
			- Social badges
			- [Steering Committee Synergy Log 2024-09](https://docs.google.com/document/d/1TYWvktmtgCtZykGmGiFxQzKTAyr6V55N81ZIUM3bvOU/edit?authuser=2) (page 17)
			- Content sharing monetisation
			- DAO out of the box, link with Status communities
			- Crowd funding and tipping
			- Tribute to talk (pay to get a highlighted message in someone's DMs)
			- Proof of membership / participation protocol
			- (Cyp said these were the focus of the L2)
		- Semi-Optimism bounties for projects building within a certain type:
			- [x] Suggest this  
			- (Don't do this, mercenary teams can just tick boxes and not add value)
			- Optimism bounties are the best as teams that add value can be rewarded only after their impact has been proven.
			- Social
			- Gaming
			- Identity
	- Legal document in 2.31:
		- We could do a content hash rather than a full file hash and that should create identical hash
# 2024-10-03
## Schedule
- Call with Volo
	- Discussed a number of topics:
	- 
# 2024-10-02
## Schedule
- Status-go Guild Weekly Sync
	- [[️status-go guild Weekly Sync 2024-10-02]]
	- Action points:
		- - Work on getting Policies
		    - 0 - about the form of agreement on all further policies
		    - 1 - Tests policy
		    - 2 - Breaking changes policy:
# 2024-10-01
## Ad Hoc
- Got back to Volo about permissions issue on this https://github.com/status-im/least-authority-audit.
	- I believe he needs to check that his Github account is part of the "`Status`" team on the `status-im` repo
