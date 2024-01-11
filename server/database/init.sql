-- create new database sampler_test_db and use it
drop database if exists sampler_test_db;
create database sampler_test_db;
use sampler_test_db;
go

-- create a new table called '[Accounts]' in schema '[dbo]'
if object_id('[dbo].[Accounts]', 'U') is not null
    drop table [dbo].[Accounts]
go

-- create the table in the specified schema
create table [dbo].[Accounts]
(
    [Id]        varchar(100)  not null primary key default newid(),
    [Email]     nvarchar(50)  not null unique,
    [Password]  nvarchar(200) not null,
    [Name]      nvarchar(200) not null,
    [CreatedAt] DATETIME      not null             default getdate(),
);
go

-- delete procedure if exists
if object_id('insert_account', 'P') is not null
    drop procedure insert_account
go

-- create a procedure to insert a new row into the accounts table
-- the procedure should take 3 parameters: email, password and name
-- the procedure should return all the columns of the newly inserted row
create procedure insert_account @email varchar(100),
                                @password varchar(500),
                                @name varchar(100)
as
begin
    insert into Accounts (Email, Password, Name)
    values (@email, @password, @name)

    select TOP (1) Id, Email, Name
    from Accounts
    where Email = @email
end
go

-- delete procedure if exists
if object_id('get_account_by_email', 'P') is not null
    drop procedure get_account_by_email
go

-- create a procedure to get an account by email
-- the procedure should take 1 parameter: email
-- the procedure should return all the columns of the account
create procedure get_account_by_email @email varchar(100)
as
begin
    select TOP (1) Id, Email, Password, Name, CreatedAt
    from Accounts
    where Email = @email
end
go
