function formatMeasurement(name: string, tags: { [key: string]: string }, value: number): string {
  let tagStr = "";
  for (const [tagKey, tagVal] of Object.entries(tags)) {
    if (tagStr.length > 0) {
      tagStr += ",";
    }
    tagStr += `${tagKey}="${tagVal}"`;
  }
  if (tagStr.length > 0) {
    tagStr = `{${tagStr}}`;
  }

  const measurement = `${name}${tagStr} ${value} ${new Date().getTime()}`;
  return measurement;
}

function log(msg: string, ...args: unknown[]): void {
  console.log(`[${new Date().toISOString()}] ${msg}`, ...args);
}

export { formatMeasurement, log };
