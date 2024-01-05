-- Create a new table called '[Accounts]' in schema '[dbo]'
-- Drop the table if it already exists
IF OBJECT_ID('[dbo].[Accounts]', 'U') IS NOT NULL
    DROP TABLE [dbo].[Accounts]
GO
-- Create the table in the specified schema
CREATE TABLE [dbo].[Accounts]
(
    [Id]        UNIQUEIDENTIFIER NOT NULL PRIMARY KEY DEFAULT NEWID(),
    [Email]     NVARCHAR(50)     NOT NULL,
    [Password]  NVARCHAR(50)     NOT NULL,
    [Name]      NVARCHAR(50)     NOT NULL,
    [CreatedAt] DATETIME         NOT NULL DEFAULT GETDATE(),
    -- Specify more columns here
);
GO