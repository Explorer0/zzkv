# zzkv
simple key-value storager

# Directory Description
```
├── LICENSE 
├── README.md
├── abstraction.go           //数据抽象文件*
├── go.mod
├── go.sum
├── storage.go               //存储器实现文件*
├── test                     //单元测试包
│   ├── bitcher.zzkv         //测试生成
│   ├── compression_test.go  //压缩器测试
│   ├── fucker.zzkv          //测试生成
│   ├── storager_test.go     //存储器测试
│   ├── test.sh
│   └── zzkv_test.go         //总体测试
├── tmp_test                 //临时测试文件夹
│   └── test.go
├── tree.txt
└── zzkv.go                  //zzkv主文件
```