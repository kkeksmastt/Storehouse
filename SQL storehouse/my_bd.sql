CREATE TABLE storehouse (
    shelf_id bigint primary key,
    shelf_name character varying(100) NOT NULL
);

INSERT INTO storehouse VALUES (1, 'A');
INSERT INTO storehouse VALUES (2, 'B');
INSERT INTO storehouse VALUES (3, 'C');
INSERT INTO storehouse VALUES (4, 'D');
INSERT INTO storehouse VALUES (5, 'E');


CREATE TABLE goods (
    product_id bigint primary key,
    product_name character varying(100) NOT NULL);

INSERT INTO goods VALUES (1, 'laptop');
INSERT INTO goods VALUES (2, 'TV');
INSERT INTO goods VALUES (3, 'phone');
INSERT INTO goods VALUES (4, 'whatch');
INSERT INTO goods VALUES (5, 'microphone');
INSERT INTO goods VALUES (6, 'coffeemaker');
INSERT INTO goods VALUES (7, 'electric kettle');

CREATE TABLE goods_shelf (
    id bigint primary key,
    product bigint NOT NULL,
    main_shelf bigint NOT NULL,
    additional_shelf integer[] NOT NULL DEFAULT '{}'::integer[],
    constraint fk_goods_shelf_product foreign key (product) references goods(product_id),
    constraint fk_goods_shelf_main_shelf foreign key (main_shelf) references storehouse(shelf_id)
    );

INSERT INTO goods_shelf VALUES (1, 1, 1, '{2,4}');
INSERT INTO goods_shelf VALUES (2, 2, 1, '{3,6}');
INSERT into goods_shelf VALUES (3, 3, 2);
INSERT into goods_shelf VALUES (4, 4, 3);
INSERT into goods_shelf VALUES (5, 5, 5);


CREATE TABLE orders (
    id bigint primary key,
    order_number integer NOT NULL,
    product integer NOT NULL,
    amount integer NOT NULL,
    constraint fk_orders_product foreign key (product) references goods (product_id)
);

INSERT INTO orders VALUES (1, 10, 1, 2);
INSERT INTO orders VALUES (2, 10, 3, 1);
INSERT INTO orders VALUES (3, 11, 2, 3);
INSERT INTO orders VALUES (4, 14, 1, 3);
INSERT INTO orders VALUES (5, 14, 6, 5);
INSERT INTO orders VALUES (6, 15, 4, 1);
INSERT INTO orders VALUES (7, 15, 5, 1);

/*select ord.order_number, ord.product as product_id, ord.amount, goods.product_name, shelfs.additional_shelf_id, sh.shelf_name from orders ord, goods, goods_shelf shelfs, storehouse sh
	where ord.product = goods.product_id and ord.product = shelfs.product and shelfs.main_shelf = sh.shelf_id
	order by shelf_name, product_id*/
