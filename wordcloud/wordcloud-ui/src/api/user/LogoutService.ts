/**
* Generated by unionj-generator.
* Don't edit!
*
* @module Logout
*/
import BizService from "./BizService";
import type {
  GetLogoutResp,
} from "./types"

export class LogoutService extends BizService{

  constructor(axios: any) {
    super(axios);
  }

  /**
  * GET /logout
  *
  * GetLogout 注销token
  * GetLogout is used for revoking a token
  * https://github.com/dgrijalva/jwt-go/issues/214
  * @returns Promise<GetLogoutResp> 
  */
  getLogout(
  ) :Promise<GetLogoutResp> {
    let client = this.axios.get
    if(this.axios.$get) {
      client = this.axios.$get
    }
    return client(this.addPrefix(`/logout`),
          {
          }
        )
  }

}

export default LogoutService;

