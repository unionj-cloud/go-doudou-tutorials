const ENV = {
  "env_dev": {
    "GDD_ENV": "dev"
  },
  "env_prod": {
    "GDD_ENV": "prod"
  },
  "env_test": {
    "GDD_ENV": "test"
  }
};

module.exports = {
  "apps": [
    {
      "name": "usersvc",
      "script": "./api",
      "cwd": "/root/deploy/go-doudou-tutorials/usersvc",
      "exec_interpreter": "",
      "exec_mode": "fork",
      ...ENV
    },
  ],
  deploy: {
    test: {
      host: ['162.14.116.92'],
      user: 'root',
      ssh_options: "StrictHostKeyChecking=no",
      ref: 'origin/master',
      repo: "git@github.com:unionj-cloud/go-doudou-tutorials.git",
      path: "/root/deploy/go-doudou-tutorials",
      "post-deploy": "sh ./usersvc/deploy.sh test",
    }
  }
};
