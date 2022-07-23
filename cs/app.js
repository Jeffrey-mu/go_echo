import { save } from "./routers/index.js";
save({ name: 'zs', email: 'zs@example.com' }).then(res => {
  console.log(res);
})
