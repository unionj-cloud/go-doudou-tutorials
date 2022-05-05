// 不同部署环境下程序所需要读取的环境变量配置
const ENV = {
  // 开发环境
  "env_dev": {
    // GDD_ENV环境变量相当于Java生态的spring boot框架的spring.profiles.active配置
    "GDD_ENV": "dev"
  },
  // 生成环境
  "env_prod": {
    "GDD_ENV": "prod"
  },
  // 测试环境
  "env_test": {
    "GDD_ENV": "test"
  }
};

module.exports = {
  // 可以在apps属性里配置多个应用
  "apps": [
    {
      // 应用名称
      "name": "usersvc",
      // 启动文件，这里是go build编译以后的二进制可执行文件，可以是相对路径，也可以是绝对路径
      "script": "./api",
      // 工作目录，因为go-doudou-tutorials是一个monorepo仓库，所以需要加上这个配置，
      // 默认的工作目录是/root/deploy/go-doudou-tutorials/current
      "cwd": "/root/deploy/go-doudou-tutorials/current/usersvc",
      // 运行时环境，默认是nodejs的node。因为我们部署的是二进制可执行文件，所以不需要编译器
      "exec_interpreter": "",
      // 运行模式，pm2支持cluster和fork两种模式。
      // cluster模式，可以由pm2做负载均衡，但是只支持nodejs应用，
      // 所以如果是非nodejs应用，这里只能配置成fork
      "exec_mode": "fork",
      // javascript中的析构语法，相当于把ENV对象嵌入到这个对象里
      ...ENV
    },
  ],
  deploy: {
    // 部署test环境相关配置，可以配置任意多个任意名称的环境，
    // 比如uat，beta，release，production等等
    test: {
      // 可以配置多个ip地址，同时部署到多台测试服务器
      host: ['162.14.116.92'],
      // 服务器用户名
      user: 'root',
      // ssh相关配置
      ssh_options: "StrictHostKeyChecking=no",
      // 要部署的代码分支
      ref: 'origin/master',
      // git clone命令的代码仓库地址
      repo: "git@github.com:unionj-cloud/go-doudou-tutorials.git",
      // 代码的服务器磁盘路径
      path: "/root/deploy/go-doudou-tutorials",
      // post-deploy回调命令
      // 这里配置的命令的意思是切到usersvc路径下，然后执行deploy.sh脚本，传入test参数
      "post-deploy": "cd usersvc && sh deploy.sh test",
    }
  }
};