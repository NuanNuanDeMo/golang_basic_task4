<h3>1. github上拉取代码</h3>
   git clone https://github.com/NuanNuanDeMo/golang_basic_task4.git
<h3>2. 进入目录</h3>
    cd golang_basic_task4
<h3>3. 安装依赖</h3>
    <h4>go mod tidy</h4>
    <h4>go install github.com/gravityblast/fresh@latest</h4>
<h3>4. 初始化数据库</h3>
    mysq8,然后执行task4.sql语句;执行完成后修改module/jdbccon.go里面的ip账号密码
<h3>5. 运行工程</h3>
    go run main.go或者fresh
<h3>测试</h3>
在Postman中导入"任务4.postman_collection.json"文件即可进行测试

