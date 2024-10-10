import request from '@/utils/request'

export function getchart(token) {
  return request({
    url: '/data/chart',
    method: 'get',
    params: { token }
  })
}
