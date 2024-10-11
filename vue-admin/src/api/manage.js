import request from '@/utils/request'

export function getUserInfo(params) {
  return request({
    url: '/manage/getUserInfo',
    method: 'get',
    params
  })
}

export function addUser(data) {
  return request({
    url: '/manage/addUser',
    method: 'post',
    data
  })
}
