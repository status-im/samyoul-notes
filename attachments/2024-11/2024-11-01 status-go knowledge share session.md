Overview

The "status-go knowledge share session" focused on the historical context and technical challenges surrounding the Status Go project, emphasizing its evolution from initial integrations of mobile storage and messaging to the development of the Waku protocol. Discussions revealed that while end-to-end encryption and user privacy are prioritized, they complicate reliability and synchronization within the application. The team considered the potential transition from Go to Nim for development, weighing the pros and cons of such a significant rewrite against the need to maintain functionality and address technical debt. Product development priorities were established, highlighting the importance of documentation, knowledge sharing, and gradual refactoring of the Status Go code. Collaboration with research teams from Waku and Nimbus was encouraged, and the meeting concluded with actionable steps, including scheduling a follow-up meeting at Devcon and enhancing current documentation.

Notes

Historical Context and Project Origins (05:04 - 23:38)

- Status initially aimed to integrate storage, messaging, and execution on a mobile node
    
- Ethereum Foundation dropped Whisper and Swarm, which Status took on
    
- Nimbus was created to have a seat at the table in developing consensus protocol
    
- Waku was developed as a replacement for Whisper, offering better scalability
    
- Status Go development focused on end-user use cases, not fundamental chat application aspects
    
- Reliability became a neglected aspect between Waku protocol and Status app development
    

️ Technical Challenges and Architecture (23:38 - 53:40)

- End-to-end encryption and reliability must happen inside the encrypted stream
    
- Status Go handles end-to-end encryption as it owns the user's keys
    
- Waku aims for sender and receiver anonymity, complicating reliability implementation
    
- Communities (group chats) introduced additional complexities in synchronization and key sharing
    
- Status Go code became complex, mixing various functionalities in long functions
    
- Proposal to integrate Enwaku into Status Go as a step towards better architecture
    

Integration and Language Considerations (53:41 - 01:22:41)

- Discussion on replacing Go with Nim for Status Go
    
    - Discussion centered around the potential benefits of replacing Go with Nim for Status Go, emphasizing improved integration and efficiency.
        
    - Concerns were raised about the complexity and historical challenges of the current Status Go codebase, highlighting the need for better documentation and understanding of the existing protocol.
        
    - Participants agreed on the importance of maintaining a clear separation of concerns within the code to facilitate future development and potential rewrites.
        
    
- Concerns raised about the feasibility and necessity of rewriting Status Go
    
- Emphasis on documenting existing protocol and creating specifications
    
- Suggestion to refactor Status Go gradually, extracting and isolating components
    
- Consideration of language interoperability, especially Go's challenges in integration
    

Product Development and Priorities (01:22:41 - 01:54:21)

- Focus on building a sustainable product and maintaining V1 features
    
- Need to identify target users and address existing bugs
    
- Rewriting Status Go considered unrealistic and potentially harmful to the project
    
- Discussion on the importance of documentation and knowledge sharing
    
- Balancing technical debt reduction with product development pressures
    

Collaboration and Future Steps (01:54:21 - 02:20:49)

- Encouragement to leverage research from Waku and Nimbus teams
    
- Suggestion to think about protocol specification and modularity during development
    
- Proposal for follow-up discussions at Devcon
    
- Emphasis on the support available from research teams for integration efforts
    
- Reminder of the long-term benefits of improving code structure and documentation
    

Action items

Igor

- Schedule a follow-up meeting at Devcon to discuss more concrete details of Status Go integration with Enwaku (02:19:09)
    

unassigned

- Review and update Status Go documentation to ensure it's current and comprehensive (01:57:30)
    

Status Go team

- Begin process of extracting and isolating components within Status Go for easier future modifications (01:52:22)
    

Wallet team

- Explore possibilities of integrating Nimbus' verifying proxy for improved security in wallet transactions (01:03:08)