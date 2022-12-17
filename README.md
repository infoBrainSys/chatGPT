# chatGPT for workflow(快捷指令)
由GPT机器人编写的一个简单的chatGPT机器人后端
疯狂套娃GPT持续开发中

### 运行容器
`docker run -itd --name chatGPTForWorkflow -p 80:80 chatgpt-for-workflow `

### 配置
进入容器

`docker exec -it chatGPTForWorkflow bash`

修改配置文件 `config/config.json`，修改key

### 启动
`go build -o chatGPTForWorkflow main.go`

`./chatGPTForWorkflow`

## 快捷指令配置
下载地址(Safari打开): 

https://www.icloud.com/shortcuts/121f16f23dd04b3a8a02ce4b5245cbc7


### 配置内容
将文本内容修改成自己的IP或者域名即可:

格式为: `http://`+`域名/IP`+`/?prompt=`
`https://your_ip/?prompt=`

### 注意⚠️
1. 如果本地部署需要保持手机wifi与服务器在同一网段
2. 域名访问可自行配置反代nginx或candy
