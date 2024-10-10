import request from "@/utils/request";
export function getInfo(para) {
  return request({
    url: "/user/info",
    method: "get",
    params: { token },
  });
}
