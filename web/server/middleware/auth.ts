export default defineEventHandler((event) => {
  const { token } = parseCookies(event);

  if (token) {
    event.node.req.headers["authorization"] = `Bearer ${token}`;
  }
});
