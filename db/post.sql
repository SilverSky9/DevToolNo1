CREATE TABLE post (
  id SERIAL PRIMARY KEY,
  title TEXT,
  details TEXT,
  tag TEXT,
  typePost TEXT
);
INSERT INTO post (title, details, tag, typePost)
VALUES (1, 'อยากได้ค้อนครับ', 'อยากได้ค้อนเอามาทุบโต๊ะเล่นครับ', 'เครื่องมือช่าง', 'ต้องการ');
INSERT INTO post (title, details, tag, typePost)
VALUES (2, 'อยากได้น้ำแข็ง', 'อากาศร้อนๆแบบนี้ เฉาก๋วยเย็นๆก็ดีเหมือนกัน', 'อาหาร', 'ต้องการ');
