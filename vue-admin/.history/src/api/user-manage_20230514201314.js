import request from '@/utils/request'

export function getUserInfo(params) {
  return request({
    url: '/manage/G',
    method: 'get',
    params
  })
}
