\c general_purpose_server;

CREATE TABLE applications (
    id VARCHAR(26) PRIMARY KEY,
    secret_key UUID NOT NULL,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE representatives (
    id VARCHAR(26) PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    corporate_name VARCHAR(100) NOT NULL,
    mail VARCHAR(255) NOT NULL,
    phone_number VARCHAR(15) NOT NULL,
    address VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    application_id VARCHAR(26) REFERENCES applications(id)
);