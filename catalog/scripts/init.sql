CREATE TABLE products (
id SERIAL NOT NULL PRIMARY KEY,
title VARCHAR(100) NOT NULL,
category VARCHAR(50) NOT NULL,
description VARCHAR(200) NOT NULL,
price NUMERIC(6,2)
);

INSERT INTO products(title, category, description, price) VALUES("Test Product", "Test", "This is dumb test product", 300.50);