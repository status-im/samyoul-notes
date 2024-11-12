# TLDR

We should consider deleting Feature Upvote and transitioning to a Discourse plugin for our feature suggestion and voting needs.

# What Am I Talking About?

https://status.app/feature-upvote

Simply put, Feature Upvote is a feature suggestion and voting functionality accessible through the [status.app](http://status.app/) website. It allows users to submit ideas and vote on feature requests, providing valuable insight into what our community would like to see developed next. However, the tool has some notable downsides that impact our independence, security, and costs, prompting a need for discussion about its future.

# Why Is It a Problem?

There are several reasons why Feature Upvote is no longer suitable for our needs. Let’s break them down.

## 1. Aligning with Status’ Decentralisation and Open Source Principles

One of Status’ core principles is openness and decentralisation. By removing reliance on an external SaaS product, we are taking an active step toward reducing central points of control and dependency. This shift would serve as a statement that Status prioritises community empowerment and open source solutions, which will resonate a lot better with our user base.

## 2. Feature Upvote Is a 3rd Party Service

Feature Upvote is a separate service not owned or maintained by Status. While it allows us to quickly set up and manage feature requests, it relies on external hosting and maintenance by https://featureupvote.com/.

By depending on a 3rd party service, we introduce a potential risk factor, as any changes to their policies or availability could impact our ability to use the service. Additionally, it introduces data privacy concerns since we do not fully control the environment where our users’ data is stored and accessed.

## 3. Feature Upvote Is a Closed Source SaaS Solution

This tool is a closed-source SaaS product, which presents limitations for us as an organisation that values transparency and decentralisation. We cannot self-host it, meaning we have no control over the underlying database or its data handling processes. As part of our website integration, we’re embedding Feature Upvote via an `<iframe>` from an external domain.

Take a look at the example below:

- The `<iframe>` points to https://status-desktop.featureupvote.com/, a domain fully managed by Feature Upvote.

![image|690x289](upload://x2OE8k1lbQ89xfwKvgFN6O0Izpj.jpeg)

The `<iframe>` points to the following url: https://status-desktop.featureupvote.com/

![image|690x332](upload://cMXGkKv148bPjWIcF9BXPoXZJl9.jpeg)

This is a problem because it is inconsistent with our vision of an open, community-driven project. Our reliance on this third-party platform is a step away from transparency and autonomy, which we strive to prioritise.

## 4. Annual Cost of $1600+ for Feature Upvote

Please see our Application Catalogue for details of our declared costs related to Feature Upvote.

[Feature UpVote](https://www.notion.so/Feature-UpVote-2bfa1c1fd089430b92497271e5685dc7?pvs=21) 

The cost of using Feature Upvote is substantial, adding up to over $1600 annually.

While this may be a manageable expense, it’s worth questioning whether we could allocate these funds more effectively, especially when considering alternative open-source or in-house options that would provide us with more control and flexibility.

Here’s a snapshot of their pricing model:

- https://featureupvote.com/pricing/

![image|601x490, 50%](upload://iRGXqLeAJhmoKa3EwhgC3p4qvmE.png)![image|690x435, 50%](upload://rRlEEzE43KFm9WJ34mP1jU3NlQh.jpeg)

Considering that this cost is recurring and does not directly align with our decentralised principles, it is essential to weigh whether we should continue this expenditure or look for a more cost-effective, in-house solution.

## 5. Data Privacy and Security Concerns

Since Feature Upvote is a closed-source SaaS, there’s limited visibility into how data is handled. By moving to an in-house, open-source solution like Discourse, we can ensure greater privacy controls, giving us full access to log data, analytics, and potential security patches. This is especially relevant in light of data protection regulations and users’ increasing demand for transparency about how their data is used.

## 6. Integration and Customisation with our Existing Workflows

Feature Upvote, as a separate tool, requires users to adapt to a different interface and set of features that may not fully match the rest of our ecosystem. With a Discourse plugin, we can customise the voting and suggestion process to better align with our community’s workflow, providing a more cohesive and integrated experience. For instance:

- Voting could tie into Discourse’s existing notification and ranking features.
- Feature discussions can occur directly within the platform, making it easier for community members to track updates, comments, and related discussions without switching between platforms.

An in-house solution, particularly with open-source tools like Discourse, allows us to adapt and scale the feature suggestion process over time. If new requirements arise (such as tagging features by priority or integrating them into development pipelines), we can make these customisations directly or through plugins

This flexibility would not be available with Feature Upvote, where we are limited to their feature set and update cycle.

## 7. Enhanced Moderation and Community Standards

Feature voting systems often need strong moderation to ensure the suggestions align with community standards and project goals.

With Discourse, we would have direct control over moderation features and can use existing tools to enforce guidelines, merge duplicate suggestions, and prevent misuse. Moderators would also benefit from centralised tools, rather than having to manage two separate platforms.

## 8. Encouraging a More Active Community

Since Discourse is already our primary community platform, having a feature suggestion and voting functionality within the same space could boost engagement.

Users may be more likely to participate if they are already familiar with the interface and don’t have to go to a separate tool. It could also increase engagement in other areas of the forum, as users browse, comment, and discuss other threads while voting on features.

See: https://forum.vac.dev/ and https://discuss.status.app/

# Proposed Solution: Transition to a Discourse Plugin

Discourse, the forum software we already use for our community discussions, has a number of plugins available that enable feature suggestions and voting capabilities similar to what Feature Upvote offers. By transitioning to a Discourse-based system, we can leverage our existing infrastructure to consolidate feature requests, streamline user interaction, and cut down on costs.

You can see the old feature requests here:
https://discuss.status.app/c/archive/features/51

Even without an upvote functionality many open source projects rely on Discourse to provide a discussion space for their community to discuss feature requests:

- [Nextcloud feature requests](https://help.nextcloud.com/tag/featurerequest)
- [Docker Feature Requests](https://forums.docker.com/c/archive/feature-requests/5)
- [Discourse Feature Discussion](https://meta.discourse.org/c/feature/2)
- [Mozilla Feature Requests](https://discourse.mozilla.org/tag/feature-request)
- [OpenWrt Feature Requests](https://forum.openwrt.org/c/feature-requests/16)
- [GitLab Feature Discussion](https://forum.gitlab.com/tag/feature)
- [OpenStreetMap Feature Requests](https://community.openstreetmap.org/tag/feature-request)

Other open source projects use Discourse plugins or a custom solution for their self hosted feature management:

https://community.anytype.io/c/feature-requests/6/none
https://community.home-assistant.io/c/feature-requests/13

### Benefits of a Discourse Plugin Solution:

- **Control and Ownership**: By hosting everything on our own infrastructure, we would retain full control over the data and the platform’s functionality, ensuring it aligns with Status’ commitment to transparency.
- **Cost Efficiency**: Eliminating the recurring $1600+ annual fee would allow us to allocate resources more effectively, potentially funding further development on other important community-driven initiatives.
- **Streamlined User Experience**: By keeping feature requests within Discourse, we create a seamless experience for users, eliminating the need to navigate to a separate site or engage with an embedded iframe.
- **Community Moderation and Engagement**: Discourse’s moderation tools and community-driven ethos provide a platform where users can engage more deeply in discussions about feature ideas, fostering a richer dialogue.

# What Next?

Feature Upvote served us well as a quick solution for collecting community feedback on potential features. However, due to its third-party nature, closed-source model, and associated costs, it does not align with our priorities or principles. Transitioning to a Discourse plugin offers a more transparent, cost-effective, and community-friendly approach that fits well within our mission.

Let’s discuss this proposal further and explore the specific Discourse plugin options that might meet our needs, ensuring a smooth transition for both our community and team.