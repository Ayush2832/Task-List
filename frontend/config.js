// config.js
// This file is generated at deploy time from .env
// For local dev: set window.API_BASE_URL manually or use the .env.local file

window.APP_CONFIG = {
  // In production, this is replaced by your CI/CD pipeline using the .env file
  // e.g., sed -i "s|__API_BASE_URL__|$API_BASE_URL|g" config.js
  API_BASE_URL: "http://localhost:8080",
};