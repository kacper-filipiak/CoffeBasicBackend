CREATE TABLE coffe_bags (
id SERIAL PRIMARY KEY,
weight FLOAT NOT NULL,
roast_date INT,
coffe_type_id INT,
CONSTRAINT coffee_bag_coffee_type_fk FOREIGN KEY (coffe_type_id) REFERENCES coffe_types(id)
); 
