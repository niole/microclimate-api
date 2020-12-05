CREATE DATABASE IF NOT EXISTS microclimates;
CREATE TABLE IF NOT EXISTS microclimates.Deployments (
    Id varchar(36) PRIMARY KEY NOT NULL,
    OwnerUserId varchar(36) NOT NULL,
    Status ENUM('unreachable', 'starting_up', 'failed', 'running', 'shutting_down') NOT NULL,
    Name varchar(255) NOT NULL UNIQUE
);
CREATE TABLE IF NOT EXISTS microclimates.Users (
    Id varchar(36) PRIMARY KEY NOT NULL,
    Email varchar(255) NOT NULL UNIQUE,
    Name varchar(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS microclimates.Peripherals (
    Id varchar(36) PRIMARY KEY NOT NULL,
    OwnerUserId varchar(36) NOT NULL,
    DeploymentId varchar(36) NOT NULL,
    HardwareId varchar(36) NOT NULL,
    Type ENUM('THERMAL', 'PARTICLE') NOT NULL,
    Unit ENUM('F', 'C', 'PM2.5', '%') NOT NULL,
    Name varchar(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS microclimates.PeripheralEvents (
    Id varchar(36) PRIMARY KEY NOT NULL,
    PeripheralId varchar(36) NOT NULL,
    DeploymentId varchar(36) NOT NULL,
    Value float NOT NULL,
    Timestamp timestamp NOT NULL
);
LOAD DATA INFILE '/docker-entrypoint-initdb.d/init_data/init_users.csv'
INTO TABLE microclimates.Users
FIELDS TERMINATED BY ','
LINES TERMINATED BY '\n' (Id,Email,Name);

LOAD DATA INFILE '/docker-entrypoint-initdb.d/init_data/init_deployments.csv'
INTO TABLE microclimates.Deployments
FIELDS TERMINATED BY ','
LINES TERMINATED BY '\n' (Id,OwnerUserId,Status,Name);

LOAD DATA INFILE '/docker-entrypoint-initdb.d/init_data/init_peripherals.csv'
INTO TABLE microclimates.Peripherals
FIELDS TERMINATED BY ','
LINES TERMINATED BY '\n' (Id,OwnerUserId,DeploymentId,HardwareId,Type,Unit,Name);

LOAD DATA INFILE '/docker-entrypoint-initdb.d/init_data/init_events.csv'
INTO TABLE microclimates.PeripheralEvents
FIELDS TERMINATED BY ','
LINES TERMINATED BY '\n' (Id,PeripheralId,DeploymentId,Value,Timestamp);
