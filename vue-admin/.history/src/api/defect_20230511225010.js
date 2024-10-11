import request from '@/utils/request'

export function defect(data) {
  return request({
    url: 'user/defectimg',
    method: 'post',
    data
  })
}
