# Overview
The Core Teams Sync (Desktop + Mobile) meeting focused on critical aspects affecting both platforms, beginning with a release coordination update highlighting the need for version compatibility amongst releases 229, 230, and 231, with the mobile team targeting an end-of-October release. The session addressed improvements in the bug reporting process, emphasizing the challenges of user privacy when collecting logged information, and explored the potential use of Sentry or Google Crash Reporting. The group discussed the establishment of a central repository for legal documents, stressing the importance of versioning and handling diverging features across platforms, while also raising concerns regarding the accuracy of app store information and the need for comprehensive privacy policies. Additionally, issues with community description synchronization were identified, particularly regarding the ‘last opened at’ property and its bandwidth implications. A discussion on state management revealed the misuse of Waku, prompting consideration of alternative solutions and a plan for implementing necessary changes. Action items were assigned to various members for follow-up on compatibility checks, error logging integration, privacy policy inquiries, synchronization removal, and community description contributions.
# Notes
## 🚀 Release Coordination (00:04 - 06:32)
- Discussion on compatibility between versions 229, 230, and 231
- Pairing fallback needs to be on the same version for both mobile and desktop
- Mobile team considering release by the end of October
- Wallet team expected to be ready by mid-October

## 🐛 Bug Reporting and User Experience (07:20 - 17:41)
- Discussion on improving bug reporting process
- Challenges with requesting logs from users due to privacy concerns
- Consideration of using Sentry or Google Crash Reporting for error logging
- Need to be cautious about potentially sensitive data in logs

## 📄 Legal Documents Repository (17:45 - 28:16)
- Sam introduced a central repository for legal documents
- Discussion on versioning and platform-specific policies
- Consideration of how to handle diverging features across platforms

## 🔐 Privacy Policy and App Store Information (28:16 - 36:14)
- Icaro raised concerns about accuracy of app store information
- Need for legal team to review and approve app store descriptions
- Discussion on potential need for platform-specific privacy policies

## 🔄 Community Description Synchronization (36:31 - 44:12)
- Patrick identified issue with community description propagation
- Discussion on the 'last opened at' property and its impact on bandwidth
- Consideration of removing or modifying the synchronization process

## 🗄️ State Management and Waku Usage (44:12 - 56:27)
- Discussion on misuse of Waku for state management
- Consideration of alternative solutions like Codex or blockchain
- Need for a strategy to implement breaking changes over time
- Invitation to contribute to ongoing discussion thread on community description
