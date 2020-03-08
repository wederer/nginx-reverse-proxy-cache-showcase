# Caching with nginx as a reverse proxy

This demo shows the caching mechanic of a nginx reverse proxy (https://www.nginx.com/blog/nginx-caching-guide/).
The webserver takes 5 seconds to send a response and the response is cached for 10 seconds by the nginx reverse proxy.

For the config of the nginx see `nginx.conf`.

## How to run:
1. `docker-compose up`
2. Open browser and load `http://localhost`
3. The request will take 5 seconds and the container `web_1` will log `Request received`.
4. After the page finishes loading -> reload the page to see that the page loads instantly and no `Request received` is logged (the cache valid for 10 seconds)
5. Try to load the cached page in multiple browsers to see the functionality of a shared cache

## Bottom line
Using nginx as a reverse proxy is a great way to add caching and reduce load on the backend.
