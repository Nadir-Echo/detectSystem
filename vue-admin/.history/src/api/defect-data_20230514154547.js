import request from '@/utils/request'

export function getDataForm(params) {
  return request({
    url: '/data/list',
    method: 'get',
    params
  })
}
