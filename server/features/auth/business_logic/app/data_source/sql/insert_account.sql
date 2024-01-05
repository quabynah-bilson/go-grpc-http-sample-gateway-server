-- create a procedure to insert a new row into the accounts table
-- the procedure should take 2 parameters: email and password
create procedure insert_account
    @email varchar(100),
    @password varchar(100)
as
begin
    insert into Accounts (Email, Password)
    values (@email, @password)
end
go