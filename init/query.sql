CREATE TABLE translation (
    id SERIAL PRIMARY KEY, 
    word VARCHAR (255) NOT NULL,
    indonesian VARCHAR (255) NOT NULL,
    notes VARCHAR (255), 
    created_at TIMESTAMP with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);