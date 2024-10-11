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
            <input type="file" accept="image/*" name="file">
            <input type="submit" value="提交">
          </div>

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

import { getToken } from '@/utils/auth'
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

      // add token
      const myHeaders = new Headers()
      myHeaders.append('X-Token', getToken())

      // 发送POST请求
      fetch('http://localhost:8010/defectimg', {
        method: 'POST',
        body: formData,
        headers: myHeaders
      })
        .then((response) => response.json()) // 解析响应数据为JSON对象
        .then((jsonResponse) => {
          //judge jsonResponse.data is 50009
          if (jsonResponse.code === 50009) {
            this.$message({
              message: '图片中未检测到缺陷',
              type: 'warning'
            })
            return
          }else{

          }
         
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
  align-items: flex-start;
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
