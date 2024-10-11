<template>
  <div class="demo-image__placeholder">
    <div class="block">
      <div style="position: relative">
        <el-image :src="src" style="display: inline-block">
          <span  class="demonstration">默认</span>
        </el-image>
      </div>
    </div>
    <div class="block">
      <div style="position: relative">
        <el-image :src="src" style="display: inline-block">
          <div
            slot="placeholder"
            class="image-slot"
            style="
              position: absolute;
              bottom: -20px;
              left: 50%;
              transform: translateX(-50%);
            "
          >
            自定义<span class="dot">...</span>
          </div>
        </el-image>
      </div>
    </div>
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
