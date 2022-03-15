import {getToken} from '/@/utils/auth';


export const requestUse = (config: any) => {
  const token = getToken()
  if (token) {
    config.headers['Authorization'] = `Bearer ${token}` // 让每个请求携带自定义 token 请根据实际情况自行修改
  }
  return config
}
