DROP DATABASE IF EXISTS restaurant;
CREATE DATABASE IF NOT EXISTS restaurant;
USE restaurant;

SELECT 'CREATING DATABASE STRUCTURE' as 'INFO';
create table Tables
(
    table_id     int auto_increment
        primary key,
    chair_number int null,
    status       int null
);

create table category
(
    category_id   int auto_increment
        primary key,
    category_name varchar(100) null
);

create table customer
(
    customer_id int auto_increment
        primary key,
    number      int null
);

create table employee
(
    employee_id int auto_increment
        primary key,
    first_name  varchar(100) null,
    last_name   varchar(100) null,
    status      tinyint      null
);

create table Orders
(
    order_id     int auto_increment
        primary key,
    table_id     int          null,
    order_date   timestamp    null,
    total_amount int          null,
    status       int          null,
    customer_id  int          not null,
    employee_id  int          not null,
    review       varchar(250) null,
    constraint FK_customer_TO_order
        foreign key (customer_id) references customer (customer_id),
    constraint FK_employee_TO_order
        foreign key (employee_id) references employee (employee_id),
    constraint FK_table_TO_order
        foreign key (table_id) references Tables (table_id)
);

create table food
(
    food_id     int auto_increment
        primary key,
    food_name   varchar(100) null,
    description varchar(250) null,
    price       int          null,
    available   tinyint(1)   null,
    calorie     int          null,
    category_id int          not null,
    constraint FK_category_TO_food
        foreign key (category_id) references category (category_id)
);

create table order_list
(
    quantity      int null,
    order_list_id int auto_increment
        primary key,
    order_id      int not null,
    food_id       int not null,
    constraint FK_food_TO_order_list
        foreign key (food_id) references food (food_id),
    constraint FK_order_TO_order_list
        foreign key (order_id) references Orders (order_id)
);


INSERT INTO restaurant.customer (customer_id, number) VALUES (1, 2);
INSERT INTO restaurant.customer (customer_id, number) VALUES (2, 3);
INSERT INTO restaurant.customer (customer_id, number) VALUES (3, 3);
INSERT INTO restaurant.customer (customer_id, number) VALUES (4, 3);
INSERT INTO restaurant.customer (customer_id, number) VALUES (5, 1);
INSERT INTO restaurant.customer (customer_id, number) VALUES (6, 1);
INSERT INTO restaurant.customer (customer_id, number) VALUES (7, 1);
INSERT INTO restaurant.customer (customer_id, number) VALUES (8, 1);
INSERT INTO restaurant.customer (customer_id, number) VALUES (9, 4);
INSERT INTO restaurant.customer (customer_id, number) VALUES (10, 7);
INSERT INTO restaurant.customer (customer_id, number) VALUES (11, 2);
INSERT INTO restaurant.customer (customer_id, number) VALUES (12, 4);
INSERT INTO restaurant.customer (customer_id, number) VALUES (13, 4);
INSERT INTO restaurant.customer (customer_id, number) VALUES (14, 4);
INSERT INTO restaurant.customer (customer_id, number) VALUES (15, 4);
INSERT INTO restaurant.customer (customer_id, number) VALUES (16, 5);
INSERT INTO restaurant.customer (customer_id, number) VALUES (17, 5);
INSERT INTO restaurant.customer (customer_id, number) VALUES (18, 5);
INSERT INTO restaurant.customer (customer_id, number) VALUES (19, 2);
INSERT INTO restaurant.customer (customer_id, number) VALUES (20, 2);
INSERT INTO restaurant.customer (customer_id, number) VALUES (21, 0);
INSERT INTO restaurant.customer (customer_id, number) VALUES (22, 0);
INSERT INTO restaurant.customer (customer_id, number) VALUES (23, 0);


