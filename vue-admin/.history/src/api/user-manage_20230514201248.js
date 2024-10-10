import request from '@/utils/request'

export function getUserInfo(params) {
  return request({
    url: '/user/form',
    method: 'get',
    params
  })
}
