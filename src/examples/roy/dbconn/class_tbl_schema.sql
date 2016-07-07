# 1. Login as root
# 	mysql -uroot -proot
# 2. Create school_db database
# 	create database school_db;
# 3. If testuser not exists, create new user (Note. remove "password expire" statement if you don't want to change password when testuser logs into mysql)
#	create user 'testuser'@'localhost' identified by 'testuser' password expire;
# -  if missing password expire,
#	alter user 'testuser'@'localhost' password expire;
# 4. Grant privileges to schema owner for school_db
#	grant all priviliges on school_db.* to 'testuser'@'localhost';
# 5. Logout
# 	quit
# 6. Login as testuser
#	mysql -utestuser -ptestuser
# 7. If password isn't reset yet, reset password of testuser by yourself
#	set password = password('testuser');
# 8. Change database to use school_db
#	use school_db;

# 9 Create new table
CREATE TABLE class_tbl (
	'id' int, 
	'name' varchar(50), 
	'gpa' float
);

# 10. Insert sample record
#	insert into class_tbl values (11, 'Roy Joe', 4.8);
#	insert into class_tbl values (22, 'Seoung Woo Baek', 4.9);
#	insert into class_tbl values (33, 'Keun Hye Park', 1.1);
# 	insert into class_tbl values (44, 'Myoung Bark Lee', 1.0);