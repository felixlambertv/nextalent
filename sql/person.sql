/**
 * Create Person Table
 */
CREATE TABLE people (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Country VARCHAR(255) NOT NULL
);

/**
 * Seed person table
 */
INSERT INTO people (Name, Country) VALUES
('Adam', 'Kuala Lumpur'),
('John', 'Singapore'),
('Henry', 'Singapore'),
('Dominic', 'Thailand');