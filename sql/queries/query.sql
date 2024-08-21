-- name: ListOrders :many
SELECT * FROM orders;

-- name: GetOrder :one
SELECT * FROM orders 
WHERE id = ?;

-- name: CreateOrder :exec
INSERT INTO orders (id, price, tax, final_price) 
VALUES (?, ?, ?, ?);

-- name: DeleteOrder :exec
DELETE FROM orders WHERE id = ?;