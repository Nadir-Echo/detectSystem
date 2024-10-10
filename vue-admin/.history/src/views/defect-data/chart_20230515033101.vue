<template>
  <div id="lineChart" :style="{ width: '600px', height: '600px' }" />
  <div id="myChart" :style="{ width: '600px', height: '600px' }" />
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
    this.renderChart()
  },
  methods: { renderChart() {
      const chart = echarts.init(document.getElementById('myChart'));
      
      // 在这里配置你的饼状图数据和样式
      const options = {
        // 配置项...
      };
      
      chart.setOption(options);
    },
    drawLine() {
      // 基于刚刚准备好的 DOM 容器，初始化 EChart 实例
      const lineChart = this.$echarts.init(document.getElementById('lineChart'))
      // 绘制图表
      lineChart.setOption({
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
