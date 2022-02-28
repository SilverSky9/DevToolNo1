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

## Development

### Front end
```
1. cd frontend
2. npm install
3. npm run dev
```

### Back end
**เลือกวิธีไหนก็ได้**
1. Run ด้วย Go
```
- cd backend
- go run main.go
```

2. Run ด้วย Docker <Br/>
- ติดตั้ง Docker Desktop สำหรับเอาไว้ Run Docker ดูวิธีติดตั้งและตั้งค่าได้ที่นี่ [สำหรับ Window](https://blog.codedsir.com/how-to-install-docker-desktop-on-windows/) และ [สำหรับ OSX](https://medium.com/@thep_p/%E0%B8%95%E0%B8%B4%E0%B8%94%E0%B8%95%E0%B8%B1%E0%B9%89%E0%B8%87-docker-%E0%B8%89%E0%B8%9A%E0%B8%B1%E0%B8%9A%E0%B9%82%E0%B8%AB%E0%B8%A5%E0%B8%94%E0%B8%9B%E0%B8%B1%E0%B9%8A%E0%B8%9B%E0%B9%80%E0%B8%9B%E0%B8%B4%E0%B8%94%E0%B8%9B%E0%B8%B8%E0%B9%8A%E0%B8%9B-8bace185d36b) <Br/>

```
- docker compose build
- docker compose up
```

## Testing
Backend
```
docker compose build
docker compose up -d
docker exec -it go_container sh
go test ./... -cover
```
Frontend
```
cd frontend
npm run dev
npx cypress open
```
1. เมื่อรันแล้วจะมีหน้า Chrome Pop-up ขึ้นมา
2. เลือกรันไฟล์ ``home_spec.js`` , ``index_spce.js``, ``End-to-end_spce.js``
3. Cypress จะรันตาม Testcase ที่เขียนเอาไว้
