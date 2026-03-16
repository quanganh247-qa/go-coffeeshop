-- name: ListProducts :many
SELECT * FROM product.products;

-- name: GetProductsByNames :many
SELECT * FROM product.products WHERE name = ANY($1::text[]);

-- name: GetProductByID :one
SELECT * FROM product.products WHERE id = $1;

-- name: CreateProduct :one
INSERT INTO product.products (name, type, price, stock, image)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;
