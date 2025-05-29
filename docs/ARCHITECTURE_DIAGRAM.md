```mermaid
graph TD
    subgraph Client
        C1[Web Client]
        C2[Mobile App]
    end

    subgraph API Layer
        A1[/api/transfers]
        A2[/api/accounts/balance]
        A3[/api/webhook/transfer]
    end

    subgraph Service Layer
        S1[TransferService]
        S2[AccountService]
        S3[NotificationService]
    end

    subgraph Repository Layer
        R1[AccountRepository]
        R2[TransferRepository]
        R3[NotificationRepository]
    end

    subgraph Database
        DB1[accounts]
        DB2[transfers]
        DB3[notifications]
    end

    C1 --> A1
    C2 --> A1
    A1 --> S1
    A2 --> S2
    A3 --> S3
    S1 --> R1
    S1 --> R2
    S2 --> R1
    S3 --> R3
    R1 --> DB1
    R2 --> DB2
    R3 --> DB3

    style C1 fill:#f9f,stroke:#333,stroke-width:2px
    style C2 fill:#f9f,stroke:#333,stroke-width:2px
    style A1 fill:#bbf,stroke:#333,stroke-width:2px
    style A2 fill:#bbf,stroke:#333,stroke-width:2px
    style A3 fill:#bbf,stroke:#333,stroke-width:2px
    style S1 fill:#bfb,stroke:#333,stroke-width:2px
    style S2 fill:#bfb,stroke:#333,stroke-width:2px
    style S3 fill:#bfb,stroke:#333,stroke-width:2px
    style R1 fill:#fbb,stroke:#333,stroke-width:2px
    style R2 fill:#fbb,stroke:#333,stroke-width:2px
    style R3 fill:#fbb,stroke:#333,stroke-width:2px
    style DB1 fill:#bfb,stroke:#333,stroke-width:2px
    style DB2 fill:#bfb,stroke:#333,stroke-width:2px
    style DB3 fill:#bfb,stroke:#333,stroke-width:2px
```
