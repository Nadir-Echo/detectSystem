<template>
  <div class="demo-image__placeholder">
    <el-row type="flex" justify="center">
      <el-col :span="8">
        <el-image :src="ori_url" style="display: block" />
      </el-col>
      <el-col :span="8" style="text-align: center">
        <el-upload
          ref="upload"
          class="upload-demo"
          
          :on-preview="handlePreview"
          :on-remove="handleRemove"
          :file-list="fileList"
          :auto-upload="false"
        >
          <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
          <el-button
            style="margin-left: 10px"
            size="small"
            type="success"
          >检测</el-button>
          <!-- <div slot="tip" class="el-upload__tip">
            只能上传jpg/png文件，且不超过500kb
          </div> -->
        </el-upload>
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
import { submitDefect } from '@/api/defect'
<script>
export default {
  data() {
    return {
      ori_url: '', // 用于存储原始图片的 url
      result_url: '', // 用于存储检测结果图片的 url
      fileList: [] // 用于存储上传文件列表的数组
    }
  },
  methods: {
    // 点击检测按钮时触发
    submitDefect() {
      // 获取上传文件的信息
      const file = this.fileList[0]
      const formData = new FormData()
      formData.append('file', file.raw)

      // 发送 HTTP POST 请求
      this.$http.post('http://localhost:8010/defectimg', formData).then(response => {
        const data = response.data

        // 在响应中获取 ori_url, result_url 和 code
        this.ori_url = data.ori_url
        this.result_url = data.result_url
        const code = data.code

        // 根据 code 判断检测结果是否成功
        if (code === 0) {
          // 检测成功
          console.log('检测成功')
        } else {
          // 检测失败
          console.log(`检测失败，错误码：${code}`)
        }
      }).catch(error => {
        console.log('出现错误：', error)
      })
    },

    // 点击预览按钮时触发
    handlePreview(file) {
      console.log('预览文件：', file)
    },

    // 点击删除按钮时触发
    handleRemove(file, fileList) {
      console.log('删除文件：', file, fileList)
    }
  }
}
// export default {
//   data() {
//     return {
//       ori_url: '',
//       result_url: '',
//       fileList: []
//     }
//   },
//   methods: {
//     handlePreview() {
//       // do something
//     },
//     handleRemove() {
//       // do something
//     },
//     submitDefect() {
//       // if (this.fileList.length === 0) {
//       //   this.$message.error('请先上传图片')
//       //   return
//       // }
//       const formData = new FormData()
//       formData.append('file', this.fileList[0].raw)
//       this.$confirm('检测需要一定的时间，是否继续？', '提示', {
//         confirmButtonText: '确定',
//         cancelButtonText: '取消',
//         type: 'warning'
//       }).then(() => {
//         this.$store.dispatch('defect/defectsteel').then(response => {
//           const data = response.data
//           this.ori_url = data.ori_url
//           this.result_url = data.result_url
//           if (response.code === 200) {
//             this.$message.success('图片检测通过')
//           } else {
//             this.$message.error('图片检测不通过，请重新上传')
//           }
//         }).catch(error => {
//           console.log(error)
//           this.$message.error('图片检测失败，请稍后重试')
//         })
//       }).catch(() => {
//         this.$message.info('已取消检测')
//       })
//     }
//   }
// }
</script>
