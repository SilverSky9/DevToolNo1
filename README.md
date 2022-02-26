# DevToolNo1
This project is a part of SOFTWARE DEVELOPMENT TOOLS AND ENVIRONMENTS 2/2564

## Members
- Nutchayaporn
- Pawaris
- Nopphadon
- Sinlapawit
- Warangkhana
- Sipang

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
go test ./... -cover
```
Frontend
```
npx cypress open
```
1. เมื่อรันแล้วจะมีหน้า Chrome Pop-up ขึ้นมา
2. เลือกรันไฟล์ ``home_spec.js`` or ``index_spce.js``
3. Cypress จะรันตาม Testcase ที่เขียนเอาไว้