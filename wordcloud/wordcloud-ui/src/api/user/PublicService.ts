/**
 * Generated by unionj-generator.
 * Don't edit!
 *
 * @module Public
 */
import BizService from './BizService';
import type { PublicLogInResp, PublicSignUpResp } from './types';

export class PublicService extends BizService {
  constructor(axios: any) {
    super(axios);
  }

  /**
   * GET /public/download/avatar
   *
   * GetPublicDownloadAvatar 下载头像接口
   * 展示如何定义文件下载接口
   * 函数签名的出参里必须有且只有一个*os.File类型的参数
   * GetPublicDownloadAvatar is avatar download api
   * demo how to define file download api
   * NOTE: there must be one and at most one *os.File output parameter
   * @param userId 用户ID
   user id
   * @returns Promise<any>
   */
  getPublicDownloadAvatar(params: { userId: number }): Promise<any> {
    let client = this.axios.get;
    if (this.axios.$get) {
      client = this.axios.$get;
    }
    return client(this.addPrefix(`/public/download/avatar`), {
      params: {
        userId: params.userId,
      },
      responseType: 'blob',
    });
  }

  /**
   * POST /public/log/in
   *
   * PublicLogIn 用户登录接口
   * 展示如何鉴权并返回token
   * PublicLogIn is user login api
   * demo how to do authentication and issue token
   * @returns Promise<PublicLogInResp>
   */
  postPublicLogIn(params: { username: string; password: string }): Promise<PublicLogInResp> {
    let client = this.axios.post;
    if (this.axios.$post) {
      client = this.axios.$post;
    }
    const urlSearchParams = new URLSearchParams();
    urlSearchParams.append('username', params.username);
    urlSearchParams.append('password', params.password);
    return client(this.addPrefix(`/public/log/in`), urlSearchParams, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },
    });
  }

  /**
   * POST /public/sign/up
   *
   * PublicSignUp 用户注册接口
   * 展示如何定义POST请求且Content-Type是application/x-www-form-urlencoded的接口
   * PublicSignUp is user signup api
   * demo how to define post request api which accepts application/x-www-form-urlencoded content-type
   * @returns Promise<PublicSignUpResp>
   */
  postPublicSignUp(params: { username: string; password: string }): Promise<PublicSignUpResp> {
    let client = this.axios.post;
    if (this.axios.$post) {
      client = this.axios.$post;
    }
    const urlSearchParams = new URLSearchParams();
    urlSearchParams.append('username', params.username);
    urlSearchParams.append('password', params.password);
    return client(this.addPrefix(`/public/sign/up`), urlSearchParams, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },
    });
  }
}

export default PublicService;
