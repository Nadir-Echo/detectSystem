import request from '@/utils/request'

export function v(params) {
  return request({
    url: '/data/list',
    method: 'get',
    params
  })
}
