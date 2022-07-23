import fetch from 'node-fetch';
import { FormData } from 'formdata-node';
// import FormData from 'node-fetch'; 不支持
export function get(url, options) {
  return fetch(url, options).then(response => {
    return response.text();
  });
}
export function post(url, data) {
  let formData = new FormData()
  Object.keys(data).forEach(key => {
    formData.set(key, data[key])
  })
  return fetch(url, {
    method: 'POST',
    credentials: 'same-origin',
    body: formData
  }).then(response => {
    return response.text();
  })
}
