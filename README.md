# chatGPT api
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
