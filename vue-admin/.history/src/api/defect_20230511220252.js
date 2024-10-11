import request from "@/utils/request";

export function submitDefect(data) {
  return request({
    url: "/defectimg",
    method: "post",
    data,
  });
}