INSERT INTO restaurant.employee (employee_id, first_name, last_name, status) VALUES (1, 'Liam', 'Williams', 0);
INSERT INTO restaurant.employee (employee_id, first_name, last_name, status) VALUES (2, 'Noah', 'Jones', 1);
INSERT INTO restaurant.employee (employee_id, first_name, last_name, status) VALUES (3, 'Elijah', 'Davis', 1);
INSERT INTO restaurant.employee (employee_id, first_name, last_name, status) VALUES (4, 'William', 'Rodriguez', 0);
INSERT INTO restaurant.employee (employee_id, first_name, last_name, status) VALUES (5, 'Benjamin', 'Martinez', 1);
INSERT INTO restaurant.employee (employee_id, first_name, last_name, status) VALUES (6, 'Lucas', 'Hernandez', 2);
INSERT INTO restaurant.employee (employee_id, first_name, last_name, status) VALUES (7, 'Henry', 'Lopez', 3);
INSERT INTO restaurant.employee (employee_id, first_name, last_name, status) VALUES (8, 'Theodore', 'Gonzalez', 3);
INSERT INTO restaurant.employee (employee_id, first_name, last_name, status) VALUES (9, 'Jack', 'Wilson', 2);
INSERT INTO restaurant.employee (employee_id, first_name, last_name, status) VALUES (10, 'Sebastian', 'Anderson', 1);
INSERT INTO restaurant.employee (employee_id, first_name, last_name, status) VALUES (11, 'p', 'i', 0);
INSERT INTO restaurant.employee (employee_id, first_name, last_name, status) VALUES (12, 'yugi', 'ho', 0);
INSERT INTO restaurant.employee (employee_id, first_name, last_name, status) VALUES (13, 'jo', 'go', 3);
INSERT INTO restaurant.employee (employee_id, first_name, last_name, status) VALUES (14, 'shiji', 'totoro', 2);

INSERT INTO restaurant.category (category_id, category_name) VALUES (1, 'Lasagne');
INSERT INTO restaurant.category (category_id, category_name) VALUES (2, 'Desserts');
INSERT INTO restaurant.category (category_id, category_name) VALUES (3, 'Beverages');
INSERT INTO restaurant.category (category_id, category_name) VALUES (4, 'Salads');
INSERT INTO restaurant.category (category_id, category_name) VALUES (5, 'Soups');
INSERT INTO restaurant.category (category_id, category_name) VALUES (6, 'Pasta');
INSERT INTO restaurant.category (category_id, category_name) VALUES (7, 'Pizza');
INSERT INTO restaurant.category (category_id, category_name) VALUES (8, 'Burgers');
INSERT INTO restaurant.category (category_id, category_name) VALUES (9, 'Sandwiches');
INSERT INTO restaurant.category (category_id, category_name) VALUES (10, 'Steaks');
INSERT INTO restaurant.category (category_id, category_name) VALUES (11, 'Breakfast');
INSERT INTO restaurant.category (category_id, category_name) VALUES (12, 'noodel');
INSERT INTO restaurant.category (category_id, category_name) VALUES (13, 'Sushi');
INSERT INTO restaurant.category (category_id, category_name) VALUES (14, 'curry');
INSERT INTO restaurant.category (category_id, category_name) VALUES (15, 'sushi');

INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (1, 'Classic Lasagne', 'Traditional Italian lasagne with beef and cheese', 13, 1, 850, 1);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (2, 'Vegetarian Lasagne', 'Lasagne made with fresh vegetables and cheese', 11, 1, 700, 1);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (3, 'Seafood Lasagne', 'Lasagne with shrimp, crab, and creamy sauce', 15, 0, 900, 1);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (4, 'Chocolate Cake', 'Rich and moist chocolate cake', 6, 1, 450, 2);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (5, 'Cheesecake', 'Creamy cheesecake with a graham cracker crust', 7, 1, 400, 2);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (6, 'Apple Pie', 'Warm apple pie with cinnamon and nutmeg', 5, 1, 350, 2);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (7, 'Coca Cola', 'Classic soft drink', 2, 1, 150, 3);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (8, 'Lemonade', 'Freshly squeezed lemonade', 3, 1, 120, 3);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (9, 'Iced Coffee', 'Chilled coffee with milk and sugar', 4, 1, 100, 3);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (10, 'Caesar Salad', 'Romaine lettuce with Caesar dressing', 8, 1, 300, 4);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (11, 'Greek Salad', 'Fresh salad with feta cheese, olives, and cucumber', 9, 1, 250, 4);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (12, 'Cobb Salad', 'Salad with chicken, avocado, and bacon', 10, 0, 400, 4);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (13, 'Tomato Soup', 'Creamy tomato soup with herbs', 5, 1, 180, 5);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (14, 'Chicken Noodle Soup', 'Classic chicken soup with noodles', 6, 1, 200, 5);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (15, 'Minestrone', 'Italian vegetable soup', 7, 1, 220, 5);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (16, 'Spaghetti Bolognese', 'Pasta with meat sauce', 12, 1, 700, 6);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (17, 'Fettuccine Alfredo', 'Pasta with creamy Alfredo sauce', 13, 1, 800, 6);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (18, 'Penne Arrabbiata', 'Spicy tomato pasta', 11, 0, 650, 6);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (19, 'Margherita Pizza', 'Classic pizza with tomato, mozzarella, and basil', 10, 1, 800, 7);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (20, 'Pepperoni Pizza', 'Pizza topped with pepperoni slices', 11, 1, 900, 7);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (21, 'BBQ Chicken Pizza', 'Pizza with BBQ chicken and onions', 12, 1, 950, 7);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (22, 'Cheeseburger', 'Classic burger with cheese, lettuce, and tomato', 9, 1, 750, 8);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (23, 'Bacon Burger', 'Burger with bacon and cheddar cheese', 10, 1, 850, 8);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (24, 'Veggie Burger', 'Burger made with plant-based patty', 8, 0, 600, 8);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (25, 'Club Sandwich', 'Triple-layer sandwich with turkey, bacon, and lettuce', 8, 1, 500, 9);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (26, 'BLT Sandwich', 'Bacon, lettuce, and tomato sandwich', 7, 1, 450, 9);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (27, 'Grilled Cheese', 'Toasted sandwich with melted cheese', 6, 1, 400, 9);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (28, 'Ribeye Steak', 'Grilled ribeye with garlic butter', 20, 1, 950, 10);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (29, 'Filet Mignon', 'Tender beef filet with sauce', 25, 1, 850, 10);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (30, 'Sirloin Steak', 'Steak with grilled vegetables', 18, 1, 900, 10);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (31, 'Pancakes', 'Fluffy pancakes with syrup', 7, 1, 500, 11);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (32, 'Omelette', 'Three-egg omelette with cheese and vegetables', 8, 1, 400, 11);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (33, 'French Toast', 'Golden-brown toast with powdered sugar', 7, 1, 450, 11);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (34, 'pt', 'good', 13, 1, 300, 10);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (35, 'su', 'class', 250, 2, 800, 2);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (36, 'Pizza Margherita', 'Classic pizza with tomatoes, mozzarella, and basil', 250, 0, 800, 2);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (37, 'Spaghetti Carbonara', 'A classic Italian pasta dish with eggs, cheese, pancetta, and pepper.', 250, 0, 500, 2);
INSERT INTO restaurant.food (food_id, food_name, description, price, available, calorie, category_id) VALUES (38, 'Spaghetti Carbonara', 'A classic Italian pasta dish with eggs, cheese, pancetta, and pepper.', 300, 0, 500, 2);

