CREATE TABLE foo (bar bool not null);

-- name: ListFoo :iter
SELECT bar FROM foo;

CREATE TABLE baz (one integer, two integer);

-- name: ListBaz :iter
SELECT one, two FROM baz;
