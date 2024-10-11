import request from "@/utils/request";

export function getInfo(token) {
  return request({
    url: "/data/chart",
    method: "get",
    params: { token },
  });
}