INSERT INTO restaurant.Tables (table_id, chair_number, status) VALUES (1, 2, 0);
INSERT INTO restaurant.Tables (table_id, chair_number, status) VALUES (2, 4, 0);
INSERT INTO restaurant.Tables (table_id, chair_number, status) VALUES (3, 4, 1);
INSERT INTO restaurant.Tables (table_id, chair_number, status) VALUES (4, 4, 1);
INSERT INTO restaurant.Tables (table_id, chair_number, status) VALUES (5, 1, 1);
INSERT INTO restaurant.Tables (table_id, chair_number, status) VALUES (6, 1, 1);
INSERT INTO restaurant.Tables (table_id, chair_number, status) VALUES (7, 1, 1);
INSERT INTO restaurant.Tables (table_id, chair_number, status) VALUES (8, 1, 1);
INSERT INTO restaurant.Tables (table_id, chair_number, status) VALUES (9, 4, 1);
INSERT INTO restaurant.Tables (table_id, chair_number, status) VALUES (10, 8, 1);
INSERT INTO restaurant.Tables (table_id, chair_number, status) VALUES (11, 8, 0);
INSERT INTO restaurant.Tables (table_id, chair_number, status) VALUES (12, 2, 0);
INSERT INTO restaurant.Tables (table_id, chair_number, status) VALUES (13, 0, 0);
INSERT INTO restaurant.Tables (table_id, chair_number, status) VALUES (14, 2, 0);
INSERT INTO restaurant.Tables (table_id, chair_number, status) VALUES (15, 0, 0);

INSERT INTO restaurant.Orders (order_id, table_id, order_date, total_amount, status, customer_id, employee_id, review) VALUES (1, 1, '2024-09-26 12:30:00', 33, 4, 1, 1, '');
INSERT INTO restaurant.Orders (order_id, table_id, order_date, total_amount, status, customer_id, employee_id, review) VALUES (2, 2, '2024-09-26 13:00:00', 29, 3, 2, 3, '');
INSERT INTO restaurant.Orders (order_id, table_id, order_date, total_amount, status, customer_id, employee_id, review) VALUES (3, 3, '2024-09-26 14:00:00', 20, 3, 3, 2, 'ซอยทำออกมาได้เยียมมาก');
INSERT INTO restaurant.Orders (order_id, table_id, order_date, total_amount, status, customer_id, employee_id, review) VALUES (4, 4, '2024-09-26 15:00:00', 82, 4, 4, 1, 'ขออันนี้เพิ่มอีก');
INSERT INTO restaurant.Orders (order_id, table_id, order_date, total_amount, status, customer_id, employee_id, review) VALUES (5, 5, '2024-09-26 16:00:00', 42, 1, 5, 2, '');
INSERT INTO restaurant.Orders (order_id, table_id, order_date, total_amount, status, customer_id, employee_id, review) VALUES (6, 6, '2024-09-26 17:00:00', 51, 2, 6, 3, '');
INSERT INTO restaurant.Orders (order_id, table_id, order_date, total_amount, status, customer_id, employee_id, review) VALUES (7, 7, '2024-09-26 18:00:00', 30, 3, 7, 1, 'อยากเปลี่ยนร้าน');
INSERT INTO restaurant.Orders (order_id, table_id, order_date, total_amount, status, customer_id, employee_id, review) VALUES (8, 8, '2024-09-26 19:00:00', 53, 4, 8, 2, 'ขมไปนิด');
INSERT INTO restaurant.Orders (order_id, table_id, order_date, total_amount, status, customer_id, employee_id, review) VALUES (9, 9, '2024-09-26 20:00:00', 65, 4, 9, 4, 'เเม่ค้าน่ารักมากก');
INSERT INTO restaurant.Orders (order_id, table_id, order_date, total_amount, status, customer_id, employee_id, review) VALUES (10, 10, '2024-09-26 21:00:00', 44, 4, 10, 1, 'ไว้จะกลับมากินอีก');
INSERT INTO restaurant.Orders (order_id, table_id, order_date, total_amount, status, customer_id, employee_id, review) VALUES (12, 1, '2024-09-26 21:00:00', 300, 1, 1, 1, '');
INSERT INTO restaurant.Orders (order_id, table_id, order_date, total_amount, status, customer_id, employee_id, review) VALUES (13, 1, '2024-09-26 21:00:00', 26, 1, 1, 1, '');
INSERT INTO restaurant.Orders (order_id, table_id, order_date, total_amount, status, customer_id, employee_id, review) VALUES (15, 11, '2024-10-01 10:39:09', 41, 1, 15, 3, '');
INSERT INTO restaurant.Orders (order_id, table_id, order_date, total_amount, status, customer_id, employee_id, review) VALUES (16, 11, '2024-10-01 10:41:36', 45, 1, 16, 2, '');
INSERT INTO restaurant.Orders (order_id, table_id, order_date, total_amount, status, customer_id, employee_id, review) VALUES (17, 11, '2024-10-01 10:44:20', 45, 3, 17, 1, 'มันเลวร้ายมาก');
INSERT INTO restaurant.Orders (order_id, table_id, order_date, total_amount, status, customer_id, employee_id, review) VALUES (18, 1, '2024-10-02 01:51:34', 45, 3, 19, 2, 'ร้อน');
INSERT INTO restaurant.Orders (order_id, table_id, order_date, total_amount, status, customer_id, employee_id, review) VALUES (19, 12, '2024-10-02 01:54:59', 45, 3, 20, 2, 'จำได้ว่ามีธุระ');
INSERT INTO restaurant.Orders (order_id, table_id, order_date, total_amount, status, customer_id, employee_id, review) VALUES (20, 13, '2024-10-03 00:55:53', 45, 1, 21, 14, '');
INSERT INTO restaurant.Orders (order_id, table_id, order_date, total_amount, status, customer_id, employee_id, review) VALUES (21, 13, '2024-10-03 01:00:47', 45, 1, 22, 3, '');
INSERT INTO restaurant.Orders (order_id, table_id, order_date, total_amount, status, customer_id, employee_id, review) VALUES (22, 13, '2024-10-03 01:05:43', 45, 4, 23, 12, 'อร่อย');

INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (2, 1, 1, 1);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 2, 1, 5);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (3, 3, 2, 8);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (2, 4, 2, 12);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 5, 3, 15);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 6, 3, 17);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 7, 4, 2);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (2, 8, 4, 3);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 9, 5, 4);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (3, 10, 5, 6);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (2, 11, 6, 7);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 12, 6, 9);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 13, 7, 10);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 14, 7, 11);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (2, 15, 8, 13);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 16, 8, 14);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 17, 9, 16);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (2, 18, 9, 18);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 19, 10, 19);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 20, 10, 20);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 21, 4, 2);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (2, 22, 4, 3);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 23, 5, 4);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (3, 24, 5, 6);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (2, 25, 6, 11);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 26, 6, 12);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (3, 27, 6, 13);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 28, 7, 14);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 29, 7, 15);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (2, 30, 8, 16);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 31, 8, 17);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 32, 9, 18);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (2, 33, 9, 19);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 34, 10, 20);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 35, 10, 21);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (2, 36, 13, 1);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (2, 37, 15, 1);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 38, 15, 3);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (2, 39, 16, 3);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 40, 16, 3);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (2, 41, 17, 3);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 42, 17, 3);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (2, 43, 18, 3);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 44, 18, 3);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (2, 45, 19, 3);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 46, 19, 3);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (2, 47, 20, 3);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 48, 20, 3);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (2, 49, 21, 3);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 50, 21, 3);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (2, 51, 22, 3);
INSERT INTO restaurant.order_list (quantity, order_list_id, order_id, food_id) VALUES (1, 52, 22, 3);

UPDATE Orders o
JOIN (
    SELECT ol.order_id, SUM(f.price * ol.quantity) AS total
    FROM order_list ol
    JOIN food f ON ol.food_id = f.food_id
    GROUP BY ol.order_id
) ol ON o.order_id = ol.order_id
SET o.total_amount = ol.total;

GRANT ALL PRIVILEGES ON restaurant.* TO 'user'@'%';
FLUSH PRIVILEGES;