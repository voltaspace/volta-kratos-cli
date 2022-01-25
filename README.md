# volta-kratos-cli
~~~~
go get -u github.com/voltaspace/volta-kratos-cli/kratos/v2@latest 脚手架增强版
~~~~
增强版相比阿B官方提供的脚手架增加以下功能:   
1.去除生成的结构体json:"omitempty"标签  
2.支持protobuf文件使用注解，可以结合gorm validate等功能使用，注解最终转换为结构体tag  
~~~~
例子:
message testMessage {  
    // @tag: gorm:"-"  
    string test1 = 1;  
    string test2 = 2;  
}
生成结果:  
type TestMessage struct{    
    Test1 string `protobuf:"varint,1,opt,name=test1,proto3" json:"test1" gorm:"-"`   
    Test2 string `protobuf:"varint,2,opt,name=test2,proto3" json:"test2"`  
`}  
~~~~
