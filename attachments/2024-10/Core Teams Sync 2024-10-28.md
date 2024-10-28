# Overview
The Core Teams Sync meeting for Desktop and Mobile focused on several critical issues affecting project progress.

A significant delay in Mixpanel event tracking was identified, leading Jonathan to suggest testing with a new account while Ícaro volunteered to review the code for errors.

The mobile team has successfully completed 30% of the regression test suite, although they encountered new blockers; the desktop team faced minor issues related to logging improvements.

Discussions included simplifying the onboarding flow and username requirements, contemplating the use of public key alternatives.

Concerns about preserving protocol knowledge after Andrea's departure highlighted the need for updated documentation and better specs. Additionally, Igor announced an upcoming call with Jacek to discuss planning for Status Go's future, speculating on a possible rewrite in Nim.

The removal of ID verification code from the desktop was confirmed by Jonathan.

Proposal to enhance transparency through a migration of documentation from Notion to GitHub.

# Notes
## 🌟 Mixpanel Event Tracking Issue (00:03 - 06:51)
- Significant delay (40-50 minutes) in Mixpanel event tracking identified
- Team using free plan of Mixpanel, possibly causing throttling
- Jonathan suggested testing with a new account to isolate the issue
- Ícaro offered to review the code and check for logging errors
## 📊 Release Progress Update (06:57 - 17:11)
- Mobile team reached 30% of regression test suite
- New blocker identified on mobile side (not status go related)
- Desktop team reported multiple issues, but nothing crucial
- Get logger being replaced with sublogger in status go codebase
- Discussion on improving log analysis tools for better debugging
## 🔄 Onboarding Flow and Username Requirements (17:31 - 30:03)
- New simplified onboarding flow discussed
- Concern raised about status go potentially requiring display name
- Suggestion to use public key or compressed key characters as default username
- Debate on using three-word mnemonics vs. public key for default usernames
## 🧠 Protocol Knowledge and Documentation (30:03 - 40:15)
- Concern about loss of protocol knowledge after Andrea's departure
- Discussion on updating and maintaining documentation for double ratchet and other protocols
- Suggestion to reach out to Andrea for help in documenting some aspects
- Need for better specs to allow potential rebuilding of the entire system
## 📞 Upcoming Call with Jacek (40:15 - 46:43)
- Igor announced a call with Jacek to discuss Status Go's future
- Speculation about potential suggestions for rewriting in Nim
- Importance of having clear specs for any potential rewrite
## 🔒 ID Verification Code Removal (46:48 - 55:48)
- Jonathan removed ID verification code from desktop (1500 lines)
- Discussion on keeping only local trust state (trust/untrust button)
- Sam shared notes on feature lifecycle process changes
- Proposal to migrate from Notion to GitHub for better transparency
# Action items

## Tetiana
Create an issue regarding Mixpanel delay with logs (09:05)
## Ícaro
Review Mixpanel code and check for logging errors (08:37)
## Igor
Check if display name is necessary for status go (34:32)
## Jonathan
Wait for mobile team to remove ID verification code before cleaning up status go (49:25)
## All team members
	Review and provide feedback on Sam's  qshared call notes and analysis document (50:00)
