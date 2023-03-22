# **Overview**

The Fanfury-Sports chain relies on external off-chain data of matches and other sport-events. To push this data reliably to the chain, some kind of origin verification is required. The `DVM module` essentially fills this role in the Fanfury-Sports chain. The `DVM Module` verifies the following data:

- Event data pushed by the House to the chain
- Validity of Odds data using proposed ticket data during bet placement by user
