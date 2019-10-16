import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/home/dnslist',
    method: 'get',
    params
  })
}
