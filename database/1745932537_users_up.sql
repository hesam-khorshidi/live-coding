BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    phone_number varchar(30)
);

COMMIT TRANSACTION;