import request from '@/utils/request'

export function getDataForm(params) {
  return request({
    url: '/data/form',
    method: 'get',
    params
  })
}

export function getImg(pid) {
  return request({
    url: '/data/get-img',
    method: 'get',
    params:{pid}
  })
}
