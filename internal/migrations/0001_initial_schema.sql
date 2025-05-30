DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.tables 
        WHERE table_schema = 'public' 
        AND table_name = 'accounts'
    ) THEN
        CREATE TABLE accounts (
            id VARCHAR(255) PRIMARY KEY,
            balance DECIMAL(15, 2) NOT NULL DEFAULT 0,
            created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
        );
    END IF;
END $$;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.tables 
        WHERE table_schema = 'public' 
        AND table_name = 'transfers'
    ) THEN
        CREATE TABLE transfers (
            id VARCHAR(255) PRIMARY KEY,
            from_account VARCHAR(255),
            to_account VARCHAR(255),
            amount DECIMAL(15, 2) NOT NULL,
            status VARCHAR(20) NOT NULL,
            created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (from_account) REFERENCES accounts(id),
            FOREIGN KEY (to_account) REFERENCES accounts(id)
        );

        CREATE INDEX idx_transfers_from_account ON transfers(from_account);
        CREATE INDEX idx_transfers_to_account ON transfers(to_account);
        CREATE INDEX idx_transfers_status ON transfers(status);
        CREATE INDEX idx_transfers_created_at ON transfers(created_at);
    END IF;
END $$;
