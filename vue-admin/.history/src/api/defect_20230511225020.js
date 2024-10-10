import request from '@/utils/request'

export function defectsteel(data) {
  return request({
    url: 'user/defectimg',
    method: 'post',
    data
  })
}
