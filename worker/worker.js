export default {
    async fetch(request, env, ctx) {
      const url = new URL(request.url);
      const targetUrl = url.searchParams.get('url');
  
      if (!targetUrl) {
        return new Response(JSON.stringify({
          status: 'Error',
          message: 'Usage: ?url=URL_TO_CHECK'
        }), {
          headers: {
            'Content-Type': 'application/json',
            'Access-Control-Allow-Origin': '*'
          }
        });
      }
  
      try {
        const response = await fetch(targetUrl);
        const status = response.ok ? 'Up' : 'Down';
        return new Response(JSON.stringify({
          status,
          message: response.status.toString()
        }), {
          headers: {
            'Content-Type': 'application/json',
            'Access-Control-Allow-Origin': '*'
          }
        });
      } catch (error) {
        return new Response(JSON.stringify({
          status: 'Down',
          message: error.message
        }), {
          headers: {
            'Content-Type': 'application/json',
            'Access-Control-Allow-Origin': '*'
          }
        });
      }
    }
  }