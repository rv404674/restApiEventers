CREATE SCHEMA EventersTestSchema;
GO

CREATE TABLE EventersTestSchema.EventDetails(
  Id INT IDENTITY (1,1) NOT NULL PRIMARY KEY,
  Name NVARCHAR(50),
  Location NVARCHAR(50),
  eventdate NVARCHAR(50)
);
GO

INSERT INTO EventersTestSchema.EventDetails (Name, Location, eventdate) VALUES
(N'Wedding', N'London', N'2021-06-31'),
(N'Birthday', N'India', N'2020-05-28'),
(N'Engagement', N'Paris', N'2030-05-10');
GO

SELECT  * FROM EventersTestSchema.EventDetails;
GO
