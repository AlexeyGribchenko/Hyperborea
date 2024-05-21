CREATE TABLE t_product (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description TEXT,
    shop_id INT,
    current_price DECIMAL,
    priority_value INT
);

CREATE TABLE t_user (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255),
    phone_number VARCHAR(18)
);

CREATE TABLE t_product_to_cart (
    id SERIAL PRIMARY KEY,
    product_id INT,
    user_id INT,
    add_date TIMESTAMP
);

CREATE TABLE t_product_to_favourites (
    id SERIAL PRIMARY KEY,
    product_id INT,
    user_id INT,
    add_date TIMESTAMP
);

CREATE TABLE t_category_to_product (
    id SERIAL PRIMARY KEY,
    product_id INT,
    category_id INT
);

CREATE TABLE t_category (
    id SERIAL PRIMARY KEY,
    parent_id INT,
    name VARCHAR(255)
);

CREATE TABLE t_product_characteristic (
    id SERIAL PRIMARY KEY,
    product_id INT,
    name VARCHAR(255),
    value TEXT
);

CREATE TABLE t_product_price (
    id SERIAL PRIMARY KEY,
    product_id INT,
    price DECIMAL,
    discount DECIMAL,
    creation_date TIMESTAMP
);

CREATE TABLE t_shop (
    id SERIAL PRIMARY KEY,
    user_id INT,
    title VARCHAR(255),
    description TEXT
);

CREATE TABLE t_image_to_product (
    id SERIAL PRIMARY KEY,
    product_id INT,
    image_id INT
);

CREATE TABLE t_image (
    id SERIAL PRIMARY KEY,
    path TEXT
);

CREATE TABLE t_image_to_comment (
    id SERIAL PRIMARY KEY,
    comment_id INT,
    image_id INT
);

CREATE TABLE t_comment (
    id SERIAL PRIMARY KEY,
    user_id INT,
    product_id INT,
    creation_date TIMESTAMP,
    rate INT,
    content TEXT
);

CREATE TABLE t_role (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50)
);

CREATE TABLE t_role_to_user (
    user_id INT,
    role_id INT
);

CREATE TABLE t_product_to_order (
    id SERIAL PRIMARY KEY,
    order_id INT,
    product_id INT,
    amount INT
);

CREATE TABLE t_order (
    id SERIAL PRIMARY KEY,
    address_id INT,
    total_price DECIMAL,
    creation_date TIMESTAMP,
    status_id INT,
    user_id INT
);

CREATE TABLE t_pick_up_address (
    id SERIAL PRIMARY KEY,
    city VARCHAR(100),
    street VARCHAR(255),
    house VARCHAR(20)
);

CREATE TABLE t_status (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50)
);

ALTER TABLE t_product
    ADD CONSTRAINT fk_shop_id
    FOREIGN KEY (shop_id)
    REFERENCES t_shop(id);

ALTER TABLE t_product_to_cart
    ADD CONSTRAINT fk_product_id
    FOREIGN KEY (product_id)
    REFERENCES t_product(id);

ALTER TABLE t_product_to_cart
    ADD CONSTRAINT fk_user_id
    FOREIGN KEY (user_id)
    REFERENCES t_user(id);

ALTER TABLE t_product_to_favourites
    ADD CONSTRAINT fk_product_id
    FOREIGN KEY (product_id)
    REFERENCES t_product(id);

ALTER TABLE t_product_to_favourites
    ADD CONSTRAINT fk_user_id
    FOREIGN KEY (user_id)
    REFERENCES t_user(id);

ALTER TABLE t_category_to_product
    ADD CONSTRAINT fk_product_id
    FOREIGN KEY (product_id)
    REFERENCES t_product(id);

ALTER TABLE t_category_to_product
    ADD CONSTRAINT fk_category_id
    FOREIGN KEY (category_id)
    REFERENCES t_category(id);

ALTER TABLE t_product_characteristic
    ADD CONSTRAINT fk_product_id
    FOREIGN KEY (product_id)
    REFERENCES t_product(id);

ALTER TABLE t_product_price
    ADD CONSTRAINT fk_product_id
    FOREIGN KEY (product_id)
    REFERENCES t_product(id);

ALTER TABLE t_shop
    ADD CONSTRAINT fk_user_id
    FOREIGN KEY (user_id)
    REFERENCES t_user(id);

ALTER TABLE t_product_to_comment
    ADD CONSTRAINT fk_product_id
    FOREIGN KEY (product_id)
    REFERENCES t_product(id);

ALTER TABLE t_product_to_comment
    ADD CONSTRAINT fk_comment_id
    FOREIGN KEY (comment_id)
    REFERENCES t_comment(id);

ALTER TABLE t_image_to_product
    ADD CONSTRAINT fk_product_id
    FOREIGN KEY (product_id)
    REFERENCES t_product(id);

ALTER TABLE t_image_to_product
    ADD CONSTRAINT fk_image_id
    FOREIGN KEY (image_id)
    REFERENCES t_image(id);

ALTER TABLE t_image_to_comment
    ADD CONSTRAINT fk_image_id
    FOREIGN KEY (image_id)
    REFERENCES t_image(id);

ALTER TABLE t_image_to_comment
    ADD CONSTRAINT fk_comment_id
    FOREIGN KEY (comment_id)
    REFERENCES t_comment(id);

ALTER TABLE t_comment
    ADD CONSTRAINT fk_user_id
    FOREIGN KEY (user_id)
    REFERENCES t_user(id);

ALTER TABLE t_comment
    ADD CONSTRAINT fk_product_id
    FOREIGN KEY (product_id)
    REFERENCES t_product(id);

ALTER TABLE t_role_to_user
    ADD CONSTRAINT fk_user_id
    FOREIGN KEY (user_id)
    REFERENCES t_user(id);

ALTER TABLE t_role_to_user
    ADD CONSTRAINT fk_role_id
    FOREIGN KEY (role_id)
    REFERENCES t_role(id);

ALTER TABLE t_product_to_order
    ADD CONSTRAINT fk_order_id
    FOREIGN KEY (order_id)
    REFERENCES t_order(id);

ALTER TABLE t_product_to_order
    ADD CONSTRAINT fk_product_id
    FOREIGN KEY (product_id)
    REFERENCES t_product(id);

ALTER TABLE t_order
    ADD CONSTRAINT fk_address_id
    FOREIGN KEY (address_id)
    REFERENCES t_pick_up_address(id);

ALTER TABLE t_order
    ADD CONSTRAINT fk_status_id
    FOREIGN KEY (status_id)
    REFERENCES t_status(id);

ALTER TABLE t_order
    ADD CONSTRAINT fk_user_id
    FOREIGN KEY (user_id)
    REFERENCES t_user(id);