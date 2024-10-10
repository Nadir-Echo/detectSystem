<template>
  <div class="demo-image__placeholder">
    <el-row type="flex" justify="center">
      <el-col :span="8">
        <el-image :src="ori_url" style="display: block" />
      </el-col>
      <el-col :span="8" style="text-align: center">
        <form
          action="http://localhost:8010/defectimg"
          method="post"
          enctype="multipart/form-data"
          onsubmit="handleResponse(this)"
        >
          上传文件:<input type="file" name="file" />
          <input type="submit" value="提交" />
        </form>
      </el-col>
      <el-col :span="8">
        <el-image :src="result_url" style="display: block">
          <div slot="placeholder" class="image-slot">
            加载中<span class="dot">...</span>
          </div>
        </el-image>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  // 在这里定义handleResponse函数
  function handleResponse(response) {
  // 将响应数据解析为JSON对象
  const jsonResponse = JSON.parse(response);

  // 从JSON对象中获取需要的数据
  const oriUrl = jsonResponse.ori_url;
  const resultUrl = jsonResponse.result_url;
  const code = jsonResponse.code;
  const data = jsonResponse.data;

  // 在控制台输出获取的数据，用于测试
  console.log(`ori_url: ${oriUrl}`);
  console.log(`result_url: ${resultUrl}`);
  console.log(`code: ${code}`);
  console.log(`data: ${data}`);

  // 更新页面上的图片
  const oriImage = document.querySelector(".demo-image__placeholder el-image:nth-child(1)");
  const resultImage = document.querySelector(".demo-image__placeholder el-image:nth-child(3)");
  oriImage.setAttribute("src", oriUrl);
  resultImage.setAttribute("src", resultUrl);
}
}
</script>