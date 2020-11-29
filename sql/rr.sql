CREATE TABLE `RR` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `name` varchar(255), `ttl` integer, `type` varchar(255), `rdata` varchar(255));

INSERT INTO RR (name,ttl,type,rdata) VALUES ("www.test.com.",20,"A","1.1.1.1");

UPDATE RR SET ttl=60, rdata="2.2.2.2" WHERE id=1;

DELETE FROM RR;