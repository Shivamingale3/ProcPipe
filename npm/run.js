#!/usr/bin/env node

"use strict";

const { execFileSync } = require("child_process");
const path = require("path");
const fs = require("fs");

const ext = process.platform === "win32" ? ".exe" : "";
const binaryPath = path.join(__dirname, "bin", `procpipe${ext}`);

if (!fs.existsSync(binaryPath)) {
  console.error("\x1b[31mError: ProcPipe binary not found.\x1b[0m");
  console.error("");
  console.error("Try reinstalling:");
  console.error("  npm install -g procpipe");
  console.error("");
  console.error(`Expected binary at: ${binaryPath}`);
  process.exit(1);
}

try {
  execFileSync(binaryPath, process.argv.slice(2), {
    stdio: "inherit",
    env: process.env,
  });
} catch (err) {
  // execFileSync throws on non-zero exit code — propagate it
  process.exit(err.status != null ? err.status : 1);
}
