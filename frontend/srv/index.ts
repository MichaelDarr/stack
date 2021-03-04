import { createApp, serveStatic } from 'https://servestjs.org/@v1.1.9/mod.ts';

const app = createApp();
app.use(serveStatic('./public'));
app.listen({ port: 8080 });
