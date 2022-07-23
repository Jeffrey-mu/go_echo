import { post } from '../utils/fetch.js'
/**
*    @param {  Object} data {name: string, email: string}
*    @description
*    @Date 2022-07-23 20:25:43 星期六
*/
export function save(data) {
  return post('http://localhost:1323/save', data);
}
