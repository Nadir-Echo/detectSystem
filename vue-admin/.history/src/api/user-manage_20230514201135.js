import request from '@/utils/request'

export function getDataForm(params) {
  return request({
    url: '/data/form',
    method: 'get',
    params
  })
}
