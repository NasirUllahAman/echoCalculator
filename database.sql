CREATE DATABASE calculator;
use calculator;
CREATE TABLE calcu (
    ID int auto_increment primary key,
    number1 float,
    number2 float,
	operation char,
    result float,
    createdAt timestamp default current_timestamp
);
insert into calcu(number1,number2,operation,result) Values (6,8,'+',14);
Insert into calcu(number1,number2,operation,result) Values (2,3,'*',6);
select * from calcu;


