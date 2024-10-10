<template>
  <div class="demo-image__placeholder">
    <el-row type="flex" justify="center">
      <el-col :span="8">
        <el-image :src="src" style="display: block;" />
      </el-col>
      <el-col :span="8" style="text-align: center;">
        <el-upload
          class="upload-demo"
          action="https://jsonplaceholder.typicode.com/posts/"
          :show-file-list="false"
        >
          <el-button size="small" type="primary" @click.native.prevent="handledefect">上传图片</el-button>
        </el-upload>
      </el-col>
      <el-col :span="8">
        <el-image :src="src" style="display: block;">
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
      src: 'https://joplin-1305917469.cos.ap-nanjing.myqcloud.com/Snipaste_2023-04-17_16-23-26.png'
    }
  }
}
</script>
