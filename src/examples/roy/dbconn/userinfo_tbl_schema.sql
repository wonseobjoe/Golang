# 1. Login as root
# 	mysql -uroot -proot
# 2. Create test_db database
# 	create database test_db;
# 3. If testuser not exists, create new user (Note. remove "password expire" statement if you don't want to change password when testuser logs into mysql)
#	create user 'testuser'@'localhost' identified by 'testuser' password expire;
# -  if missing password expire,
#	alter user 'testuser'@'localhost' password expire;
# 4. Grant privileges to schema owner for test_db
#	grant all priviliges on test_db.* to 'testuser'@'localhost';
# 5. Logout
# 	quit
# 6. Login as testuser
#	mysql -utestuser -ptestuser
# 7. Reset password of testuser by yourself
#	set password = password('testuser');
# 8. Change database to use test_db
#	use test_db;

# 9 Create new table
CREATE TABLE `userinfo_tbl` (
    `uid` INT(10) NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(64) NULL DEFAULT NULL,
    `departname` VARCHAR(64) NULL DEFAULT NULL,
    `created` DATE NULL DEFAULT NULL,
    PRIMARY KEY (`uid`)
);