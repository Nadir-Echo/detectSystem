import request from '@/utils/request'

export function submitDefect(data) {
  return request({
    url: 'user/defectimg',
    method: 'post',
    data
  })
}
