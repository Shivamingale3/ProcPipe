#!/usr/bin/env node

"use strict";

const https = require("https");
const http = require("http");
const fs = require("fs");
const path = require("path");
const { execSync } = require("child_process");

const REPO = "Shivamingale3/ProcPipe";

const PLATFORM_MAP = {
  "linux-x64": "procpipe-linux-amd64",
  "linux-arm64": "procpipe-linux-arm64",
  "darwin-x64": "procpipe-darwin-amd64",
  "darwin-arm64": "procpipe-darwin-arm64",
  "win32-x64": "procpipe-windows.exe",
};

function getBinaryName() {
  const key = `${process.platform}-${process.arch}`;
  const binary = PLATFORM_MAP[key];

  if (!binary) {
    const supported = Object.keys(PLATFORM_MAP)
      .map((k) => k.replace("-", "/"))
      .join(", ");
    console.error(
      `\x1b[31mError: ProcPipe does not support ${process.platform}/${process.arch}\x1b[0m`,
    );
    console.error(`Supported platforms: ${supported}`);
    process.exit(1);
  }

  return binary;
}

function getVersion() {
  const pkg = JSON.parse(
    fs.readFileSync(path.join(__dirname, "package.json"), "utf8"),
  );
  return pkg.version;
}

function getLocalBinaryPath() {
  const binDir = path.join(__dirname, "bin");
  const ext = process.platform === "win32" ? ".exe" : "";
  return path.join(binDir, `procpipe${ext}`);
}

function download(url) {
  return new Promise((resolve, reject) => {
    const client = url.startsWith("https") ? https : http;

    client
      .get(url, { headers: { "User-Agent": "procpipe-npm" } }, (res) => {
        // Handle redirects (GitHub releases use 302)
        if (
          res.statusCode >= 300 &&
          res.statusCode < 400 &&
          res.headers.location
        ) {
          return download(res.headers.location).then(resolve).catch(reject);
        }

        if (res.statusCode !== 200) {
          reject(
            new Error(`Download failed with status ${res.statusCode}: ${url}`),
          );
          return;
        }

        const chunks = [];
        res.on("data", (chunk) => chunks.push(chunk));
        res.on("end", () => resolve(Buffer.concat(chunks)));
        res.on("error", reject);
      })
      .on("error", reject);
  });
}

async function install() {
  const binaryName = getBinaryName();
  const version = getVersion();
  const destPath = getLocalBinaryPath();
  const binDir = path.dirname(destPath);
  const url = `https://github.com/${REPO}/releases/download/v${version}/${binaryName}`;

  // Ensure bin directory exists
  if (!fs.existsSync(binDir)) {
    fs.mkdirSync(binDir, { recursive: true });
  }

  // Skip if binary already exists and is functional
  if (fs.existsSync(destPath)) {
    try {
      execSync(`"${destPath}" version`, { stdio: "ignore" });
      console.log("ProcPipe binary already installed, skipping download.");
      return;
    } catch {
      // Binary exists but is broken, re-download
    }
  }

  console.log(
    `Downloading ProcPipe v${version} for ${process.platform}/${process.arch}...`,
  );
  console.log(`URL: ${url}`);

  try {
    const data = await download(url);
    fs.writeFileSync(destPath, data);

    // Set executable permission on Unix
    if (process.platform !== "win32") {
      fs.chmodSync(destPath, 0o755);
    }

    console.log(
      `\x1b[32m✅ ProcPipe v${version} installed successfully!\x1b[0m`,
    );
  } catch (err) {
    console.error(`\x1b[31mFailed to download ProcPipe: ${err.message}\x1b[0m`);
    console.error("");
    console.error("You can install manually:");
    console.error(`  curl -L ${url} -o ${destPath} && chmod +x ${destPath}`);
    process.exit(1);
  }
}

install();
