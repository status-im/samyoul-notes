# Overview

The Weekly Sync Meeting for the status-go guild focused on several key areas including a discussion on Git workflows where the team agreed to maintain a linear history with merge commits and to adjust GitHub settings to enforce rebasing before merging.

Team members provided updates on their ongoing projects:
- Andrey is developing a new RPC provider
- Dario is enhancing functional tests and caching for blockchain data
- Patryk is improving logging systems.
The group discussed unifying logging libraries and the potential for a wrapper logger interface.

Sam introduced a new requirements-first process involving developers, QA, and designers to enhance communication and feasibility assessment. The meeting also included discussions on implementation timelines and storing requirements in GitHub or Notion.

Lastly, action items were reviewed, including updating policies to align with the `contributing.md` file, with no next meeting scheduled.

# Notes

## 🔄 Git Workflow Discussion (00:01 - 07:47)
- Discussed merge and squash vs. full commit history
- Agreed on keeping linear history with merge commits
- GitHub settings to be adjusted to enforce rebasing before merging
- Aim to avoid merging branches into feature branches

## 📊 Team Updates (08:16 - 12:40)
- Andrey: Added new RPC provider for node flip, working on market go
- Dario: Worked on functional tests, caching for blockchain data
- Patryk: Working on logging improvements

## 🖨️ Logging System Discussion (12:40 - 23:54)
- Aim to unify logging libraries across the application
- Considering custom solution for categorising and filtering logs
- Discussed pros and cons of creating a wrapper logger interface

## 📝 Requirements-First Process (23:56 - 40:50)
- Sam: Working on new requirements-first process
- Process involves developer, QA, and designer from ideation to maintenance
- Aims to improve communication and technical feasibility assessment
- User stories to be minimum requirement, with additional documents as needed

## 🔍 Process Implementation Details (40:50 - 53:38)
- Discussed alignment with Arwen's high-level strategy
- Timeline for introducing the new process to teams
- Considered storing requirements in GitHub or public Notion
- Discussed potential impact on status-go guild

## 📅 Meeting Wrap-up (53:56 - 57:05)
- Reviewed action points from previous meeting
- Discussed updating policies to be consistent with contributing.md
- Confirmed no next meeting scheduled

# Action items
Samuel Hawksby-Robinson
- Update policies to be consistent with `contributing.md` file (54:13)