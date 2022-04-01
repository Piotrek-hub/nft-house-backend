import { Application, Status } from 'https://deno.land/x/oak/mod.ts';

import apiRouter from './routes/api.ts';

const app: Application = new Application();

app.use(apiRouter.routes());
app.use(apiRouter.allowedMethods());

console.log('App is listening on http://localhost:8080/');
await app.listen({ port: 8080 });
