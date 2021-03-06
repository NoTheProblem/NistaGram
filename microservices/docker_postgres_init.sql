CREATE USER "auth-service" WITH PASSWORD 'test' CREATEDB;
CREATE USER "verification-service" WITH PASSWORD 'test' CREATEDB;
CREATE USER "user-service" WITH PASSWORD 'test' CREATEDB;


CREATE DATABASE "auth-db"
    WITH 
    OWNER = "auth-service"
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;


CREATE DATABASE "user-db"
    WITH 
    OWNER = "user-service"
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

CREATE DATABASE "verification-db"
    WITH 
    OWNER = "verification-service"
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;