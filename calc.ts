const process = Deno.run({
  cmd: ["C:\\Windows\\System32\\calc.exe"], // Command to launch the Windows Calculator
});

// Optionally, wait for the process to complete
await process.status();
