<template>
  <div>
    <img :src="image1" alt="Image 1" />
    <img :src="image2" alt="Image 2" />
    <button @click="loadImage">Load Image</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      image1: '',
      image2: ''
    }
  },
  methods: {
    loadImage() {
      // 获取用户选择的图片文件
      const input = document.createElement('input')
      input.type = 'file'
      input.accept = 'image/*'
      input.addEventListener('change', () => {
        const file = input.files[0]
        if (file) {
          // 创建一个 FileReader 对象，用于读取文件内容
          const reader = new FileReader()
          reader.readAsDataURL(file)
          reader.onload = () => {
            // 将图片数据显示在第一个图片框中
            this.image1 = reader.result
          }
        }
      })
      input.click()
    }
  }
}
</script>
