# DevToolNo1
This project is a part of SOFTWARE DEVELOPMENT TOOLS AND ENVIRONMENTS 2/2564

## Members
- Nutchayaporn
- Pawaris
- Nopphadon
- Sinlapawit
- Warangkhana
- Sipang


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

Reference : [commit message](https://github.com/angular/angular/blob/22b96b9/CONTRIBUTING.md#type)

## Branch strategy
#### Git Flow Branch Strategy
|Branch name |   วิธีใช้        |
|:----------:|-------------|
|``Main``        |เป็น Branch หลักของโปรเจคใช้สำหรับเมื่อโปรเจคเสร็จแล้วเรียบร้อยแล้ว|
|``Dev``        |เป็น Branch ที่ใช้สำหรับการพัฒนาของโปรเจค เมื่อมี Feature  ใดที่ทำเสร็จแล้ว จะทำการ Merge Branch ``Feature`` นั้นด้วยวิธีการ PR เข้า Branch ``Dev`` นี้|
|``Feature``        |เป็น Branch ที่แยกออกมาพัฒนาจาก ``Dev`` เพื่อแบ่ง Feature ในการพัฒนาส่วนต่างๆ |
|``Release``        |เป็น Branch ที่ใช้สำหรับการปล่อยให้โปรเจคออกมาให้พร้อมใช้งานเป็นระยะๆ|
|``Hotfix``        |เป็น Branch ไว้สำหรับการแก้ไขปัญหาที่เร่งด่วน จำเป็นต้องแก้ไขทันที|


## หลักการทำงานร่วมกันในทีมของเรา
ทีมของพวกเราเลือกใช้ Pull Request ในการ Merge Branch เข้าด้วยกัน<Br/>
โดยหลักการ Pull Request ของทีมเราคือ
- ทำการเปิด PR ทุกครั้งที่ต้องการ Merge Branch ใดๆ
- การเปิด PR ให้เพิ่ม Reviewer **ที่มีส่วนเกี่ยวข้องกับการ Code ส่วนนั้น**ให้มา Approve PR ให้ **อย่างน้อย 1 คน**
- ผู้ที่เป็น Reviewer ต้องตรวจสอบ Source Code ให้คำแนะนำ หรือ จุดต้องปรับแก้ ให้แก่ผู้ที่ต้องการ PR
- เมื่อผู้ที่เป็น Reviewer ตรวจสอบเสร็จแล้ว ต้องอมุมัติให้ผู้ขอ PR สามารถ Merge เข้าสู่ Branch ได้
- Branch ``Main`` ได้มีการปกป้อง เพื่อไม่ให้มีการเผลอ Merge ได้ มีกฎจำเป็นต้องมี Reviewer **ตั้งแต่ 2 คนขึ้นไป**