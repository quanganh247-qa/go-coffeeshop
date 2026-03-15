
-- INSERT INTO products (name, type, price, image) VALUES
-- ('CAPPUCCINO', 0, 4.5, 'img/CAPPUCCINO.png'),
-- ('COFFEE_BLACK', 1, 3.0, 'img/COFFEE_BLACK.png'),
-- ('COFFEE_WITH_ROOM', 2, 3.0, 'img/COFFEE_WITH_ROOM.png'),
-- ('ESPRESSO', 3, 3.5, 'img/ESPRESSO.png'),
-- ('ESPRESSO_DOUBLE', 4, 4.5, 'img/ESPRESSO_DOUBLE.png'),
-- ('LATTE', 5, 4.5, 'img/LATTE.png'),
-- ('CAKEPOP', 6, 2.5, 'img/CAKEPOP.png'),
-- ('CROISSANT', 7, 3.25, 'img/CROISSANT.png'),
-- ('MUFFIN', 8, 3.0, 'img/MUFFIN.png'),
-- ('CROISSANT_CHOCOLATE', 9, 3.5, 'img/CROISSANT_CHOCOLATE.png');

START TRANSACTION;

CREATE SCHEMA IF NOT EXISTS "product";

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE
    product.products (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        name VARCHAR(255) NOT NULL,
        type INTEGER NOT NULL UNIQUE,
        price DOUBLE PRECISION NOT NULL,
        stock INTEGER NOT NULL,
        image VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE UNIQUE INDEX ix_product_id ON product.products (id);

COMMIT;