import { Result } from '../utils';

const fakeUserInfo = {
  userId: '1',
  username: 'qylz',
  realName: '管理员',
  desc: 'manager',
  password: 'Abc@12345',
  token: 'fakeToken1',
  roles: [
    {
      roleName: 'Super Admin',
      value: 'super',
    },
  ],
};
export default class UserService {
  async login() {
    return Result.success(fakeUserInfo);
  }

  async getUserInfoById() {
    return Result.success(fakeUserInfo);
  }
}
