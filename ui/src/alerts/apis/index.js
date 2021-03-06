import {proxy} from 'utils/queryUrlGenerator';

export function getAlerts(proxyLink) {
  return proxy({
    source: proxyLink,
    query: "select host, value, level, alertName from alerts",
    db: "chronograf",
  });
}
