CREATE DATABASE NewCalculator;
use NewCalculator;
CREATE TABLE calcu (
    ID int auto_increment primary key,
    number1 int,
    number2 int,
	operation char,
    result float,
    createdAt timestamp
);
select * from calcu;
