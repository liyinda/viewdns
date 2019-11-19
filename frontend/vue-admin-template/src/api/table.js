import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/home/table',
    method: 'get',
    params
  })
}

export function addDomain(params) {
  return request({
    url: '/home/adddomain',
    method: 'post',
    params
  })
}

export function delDomain(params) {
  return request({
    url: '/home/deldomain',
    method: 'post',
    params
  })
}
