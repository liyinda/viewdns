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
    url: '/passport/logout',
    method: 'post',
    params
  })
}

export function updateUser(data) {
  return request({
    url: '/home/useredit',
    method: 'put',
    params
  })
}
