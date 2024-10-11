import request from "@/utils/request";
export function getInfo(params) {
  return request({
    url: "/user/info",
    method: "get",
    params: { token },
  });
}
