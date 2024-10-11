import request from '@/utils/request'

export function submitDefect(data) {
  return request({
    url: '/defect/submitDefect',
    method: 'post',
    data
  })
}