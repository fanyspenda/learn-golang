-- create database
CREATE DATABASE alta_online_shop;

-- membuat table users
CREATE TABLE users (id int PRIMARY KEY NOT NULL AUTO_INCREMENT, status smallint, gender char(1), created_at timestamp, updated_at timestamp);

-- membuat table product_types
CREATE TABLE product_types (id int PRIMARY KEY NOT NULL AUTO_INCREMENT, name varchar(255),  created_at timestamp, updated_at timestamp);

-- membuat table operators
CREATE TABLE operators (id int PRIMARY KEY NOT NULL AUTO_INCREMENT, name varchar(255),  created_at timestamp, updated_at timestamp);

-- membuat table products
CREATE TABLE products (
    id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    product_type_id int,
    operator_id int,
    name varchar(100),
    code varchar(50),
    status smallint,
    created_at timestamp, 
    updated_at timestamp,
    FOREIGN KEY (product_type_id) REFERENCES product_types(id),
    FOREIGN KEY (operator_id) REFERENCES operators(id) 
);

-- membuat table product_descriptions
CREATE TABLE product_descriptions (
    id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    description text,
    created_at timestamp, 
    updated_at timestamp
);

-- membuat table payment_methods
CREATE TABLE payment_methods (
    id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name varchar(255),
    status smallint,
    created_at timestamp, 
    updated_at timestamp
);

-- menambah kolom dob ke table user
ALTER TABLE users ADD dob date;

-- membuat table transactions
CREATE TABLE transactions (
    id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    user_id int,
    payment_method_id int,
    status varchar(10),
    total_qty int,
    total_price numeric(25,2),
    created_at timestamp, 
    updated_at timestamp,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id)
);

