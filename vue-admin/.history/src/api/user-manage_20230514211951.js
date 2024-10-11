import request from '@/utils/request'

export function getUserInfo(params) {
  return request({
    url: '/manage/getUserInfo',
    method: 'get',
    params
  })
}
export function addUser(params) {
    return request({
      url: '/manage/addUserInfo',
      method: 'post',
      params
    })
  }