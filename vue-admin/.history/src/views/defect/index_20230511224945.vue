<template>
  <div class="demo-image__placeholder">
    <el-row type="flex" justify="center">
      <el-col :span="8">
        <el-image :src="src" style="display: block" />
      </el-col>
      <el-col :span="8" style="text-align: center">
        <el-upload
          ref="upload"
          class="upload-demo"
          action="https://jsonplaceholder.typicode.com/posts/"
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
            @click="submitDefect"
          >检测</el-button>
          <!-- <div slot="tip" class="el-upload__tip">
            只能上传jpg/png文件，且不超过500kb
          </div> -->
        </el-upload>
      </el-col>
      <el-col :span="8">
        <el-image :src="src" style="display: block">
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
      src1: '',
      src2: '',
      fileList: []
    }
  },
  methods: {
    handlePreview() {
      // do something
    },
    handleRemove() {
      // do something
    },
    submitDefect() {
      if (this.fileList.length === 0) {
        this.$message.error('请先上传图片')
        return
      }
      const formData = new FormData()
      formData.append('file', this.fileList[0].raw)
      this.$confirm('检测需要一定的时间，是否继续？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.$store.dispatch('').then(response => {
          const data = response.data
          this.src1 = data.src1
          this.src2 = data.src2
          if (data.result === 'OK') {
            this.$message.success('图片检测通过')
          } else {
            this.$message.error('图片检测不通过，请重新上传')
          }
        }).catch(error => {
          console.log(error)
          this.$message.error('图片检测失败，请稍后重试')
        })
      }).catch(() => {
        this.$message.info('已取消检测')
      })
    }
  }
}
</script>