-- membuat table transaction_details
CREATE TABLE transaction_details (
    transaction_id int,
    product_id int,
    status varchar(10),
    qty int,
    price numeric(25,2),
    created_at timestamp, 
    updated_at timestamp,
    PRIMARY KEY (transaction_id, product_id),
    FOREIGN KEY (transaction_id) REFERENCES transactions(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

-- menambah kolom di table products
ALTER TABLE products ADD product_description_id text;
ALTER TABLE products MODIFY product_description_id int;
ALTER TABLE products ADD FOREIGN KEY (product_description_id) REFERENCES product_descriptions(id);

-- membuat table kurir
CREATE TABLE kurir (
    id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name varchar(255),
    created_at timestamp, 
    updated_at timestamp
);

-- tambah kolom ongkos_dasar
ALTER TABLE kurir
ADD ongkos_dasar numeric(25, 2);


-- mengubah nama table
ALTER TABLE kurir RENAME TO shipping;

-- hapus table shipping
DROP TABLE shipping;


-- PROBLEM 2

-- menambah 5 operator
INSERT INTO operators (
    name
) VALUES 
("Tinky Winky"),
("Dipsy"),
("Lala"),
("Po"),
("John Doe");

-- menambah 3 product product type
INSERT INTO product_types (
    name
) VALUES 
("type 1"),
("type 2"),
("type 3");

-- menambah 2 product dengan product_type_id 1 dan operator_id 3
INSERT INTO products (
    product_type_id,
    operator_id,
    name
) VALUES 
( 1, 3, "Gunpla 1"),
( 1, 3, "Gunpla 2");

-- menambah 2 product dengan product_type_id 2 dan operator_id 1
INSERT INTO products (
    product_type_id,
    operator_id,
    name
) VALUES 
( 2, 1, "Yugioh Cyber Dragon pack 1"),
( 2, 1, "Yugioh Cyber Fortune Lady pack 1"),
( 2, 1, "Yugioh Exodia pack 1");

-- menambah 2 product dengan product_type_id 3 dan operator_id 4
INSERT INTO products (
    product_type_id,
    operator_id,
    name
) VALUES 
( 3, 4, "Action Figure Agumon"),
( 3, 4, "Action Figure Metal-Greymon"),
( 3, 4, "Action Figure Wargreymon");

-- membuat product description
INSERT INTO product_descriptions (
    description
) VALUES
("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed eget ante lacinia, pretium lorem et, viverra orci. Nunc non augue at felis efficitur facilisis. Vestibulum scelerisque, risus quis venenatis lacinia, mi justo interdum mi, ut varius justo nulla sed felis. Quisque at velit ac felis ullamcorper pretium. Fusce interdum risus id est pretium fringilla nec sit amet dolor. Vivamus congue, risus in sodales pulvinar, ex nisl tincidunt leo, nec viverra arcu purus scelerisque urna. Cras interdum ipsum at nisl blandit laoreet. Vivamus id enim tincidunt, dignissim nisl pretium, elementum urna. Sed facilisis, ex eget euismod ullamcorper, ligula ex interdum turpis, ac consectetur nisi orci id lacus. Nulla auctor blandit risus, in tincidunt massa sollicitudin eget. In tincidunt ex nec felis gravida, id vestibulum orci varius. Nulla lacinia a lorem quis dapibus.");

-- menambah deskripsi dari products
UPDATE products SET product_description_id=1;

-- menambah 3 payment methods
INSERT INTO payment_methods (
    name, status
) VALUES 
("all", 1),
("cash", 2),
("credit", 3);

-- menambah kolom name pada table user
ALTER TABLE users ADD name VARCHAR(200);

-- menambahkan data user
INSERT INTO users (
    name,
    status,
    gender,
    dob
) VALUES 
("Coach", 2, "M", "1976-10-18"),
("Rochelle", 1, "F", "1978-12-25"),
("Nick", 3, "M", "1982-03-13"),
("Ellis", 4, "M", "1987-06-07"),
("Zoey", 2, "F", "1989-7-23");

-- menambahkan 3 transaksi tiap user dan masing-masing transaksi 3 product
INSERT INTO transactions (
    user_id,
    payment_method_id,
    status,
    total_qty
) VALUES
(5, 1, 1, 10),
(5, 2, 2, 20),
(5, 3, 3, 30),
(6, 1, 1, 15),
(6, 2, 2, 25),
(6, 3, 3, 35),
(7, 1, 1, 17),
(7, 2, 2, 27),
(7, 3, 3, 37),
(8, 1, 1, 18),
(8, 2, 2, 28),
(8, 3, 3, 38),
(9, 1, 1, 13),
(9, 2, 2, 23),
(9, 3, 3, 33);

INSERT INTO transaction_details (
    transaction_id,
    product_id
) VALUES 
(16, 1),
(16, 2),
(16, 3),
(17, 4),
(17, 5),
(17, 6),
(18, 7),
(18, 8),
(18, 9),
(19, 1),
(19, 2),
(19, 3),
(20, 4),
(20, 5),
(20, 6),
(21, 7),
(21, 8),
(21, 9),
(22, 1),
(22, 2),
(22, 3),
(23, 4),
(23, 5),
(23, 6),
(24, 7),
(24, 8),
(24, 9),
(25, 1),
(25, 2),
(25, 3),
(26, 4),
(26, 5),
(26, 6),
(27, 7),
(27, 8),
(27, 9),
(28, 1),
(28, 2),
(28, 3),
(29, 4),
(29, 5),
(29, 6),
(30, 7),
(30, 8),
(30, 9);

-- menampilkan user dengan gender "M"
SELECT * FROM users WHERE gender = "M";

-- menampilkan product dengan id = 3
SELECT * FROM products WHERE id = 3;

-- menampilkan users yang dari 7 hari yang lalu hingga sekarang dan nama mengandung kata 'a'
SELECT * FROM users WHERE created_at BETWEEN ADDDATE(now(), -7) AND now() AND name LIKE ("%a%");

-- menghitung jumlah user perempuan
SELECT COUNT(*) FROM users WHERE gender = "F";

-- menampilkan users dengan nama berurutan
SELECT * FROM users ORDER BY name ASC;

-- menampilkan hanya 5 data products
SELECT * FROM products LIMIT 5;

UPDATE products SET name ="product dummy" WHERE id = 1;

UPDATE transaction_details SET qty=3 WHERE product_id=1;

-- delete product with id=1
DELETE FROM transaction_details WHERE product_id=1;
DELETE FROM products WHERE id=1;



-- PROBLEM 3
-- menampilkan transactions dengan user_id = 1 dan 2
SELECT * FROM transactions WHERE user_id = 5 OR user_id = 6;

-- memberi value total_price
UPDATE transactions SET total_price=3000 WHERE user_id=5;

-- jumlah harga transaksi user_id=1
SELECT SUM(total_price) FROM transactions WHERE user_id=5;


SELECT SUM(total_price) FROM transactions WHERE product_type=5;

SELECT COUNT(*) FROM transactions;

-- count transactions with product_type = 2
SELECT COUNT(*) FROM (SELECT transactions.id FROM transactions 
INNER JOIN transaction_details ON transactions.id=transaction_details.transaction_id
INNER JOIN products ON products.id = transaction_details.product_id WHERE products.product_type_id=2
GROUP BY transactions.id) AS count_transactions;

-- menampilkan product dan product_type.name yang saling berhubungan
SELECT products.*, product_types.name as type_name FROM products INNER JOIN product_types ON products.product_type_id=product_types.id;

-- Tampilkan semua field table transaction, field name table product dan field name table user.
SELECT users.name, transactions.*, products.name FROM transactions 
INNER JOIN transaction_details ON transaction_details.transaction_id = transactions.id
INNER JOIN products ON transaction_details.product_id = products.id
INNER JOIN users ON transactions.user_id = users.id;

-- buat trigger ketika hapus data transaction, harus hapus data transaction details dulu
DELIMITER $$
CREATE TRIGGER delete_transaction_details
BEFORE DELETE ON transactions FOR EACH ROW
BEGIN
 DECLARE v_transaction_id INT;
 SET v_transaction_id = OLD.id;

 DELETE FROM transaction_details WHERE transaction_id=v_transaction_id;
END$$


-- testing TRIGGER
INSERT INTO transactions (
    user_id,
    payment_method_id
) VALUES
(9, 3),
(8, 3);

INSERT INTO transaction_details (
    transaction_id,
    product_id
) VALUES 
(31, 2),
(32, 2);
--------------

INSERT INTO transaction_details (
    transaction_id,
    product_id,
    qty
) VALUES 
(32, 3, 3),
(32, 4, 5),
(32, 5, 7),
(32, 6, 9);

-- update transactions table qty when transaction_detail got deleted
DELIMITER $$
CREATE TRIGGER update_transactions_qty
AFTER DELETE ON transaction_details FOR EACH ROW
BEGIN
 DECLARE v_total_qty INT;
 SELECT SUM(transaction_details.qty) INTO v_total_qty FROM transaction_details WHERE transaction_details.transaction_id = OLD.transaction_id;

 UPDATE transactions SET total_qty = v_total_qty WHERE id = OLD.transaction_id;
END$$

INSERT INTO products (
    id, 
    product_type_id,
    operator_id,
    name
    ) VALUES 
    (1, 2, 3, "Gundam Sazabi" );

SELECT products.id, products.name FROM (SELECT products.id, products.name FROM products
INNER JOIN transaction_details ON transaction_details.product_id=products.id 
GROUP BY products.id
ORDER BY products.id ASC) as used_products
RIGHT JOIN products ON used_products.id = products.id
WHERE used_products.id IS NULL;