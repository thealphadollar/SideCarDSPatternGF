const express = require("express");
const app = express();

const httpProxy = require("http-proxy");
const apiProxy = httpProxy.createProxyServer({});

const SRV_NAME = process.env.SRV_NAME || "localhost";

const Logger = require("../lib/logger");
const logger = new Logger(SRV_NAME);

const { TARGET = "http://localhost:8080", PORT = 80 } = process.env;

app.all("/*", (req, res) => {
  logger.send(TARGET, req.url, Date.now());
  apiProxy.web(req, res, { target: TARGET });
});

app.listen(PORT, () => {
  logger.send(`Listening on port ${PORT}`);
});
