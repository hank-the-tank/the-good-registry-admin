BEGIN;
CREATE SCHEMA IF NOT EXISTS schema;
COMMIT;
CREATE TABLE IF NOT EXISTS orders (
  id SERIAL PRIMARY KEY, 
  customer_name VARCHAR (255), 
  email VARCHAR(255),
  volume INT NOT NULL,
  created TIMESTAMP NOT NULL
);
COMMIT;
SELECT version();

-- INSERT INTO orders (id, customer_name, email, volume, created) VALUES (2, 'Hank', 'hank@sharesies.co.nz', 4, '2020-08-01 00:00:00.000000');
