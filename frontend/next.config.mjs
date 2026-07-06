const apiOrigin = process.env.API_PROXY_ORIGIN || "http://localhost:18080";

/** @type {import('next').NextConfig} */
const nextConfig = {
  async rewrites() {
    return [
      {
        source: "/backend/:path*",
        destination: `${apiOrigin}/:path*`
      }
    ];
  }
};

export default nextConfig;
