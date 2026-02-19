#!/usr/bin/env node

"use strict";

const { execFileSync } = require("child_process");
const path = require("path");
const fs = require("fs");

const args = process.argv.slice(2);
const subcommand = args[0];

// Intercept commands that conflict with npm's package management
const NPM_INTERCEPTED = {
  install: {
    message: "ProcPipe is already installed globally via npm.",
    hint: "It's available as 'procpipe' from any terminal.\nIf it's not working, try: npm install -g procpipe",
  },
  uninstall: {
    message: "ProcPipe was installed via npm. Use npm to uninstall:",
    hint: "npm uninstall -g procpipe",
  },
};

if (subcommand && NPM_INTERCEPTED[subcommand]) {
  const { message, hint } = NPM_INTERCEPTED[subcommand];
  console.log(`\x1b[33m${message}\x1b[0m`);
  console.log(`  ${hint}`);
  process.exit(0);
}

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
  execFileSync(binaryPath, args, {
    stdio: "inherit",
    env: process.env,
  });
} catch (err) {
  // execFileSync throws on non-zero exit code — propagate it
  process.exit(err.status != null ? err.status : 1);
}
