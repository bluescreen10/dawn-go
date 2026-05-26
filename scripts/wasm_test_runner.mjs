/**
 * Headless browser test runner for Go WASM tests.
 *
 * Usage:
 *   node wasm_test_runner.mjs <test.wasm> <wasm_exec.js> [go test flags...]
 *
 * Requires:
 *   npm install playwright
 *   npx playwright install chromium
 */

import { chromium } from "playwright";
import http from "http";
import fs from "fs";
import path from "path";

const [, , wasmPath, wasmExecPath, ...testArgs] = process.argv;

if (!wasmPath || !wasmExecPath) {
  console.error("usage: wasm_test_runner.mjs <test.wasm> <wasm_exec.js> [go test flags...]");
  process.exit(1);
}

// ── HTTP server ────────────────────────────────────────────────────────────

function serveFile(res, filePath, contentType) {
  try {
    const data = fs.readFileSync(filePath);
    res.writeHead(200, { "Content-Type": contentType });
    res.end(data);
  } catch (e) {
    res.writeHead(500);
    res.end(String(e));
  }
}

// go test flags are passed as the WASM argv via the URL query string.
const queryArgs = testArgs.map(encodeURIComponent).join("&arg=");
const argsQuery = queryArgs ? `?arg=${queryArgs}` : "";

const HTML = `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>
// Override process.exit before wasm_exec.js so we can capture the exit code.
// wasm_exec.js only creates globalThis.process if it doesn't already exist.
globalThis.process = {
  getuid()  { return -1; },
  getgid()  { return -1; },
  geteuid() { return -1; },
  getegid() { return -1; },
  getgroups() { throw new Error("ENOSYS"); },
  pid: -1, ppid: -1,
  umask()   { throw new Error("ENOSYS"); },
  cwd()     { throw new Error("ENOSYS"); },
  chdir()   { throw new Error("ENOSYS"); },
  env: {},
  // exit is called by Go's runtime.wasmExit via the Go class's this.exit()
};
</script>
<script src="/wasm_exec.js"></script>
<script>
// Parse test flags from query string.
const args = ["js"].concat(
  [...new URLSearchParams(location.search).getAll("arg")]
);

const go = new Go();
go.argv = args;

// Capture exit code via the Go class's exit hook.
go.exit = (code) => {
  window.__exitCode = code;
  window.__done = true;
};

WebAssembly.instantiateStreaming(fetch("/test.wasm"), go.importObject)
  .then(({ instance }) => go.run(instance))
  .catch((err) => {
    console.error("WASM error:", String(err));
    window.__exitCode = 1;
    window.__done = true;
  });
</script>
</head>
<body>Running Go WASM tests...</body>
</html>`;

const server = http.createServer((req, res) => {
  const url = req.url.split("?")[0];
  if (url === "/") {
    res.writeHead(200, { "Content-Type": "text/html" });
    res.end(HTML);
  } else if (url === "/wasm_exec.js") {
    serveFile(res, wasmExecPath, "application/javascript");
  } else if (url === "/test.wasm") {
    serveFile(res, wasmPath, "application/wasm");
  } else {
    res.writeHead(404);
    res.end("not found");
  }
});

await new Promise((resolve) => server.listen(0, "127.0.0.1", resolve));
const { port } = server.address();

// ── Browser ────────────────────────────────────────────────────────────────

const browser = await chromium.launch({
  headless: true,
  args: [
    // Enable WebGPU (required in many Chromium builds).
    "--enable-unsafe-webgpu",
    // Use SwiftShader (CPU Vulkan) so tests work on machines without a GPU.
    "--use-vulkan=swiftshader",
    "--enable-features=Vulkan",
    "--disable-vulkan-surface",
    "--disable-dawn-features=disallow_unsafe_apis",
  ],
});

const page = await browser.newPage();

// Forward all browser console output to the terminal.
page.on("console", (msg) => {
  process.stdout.write(msg.text() + "\n");
});
page.on("pageerror", (err) => {
  process.stderr.write("page error: " + String(err) + "\n");
});

await page.goto(`http://127.0.0.1:${port}/${argsQuery}`);

// Wait for the Go program to call go.exit().
// Timeout is generous because WebGPU device initialisation can be slow.
await page.waitForFunction("window.__done === true", { timeout: 120_000 });

const exitCode = await page.evaluate("window.__exitCode ?? 1");

await browser.close();
server.close();

process.exit(exitCode);
