Overview

In the recent "status-go: goals sync" meeting, the team discussed several key initiatives and action items.

Igor introduced specific GitHub labels for differentiating desktop, status-go, and mobile goals, addressing concerns from Samuel regarding their complexity.

The group also reviewed logger improvements led by Patryk, who is working on a PR to enhance logging practices, while Samuel raised concerns about flexibility in handling loggers.

Additionally, Andrey provided updates on the rework of the RPC provider storage and emphasized the need for discussions on merge commit policies and integration with go-ethereum.

Igor is tasked with finalizing the Code Cov integration, while Patryk and Andrey will continue their respective developments on logging and RPC provider enhancements.

Dario has been assigned to explore the integration of network-dependent unit tests into the CI process in collaboration with Jakub.

Notes

️ GitHub Labels for Goals (00:00 - 11:58)

- Igor added GitHub labels for goals as agreed
    
- Labels were renamed to be more specific
    
- Labels distinguish between desktop, status-go, and mobile goals
    
- Samuel expressed concerns about the complexity of the labels
    
- Discussion on filtering and prioritization of issues using labels
    
- Dario mentioned a small update on a PR related to network-dependent tests
    

Logger Improvements (12:37 - 20:47)

- Patryk opened a PR to replace get logger usage with sublogger
    
- Plan to substitute the root engine and reverse the direction of logger proxying
    
- Goal to implement categorized logging after cleanup
    
- Samuel raised concerns about passing loggers around by interface
    
- Discussion on the flexibility of Zap logger and named instances
    

RPC Provider Rework and Merge Commits (20:50 - 28:07)

- Andrey working on reworking RPC provider storage in the network
    
- Removed upstream provider from RPC client
    
- Dario mentioned the need to discuss go-ethereum integration
    
- Andrey raised the importance of discussing merge commits policy
    
- Igor updating on Code Cov integration progress
    

Action items

Igor Sirotin

- Finalize Code Cov integration and configure it to fit team needs (27:09)
    

Patryk Osmaczko

- Continue work on replacing get logger usage with sublogger (12:42)
    

Andrey Bocharnikov

- Continue reworking RPC provider storage in the network (20:57)
    

Dario Lipicar

- Explore integration of network-dependent unit tests into CI with Jakub (11:39)