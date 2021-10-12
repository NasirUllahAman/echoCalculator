
CREATE USER 'nsair'@'localhost' IDENTIFIED BY 'Nasir@123';

GRANT ALL ON my_db.* TO 'nasir'@'localhost';

GRANT ALL PRIVILEGES ON dbTest.* To 'nasir'@'localhost' IDENTIFIED BY 'Nasir@123';

GRANT ALL PRIVILEGES ON myDB.* To 'nasir'@'localhost'; 



UPDATE mysql.user set *password_field from above* = PASSWORD('your_new_password') where user = 'root' and host = 'localhost';


UPDATE mysql.user SET Password = PASSWORD('HAMZAANIS@1213ad1231') WHERE User = 'root';
    FLUSH PRIVILEGES;
    exit;


UPDATE mysql.user SET authentication_string=PASSWORD('YOURNEWPASSWORD'), plugin='mysql_native_password' WHERE User='root' AND Host='localhost';

