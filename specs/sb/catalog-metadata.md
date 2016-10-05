---
title: Catalog Metadata
owner: Core Services
---

Not sure we have or should keep the concept of "CLI strings".

<p class="note"><strong>Note</strong>: "CLI strings" are all lowercase, no spaces.
Keep it short; imagine someone having to type it as an argument for a longer CLI
command.</p>

costs object should be standardized based on our discussions.

An array-of-objects that describes the costs of a service, in what
currency, and the unit of measure. If there are multiple costs, all of
them could be billed to the user (such as a monthly + usage costs at
once).  Each object must provide the following keys:
`amount: { usd:float }, unit: string`
This indicates the cost in USD of the service
plan, and how frequently the cost is occurred, such as "MONTHLY" or
"per 1000 messages".

