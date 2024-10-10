import request from '@/utils/request'

export function getUserInfo(params) {
  return request({
    url: '/data/form',
    method: 'get',
    params
  })
}
