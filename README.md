Enterprise Immutable Ledger Platform

Project Overview
This is a production grade distributed financial ledger built with Go. The system is designed for high integrity transaction processing and asynchronous fraud detection. It implements principles often found in blockchain systems but optimized for the low latency needs of private banking environments.

Core Architecture Features
1. Strict Transaction Integrity: Utilizes PostgreSQL in SERIALIZABLE isolation mode to prevent race conditions, double spending, and write skew anomalies.
2. Cryptographic Audit Trail: Implements a Hash Linked Ledger where each transaction is cryptographically tied to the previous entry to ensure the history is tamper proof.
3. Idempotency Engine: Uses a Redis backed middleware to ensure that network retries do not result in duplicate financial operations.
4. Event Driven Analysis: Decouples the primary transaction flow from the fraud detection logic using a Pub/Sub worker pattern.

Technical Stack
- Language: Go (Golang) 1.22
- Primary Database: PostgreSQL (Serializable Isolation)
- Caching and Messaging: Redis 7
- Orchestration: Docker and Docker Compose
- Performance Testing: k6

Setup and Execution
1. Clone the repository and navigate to the root directory.
2. Run the build script to compile binaries: sh build.sh
3. Launch the full infrastructure stack: docker-compose up
4. Execute the load test to verify consistency: k6 run load_test.js

Security and Reliability Design
The system follows the Clean Architecture pattern to separate business logic from infrastructure. Security is enforced at multiple layers:
- Application Layer: Idempotency keys prevent duplicate processing at the API gateway.
- Database Layer: Append only triggers prevent any modification of historical data.
- Infrastructure Layer: Services run as non privileged users within isolated containers.
