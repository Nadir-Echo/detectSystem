<template>
  <div class="demo-image__placeholder">
    <el-row type="flex" justify="center">
      <el-col :span="8">
        <el-image :src="ori_url" style="display: block">
          <div slot="error" class="image-slot">
            <i class="el-icon-picture-outline" />
          </div>
        </el-image>
      </el-col>
      <el-col :span="8" style="text-align: center">
        <form
          action="http://localhost:8010/defectimg"
          method="post"
          enctype="multipart/form-data"
          @submit.prevent="handleResponse"
        >
          <div class="input-container">
            <input type="file" accept="image/*" name="file" ></input>
            <input type="submit" value="提交" ></input>
          </div>
          <!-- 上传文件:<input type="file" name="file" />
          <input type="submit" value="提交" /> -->
        </form>
      </el-col>
      <el-col :span="8">
        <el-image :src="result_url" style="display: block">
          <div slot="error" class="image-slot">
            <i class="el-icon-picture-outline" />
          </div>
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

          // // 更新页面上的图片
          // const oriImage = document.querySelector(
          //   '.demo-image__placeholder el-image:nth-child(1)'
          // )
          // const resultImage = document.querySelector(
          //   '.demo-image__placeholder el-image:nth-child(3)'
          // )
          // oriImage.setAttribute('src', this.ori_url)
          // resultImage.setAttribute('src', this.result_url)
        })
        .catch((error) => {
          console.error(error)
        })
    }
  }
}
</script>
<style>
.input-container {
  justify-content: flex-start;
  display: flex;
  /* align-items: center; */

  flex-direction: column;
  align-items: flex-start;;
}

.input-container input[type='file'] {
  margin-bottom: 10px;
  margin-right: 10px;
}

label {
  position: relative;
}
#fileinp {
  position: absolute;
  left: 0;
  top: 0;
  opacity: 0;
}
#btn {
  margin-right: 5px;
}
#text {
  color: red;
}
</style>
