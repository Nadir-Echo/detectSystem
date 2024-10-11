import request from '@/utils/request'

export function getUserInfo(params) {
  return request({
    url: '/manage/',
    method: 'get',
    params
  })
}
