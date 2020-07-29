create database items;



create table usersec(
	userid varchar(255),
	mainmenu varchar(255),
	menuname varchar(255)
);

insert into usersec values('rony','Accounting Head Entry','ACCOUNTS');
insert into usersec values('suman','Create User','ADMIN');
insert into usersec values('rony','Card Entry','INVENTORY');
insert into usersec values('rony','Purchase Product Search Details','RECEIVED GOODS');
insert into usersec values('rony','Ledger Book','ACCOUNTS');
insert into usersec values('alex','Unit Entry','INVENTORY');




create table logins
(
	username varchar(255),
	email varchar(255),
	password1 varchar(255)
);

insert into logins values('a','a@a.a','a');



create table purchases
(
	item_id int unique,
	item_name varchar(255),
	item_quantity float,
	item_rate float,
	item_purchase_date date
);

insert into purchases values(1,'pencil',20,5.5,'2020-05-09');
insert into purchases values(2,'pen',10,5.5,'2020-04-04');
insert into purchases values(3,'rubber',5,5.5,'2019-01-01');