import request from '@/utils/request'

export function defectsteel(data) {
  return request({
    url: 'api/v1/defectimg',
    method: 'post',
    data
  })
}
