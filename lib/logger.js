"use strict";

class Logger {
  constructor(service) {
    this.service = service;
  }

  send(target, route, time = Date.now()) {
    console.log(`${this.service} = ${target} - ${route} - ${time}`);
  }
}

module.exports = Logger;
