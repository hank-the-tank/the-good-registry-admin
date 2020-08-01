BEGIN;
CREATE SCHEMA IF NOT EXISTS schema;
CREATE TABLE IF NOT EXISTS customers(
  id INT GENERATED ALWAYS AS IDENTITY, 
  first_name VARCHAR (255),
  last_name VARCHAR (255),
  preferred_name VARCHAR(255),
  email VARCHAR(255),
  phone VARCHAR(255),
  company VARCHAR(255),
  created TIMESTAMP NOT NULL,
  PRIMARY KEY(id)
);
CREATE TABLE IF NOT EXISTS orders (
  id INT GENERATED ALWAYS AS IDENTITY, 
  customer_id INT, 
  email VARCHAR(255),
  volume INT NOT NULL,
  created TIMESTAMP NOT NULL,
  PRIMARY KEY(id),
  CONSTRAINT fk_customer
    FOREIGN KEY(customer_id)
      REFERENCES customers(id)
);
CREATE TABLE IF NOT EXISTS gift_codes (
  id INT GENERATED ALWAYS AS IDENTITY,
  order_id INT,
  code VARCHAR(255),
  amount INT NOT NULL,
  created TIMESTAMP NOT NULL,
  PRIMARY KEY(id),
  CONSTRAINT fk_order
    FOREIGN KEY(order_id)
      REFERENCES orders(id)
);
COMMIT;
SELECT version();

-- INSERT INTO orders (id, customer_name, email, volume, created) VALUES (2, 'Hank', 'hank@sharesies.co.nz', 4, '2020-08-01 00:00:00.000000');
