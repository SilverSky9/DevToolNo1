CREATE TABLE post (
  ID SERIAL PRIMARY KEY,
  Title TEXT,
  Details TEXT,
  Tag TEXT,
  PostType TEXT
);
INSERT INTO post (ID, title, details, tag, PostType)
VALUES (1, 'อยากได้ค้อนครับ', 'อยากได้ค้อนเอามาทุบโต๊ะเล่นครับ', 'เครื่องมือช่าง', 'ต้องการ');
INSERT INTO post (ID, title, details, tag, PostType)
VALUES (2, 'อยากได้น้ำแข็ง', 'อากาศร้อนๆแบบนี้ เฉาก๋วยเย็นๆก็ดีเหมือนกัน', 'อาหาร', 'ต้องการ');
