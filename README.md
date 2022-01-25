# volta-kratos-cli
go get -u github.com/voltaspace/volta-kratos-cli/kratos/v2@latest 脚手架增强版
增强版相比阿B官方提供的脚手架增加以下功能：    
1.去除生成的结构体json:"omitempty"标签  
2.支持protobuf文件使用注解，可以结合gorm validate等功能使用，注解最终转换为结构体tag  
例子：@tag: gorm:"-"
