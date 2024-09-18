import { check, sleep } from "k6";
import http from "k6/http";
import { Options } from "k6/options";

export let options: Options = {
  vus: 50, // Define the number of virtual users
  duration: "15m", // Define how long the test will run
  tags: {
    testid: `${getCurrentUTCTimestamp()}-runners`,
  },
};

export default () => {
  dispatchGitHubWorkflow();
};

export function getCurrentUTCTimestamp() {
  const now = new Date();
  return now.toISOString().replace("T", " ").substring(0, 19);
}

export function dispatchGitHubWorkflow() {
  const url = "https://test-api.k6.io";

  let res = http.get(url);
  check(res, {
    "is status 200": (r) => r.status === 200,
  });
  sleep(1);
}