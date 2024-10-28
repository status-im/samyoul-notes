Overview

The weekly sync for the status-go guild focused on several key areas, starting with a discussion led by Samuel on refining a draft of a zero policy document, emphasizing the need for clarity and actionable steps for implementation. The team decided to maintain the policy's focus on status-go for now and create a README file to clarify the policy process.

Igor reported progress on the Banix linter and highlighted potential integration of Sentry for crash reporting, while addressing concerns regarding data privacy.

The meeting also touched upon code coverage, with Igor tasked with finalizing integration and setting thresholds to manage flaky results.

Lastly, the team agreed to restructure their meeting schedule to bi-weekly sessions lasting one hour. Action items were assigned to Samuel and Igor to address various aspects of the discussions.

Notes

📋 Policy Discussion (00:01 - 10:40)
- Samuel presented a draft of a zero policy document
- Discussion on making the policy more specific and actionable
- Agreed to include steps for policy implementation in the repository
- Suggestion to use GitHub PRs for policy discussions instead of Notion

Policy Refinement (12:46 - 22:26)
- Discussed making the policy applicable to other Status repositories
- Agreed to keep the policy focused on status-go for now
- Decided to create a README file explaining the policy process
- Discussed the need for clear steps and conditions in the policy document

️ Banix Linter and Sentry Integration (22:41 - 30:16)

- Igor completed work on the Banix linter for status-go
- Discussed potential implementation of Sentry for crash reporting
- Concerns raised about data leakage and privacy when implementing Sentry

Privacy and Data Leakage Concerns (30:19 - 37:47)

- Discussed the importance of preventing accidental data leakage
- Acknowledged the need to balance privacy with useful error reporting
- Considered the implications of implementing new features on user privacy

Code Coverage and Integration (37:49 - 48:44)

- Igor working on finalizing code coverage integration
- Discussed issues with flaky coverage results and the need for thresholds
- Agreed to monitor and report any issues with code coverage before making it a requirement

️ Meeting Structure Discussion (48:44 - 55:32)

- Discussed reintroducing a general status-go meeting
- Debated frequency (weekly vs bi-weekly) and duration of meetings
- Decided on bi-weekly meetings with a one-hour duration

Action items

Samuel
- Refine the zero policy document based on feedback (02:10)
- Create a README file explaining the policy process (22:26)

Igor
- Continue monitoring and refining code coverage integration (38:05)
- Implement thresholds for code coverage reporting (38:05)
- Reorganize meetings to bi-weekly one-hour sessions (54:17)

All
- Keep monitoring and reporting issues with code coverage before it becomes a requirement (40:56)