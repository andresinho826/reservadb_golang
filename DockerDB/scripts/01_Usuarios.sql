ALTER SESSION SET CONTAINER=RESERVASDB;
CREATE USER TESTUSER IDENTIFIED BY 123;
GRANT CONNECT, RESOURCE TO TESTUSER;
GRANT SELECT ANY TABLE TO TESTUSER;

CREATE USER RES IDENTIFIED BY res;

GRANT CREATE SESSION, CREATE TABLE, CREATE VIEW, UNLIMITED TABLESPACE TO RES;