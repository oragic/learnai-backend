# Internal components
## Core
The core is where the heart of the application lives. It contains the essential business
logic and rules of the application.

### Ports
Think of ports as contracts or interfaces. They define how the application communicates with
external systems or services. For example, a port could specify the rules for connecting to a
database, interacting with another web services, or handling user interfaces

### Domain
Directory that contains domain models/entities representing core business concepts.

### Service 
Directory that contains the business logic or services of the application.

### Utils
Directory that contains utility functions that reused in the service package.

## Adapters
Adapters are the one who implement the contracts or interfaces defined by ports. Adapters are responsible for making sure the application can interact to databases, web services, or other things.