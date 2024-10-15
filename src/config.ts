import { readFileSync } from "fs";

enum ConfigKey {
  TargetList = "TARGET_LIST",
  Options = "OPTIONS",
  TestIntervalMs = "TEST_INTERVAL_MS",
}

const loadedConfig: Partial<{ [key in ConfigKey]: string }> = {};

function getConfig(key: ConfigKey, defaultVal?: string): string {
  if (!loadedConfig[key]) {
    loadedConfig[key] = getConfigNoCache(key);
  }
  return loadedConfig[key] || defaultVal;
}

function getConfigNoCache(key: ConfigKey): string {
  if (process.env[key + "_FILE"]) {
    return readFileSync(process.env[key + "_FILE"])
      .toString()
      .trim();
  } else if (process.env[key]) {
    return process.env[key].trim();
  } else {
    return undefined;
  }
}

export { ConfigKey, getConfig };
