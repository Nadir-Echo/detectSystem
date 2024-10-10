import request from '@/utils/request'

export function getUserInfo(params) {
  return request({
    url: '/manage/getUserInfo',
    method: 'get',
    params
  })
}
export function addUserInfo(params) {
    return request({
      url: '/manage/getUserInfo',
      method: 'get',
      params
    })
  }
