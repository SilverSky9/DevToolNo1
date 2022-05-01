# DevToolNo1
This project is a part of SOFTWARE DEVELOPMENT TOOLS AND ENVIRONMENTS 2/2564

เว็บแอปพลิเคชั่นสร้างขึ้นเพื่อแก้ไขปัญหาในการซื้อขายของกันภายในมหาวิทยาลัย

## Members

|Name|Student ID|
|:--:|-----------|
|NUTCHAYAPORN SIANGKHIO|62070059|
|PAWARIS WONGSAIED|62070116|
|NOPPHADON PHANWONG|62070101|
|SINLAPAWIT RITNAMKHAM|62070178|
|WARANGKHANA SUKPARTCHAROEN|62070166|
|SIPANG KLINHOM|62070196|

## รูปแบบของ commit message
|commit message|ความหมาย|
|:------------:|--------|
|``build``     |การเปลี่ยนแปลงที่ส่งผลต่อระบบ build (ตัวอย่างเช่น gulp, broccoli, npm)|
|``ci``        |การเปลี่ยนแปลงไฟล์ การกำหนดค่า configuration files และ scripts  (ตัวอย่างเช่น Travis, Circle, BrowserStack, SauceLabs)|
|``docs``      |การเปลี่ยนแปลงเอกสารต่าง ๆ เท่านั้น|
|``feat``      |ฟีเจอร์ใหม่|
|``fix``       |แก้ไข Bug|
|``perf``      |ปรับเปลี่ยน Code เพื่อเพิ่มประสิทธิภาพ|
|``refactor``  |การ Refactoring Code|
|``style``     |แก้ไขการเขียน Code เช่น white-space, formatting, missing semi-colons, etc.|
|``test``      |การทดสอบ หรือการแก้ไขการทดสอบที่มีอยู่|
|``add``       |เป็นการเพิ่มหรือสร้างไฟล์ในโปรเจค

## Branch strategy
#### Git Flow Branch Strategy
|Branch name |   วิธีใช้        |
|:----------:|-------------|
|``Main``        |เป็น Branch หลักของโปรเจคใช้สำหรับเมื่อโปรเจคเสร็จแล้วเรียบร้อยแล้ว|
|``Dev``        |เป็น Branch ที่ใช้สำหรับการพัฒนาของโปรเจค เมื่อมี Feature  ใดที่ทำเสร็จแล้ว จะทำการ Merge Branch ``Feature`` นั้นด้วยวิธีการ PR เข้า Branch ``Dev`` นี้|
|``Feature``        |เป็น Branch ที่แยกออกมาพัฒนาจาก ``Dev`` เพื่อแบ่ง Feature ในการพัฒนาส่วนต่างๆ |
|``Release``        |เป็น Branch ที่ใช้สำหรับการปล่อยให้โปรเจคออกมาให้พร้อมใช้งานเป็นระยะๆ|
| ``Test``          | เป็น Branch ที่ใช้สำหรับการทำ Testing ต่างๆ
|``Hotfix``        |เป็น Branch ไว้สำหรับการแก้ไขปัญหาที่เร่งด่วน จำเป็นต้องแก้ไขทันที|
