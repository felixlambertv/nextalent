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

/**
 * Select prodcedure function
 */
CREATE OR REPLACE FUNCTION get_country_by_name(p_name VARCHAR)
RETURNS VARCHAR AS $$

DECLARE result_country VARCHAR;
BEGIN
SELECT country INTO result_country
FROM people
WHERE name = p_name;
RETURN result_country;
END;

$$ LANGUAGE plpgsql;


 /**
  * Select from procedure
  */
SELECT get_country_by_name('Adam');