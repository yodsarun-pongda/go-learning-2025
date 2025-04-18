# My own Go lang Learning

## Setup project
1. Install Package
```
go mod tidy
```
2. Start the application
```
go run MyLearningApplication.go
```

## Build Go Application
1. build application
```
go build
```
1.1 custom architech
GOOS คือ เลือกระบบปฏิบัติการที่จะ Build
GOARCH คือ สถาปัตยกรรมของ CPU ที่จะ Build
```
env GOOS=windows GOARCH=amd64 go build
```
2. Running go Application
```
./goapp
```

Yodsarun Pongda\
- Senior Software Engineer\
- Master Degree - Information Technology And Management @ King Mongkut's Institute of Technology Ladkrabang - KMITL\
- Bachelor Degree - Computer Engineering @ King Mongkut's University of Technology Thonburi - KMUTT