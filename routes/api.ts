import { Application, Router, Status } from 'https://deno.land/x/oak/mod.ts';

const router: Router = new Router();

router.post('/api', (ctx, next) => {
	ctx.response.body = { content: 'Add House' };
	ctx.response.type = 'json';
	ctx.response.status = Status.OK;
});

router.get('/api', (ctx, next) => {
	ctx.response.body = { content: 'Get House' };
	ctx.response.type = 'json';
	ctx.response.status = Status.OK;
});

router.get('/api/:id', (ctx, next) => {
	ctx.response.body = { content: 'Get House', id: ctx.params.id };
	ctx.response.type = 'json';
	ctx.response.status = Status.OK;
});

export default router;
