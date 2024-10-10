import request from '@/utils/request'

export function getChart(token) {
  return request({
    url: '/data/chart',
    method: 'get',
    params: { token }
  })
}
