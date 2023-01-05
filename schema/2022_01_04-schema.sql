create table if not exists erajaya_product (
    id VARCHAR(255) not null primary key,
    name VARCHAR(255),
    price INTEGER,
    description VARCHAR,
    quantity INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
)