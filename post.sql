CREATE TABLE post (
  PostId SERIAL PRIMARY KEY,
  ProductName TEXT,
  PostDate TIME,
  ProductOption TEXT,
  Price INT,
  Amount INT,
  PinId INT,
  TagId INT
);
INSERT INTO post (PostId, ProductName, PostDate, ProductOption, Price, Amount, PinId, TagId)
VALUES (1, 'อยากได้ค้อนครับ','00:00:00.0000000','ต้องการ', 123, 1 , 1234, 1346)
