-- name: ListProducts :many
SELECT * FROM product.products;

-- name: GetProductsByNames :many
SELECT * FROM product.products WHERE name = ANY($1::text[]);

-- name: GetProductByID :one
SELECT * FROM product.products WHERE id = ($1::uuid);
