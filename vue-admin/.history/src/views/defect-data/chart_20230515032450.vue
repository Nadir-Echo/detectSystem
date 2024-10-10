<template>
  <div id="myChart" :style="{ width: '300px', height: '300px' }" center/>
</template>

<script>
export default {
  name: 'Hello',
  data() {
    return {
      msg: 'Welcome to kalacloud.com'
    }
  },
  mounted() {
    this.drawLine()
  },
  methods: {
    drawLine() {
      // 基于刚刚准备好的 DOM 容器，初始化 EChart 实例
      const myChart = this.$echarts.init(document.getElementById('myChart'))
      // 绘制图表
      myChart.setOption({
        title: { text: '太阳系八大行星的卫星数量' },
        tooltip: {},
        xAxis: {
          data: ['水星', '金星', '地球', '火星', '木星', '土星', '金王星', '海王星']
        },
        yAxis: {},
        series: [{
          name: '数量',
          type: 'bar',
          data: [0, 0, 1, 2, 79, 82, 27, 14]
        }]
      })
    }
  }
}
</script>
