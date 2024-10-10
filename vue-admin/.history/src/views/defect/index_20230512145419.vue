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
          @submit.prevent="handleResponse"
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
  data() {
    return {
      ori_url: '',
      result_url: ''
    }
  },
  methods: {
    handleResponse(event) {
      // 阻止表单默认提交行为
      event.preventDefault()

      // 获取表单数据
      const formData = new FormData(event.target)

      // 发送POST请求
      fetch('http://localhost:8010/defectimg', {
        method: 'POST',
        body: formData
      })
        .then((response) => response.json()) // 解析响应数据为JSON对象
        .then((jsonResponse) => {
          // 从JSON对象中获取需要的数据
          this.ori_url = jsonResponse.ori_url
          this.result_url = jsonResponse.result_url
          const code = jsonResponse.code
          const data = jsonResponse.data

          // 在控制台输出获取的数据，用于测试
          console.log(`ori_url: ${this.ori_url}`)
          console.log(`result_url: ${this.result_url}`)
          console.log(`code: ${code}`)
          console.log(`data: ${data}`)

          // 更新页面上的图片
          const oriImage = document.querySelector(
            '.demo-image__placeholder el-image:nth-child(1)'
          )
          const resultImage = document.querySelector(
            '.demo-image__placeholder el-image:nth-child(3)'
          )
          oriImage.setAttribute('src', oriUrl)
          resultImage.setAttribute('src', resultUrl)
        })
        .catch((error) => {
          console.error(error)
        })
    }
  }
}
</script>
