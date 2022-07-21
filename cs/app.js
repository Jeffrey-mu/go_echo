import fetch from 'node-fetch';
const LEN = 20
for (var i = 0; i < 20; i++) {

  get('name' + i, data)
}
async function get(name, data) {
  const response = await fetch('http://localhost:1323/show?name=' + name + '&data=' + data)
  const data = await response.text();
  console.log(data);
}
get()
